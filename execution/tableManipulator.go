package execution

import (
	"ZetaDB/container"
	"ZetaDB/storage"
	"ZetaDB/utility"
	"errors"
	"sync"
)

type TableManipulator struct {
}

//for singleton pattern
var tmInstance *TableManipulator
var tmOnce sync.Once

//to get TableManipulator, call this function
func GetTableManipulator() *TableManipulator {
	tmOnce.Do(func() {
		tmInstance = &TableManipulator{}
	})

	return tmInstance
}

//create a new tail page in mode0 for this table, return this newly created page
func (tm *TableManipulator) NewTailPageMode0ToTable(tableId uint32) *storage.DataPage {
	ktm := GetKeytableManager()
	se := storage.GetStorageEngine()
	transaction := storage.GetTransaction()

	//get tailPageId, lastTupleId, tupleNum
	_, tailPageId, lastTupleId, tupleNum, _ := ktm.Query_k_table(tableId)

	//get schema of this table
	schema, _ := ktm.Query_k_tableId_schema_FromTableId(tableId)

	//get old tailPage
	oldTailPage, _ := se.GetDataPage(tailPageId, schema)

	//get a vacant dataPageId
	newTailPageId := ktm.GetVacantDataPageId()

	//update old tailPage
	oldTailPage.DpSetNextPageId(newTailPageId)
	transaction.InsertDataPage(oldTailPage)

	//update k_table
	ktm.Update_k_table(tableId, newTailPageId, lastTupleId, tupleNum)

	//create a new page
	newTailPage := storage.NewDataPageMode0(newTailPageId, tableId, oldTailPage.DpGetPageId(), newTailPageId)
	se.InsertDataPage(newTailPage)
	transaction.InsertDataPage(newTailPage)

	return newTailPage
}

//create a new tail page in mode1 for this table, insert data into this page
//if current page can not hold whole data, create new mode2 page to hold remain data
//for k_table, only update tailPageId here
//return mode1Page
func (tm *TableManipulator) NewTailPageMode12GroupToTable(tableId uint32, data []byte) *storage.DataPage {
	se := storage.GetStorageEngine()
	transaction := storage.GetTransaction()
	ktm := GetKeytableManager()

	//get schema
	schema, _ := ktm.Query_k_tableId_schema_FromTableId(tableId)

	//get tailPageId, lastTupleId, tupleNum
	_, tailPageId, lastTupleId, tupleNum, _ := ktm.Query_k_table(tableId)

	//get oldTailPageId
	oldTailPage, _ := se.GetDataPage(tailPageId, schema)

	//create mode1Page
	mode1PageId := ktm.GetVacantDataPageId()
	mode1Page := storage.NewDataPageMode1(mode1PageId, tableId, tailPageId, mode1PageId, mode1PageId, data[:utility.DEFAULT_PAGE_SIZE-32])
	se.InsertDataPage(mode1Page)
	transaction.InsertDataPage(mode1Page)

	data = data[utility.DEFAULT_PAGE_SIZE-32:]

	oldTailPage.DpSetNextPageId(mode1PageId)
	transaction.InsertDataPage(oldTailPage)

	//loop, create new mode2Page and keep pushing data into new pages
	linkPrePage := mode1Page
	linkPrePageId := mode1PageId
	for {
		if len(data) == 0 { //all data pushed into pages
			break
		}

		var dataToPush []byte
		if len(data) > utility.DEFAULT_PAGE_SIZE-32 {
			dataToPush = data[:utility.DEFAULT_PAGE_SIZE-32]
			data = data[utility.DEFAULT_PAGE_SIZE-32:]
		} else {
			dataToPush = data
			data = data[0:0]
		}

		newMode2PageId := ktm.GetVacantDataPageId()
		newMode2Page := storage.NewDataPageMode2(newMode2PageId, tableId, int32(len(dataToPush)), linkPrePageId, newMode2PageId, dataToPush)
		se.InsertDataPage(newMode2Page)
		transaction.InsertDataPage(newMode2Page)

		linkPrePage.DpSetLinkNextPageId(newMode2PageId)
		transaction.InsertDataPage(linkPrePage)

		linkPrePage = newMode2Page
		linkPrePageId = newMode2PageId
	}

	//update k_table
	ktm.Update_k_table(tableId, mode1PageId, lastTupleId, tupleNum)
	return mode1Page
}

//delete a mode 0 page from this table
//update k_table if tuple number or tailPageId is changed
func (tm *TableManipulator) DeletePageMode0FromTable(tableId uint32, pageId uint32) {
	se := storage.GetStorageEngine()
	transaction := storage.GetTransaction()
	ktm := GetKeytableManager()

	//get tailPageId, lastTupleId, tupleNum according to tableId
	_, tailPageId, lastTupleId, tupleNum, _ := ktm.Query_k_table(tableId)

	//get schema
	schema, _ := ktm.Query_k_tableId_schema_FromTableId(tableId)

	//get page that is going to be deleted
	deletedPage, _ := se.GetDataPage(pageId, schema)

	//return pageId of deletedPage
	ktm.InsertVacantDataPageId(pageId)

	//get prior and next page of deletedPage
	priorPageId, _ := deletedPage.DpGetPriorPageId()
	nextPageId, _ := deletedPage.DpGetNextPageId()
	priorPage, _ := se.GetDataPage(priorPageId, schema)
	nextPage, _ := se.GetDataPage(nextPageId, schema)

	//check if page that is going to be deleted is tail page
	//if so, update tailPageId
	if nextPageId == deletedPage.DpGetPageId() {
		tailPageId = priorPageId

		priorPage.DpSetNextPageId(priorPageId)
		transaction.InsertDataPage(priorPage)

		deletedPage.DpSetPriorPageId(deletedPage.DpGetPageId())
		transaction.InsertDataPage(deletedPage)
	} else {
		priorPage.DpSetNextPageId(nextPageId)
		transaction.InsertDataPage(priorPage)

		nextPage.DpSetPriorPageId(priorPageId)
		transaction.InsertDataPage(nextPage)

		deletedPage.DpSetPriorPageId(tableId)
		deletedPage.DpSetNextPageId(tableId)
		transaction.InsertDataPage(deletedPage)
	}

	//update k_table
	ktm.Update_k_table(tableId, tailPageId, lastTupleId, tupleNum)
}

//delete a group of mode1 and mode2 pages
//inputed pageId is the pageId of mode1 page
func (tm *TableManipulator) DeletePageMode1And2FromTable(tableId uint32, pageId uint32) {
	se := storage.GetStorageEngine()
	transaction := storage.GetTransaction()
	ktm := GetKeytableManager()

	//get tailPageId, lastTupleId, tupleNum
	_, tailPageId, lastTupleId, tupleNum, _ := ktm.Query_k_table(tableId)

	//get schema
	schema, _ := ktm.Query_k_tableId_schema_FromTableId(tableId)

	//get mode1Page
	mode1Page, _ := se.GetDataPage(pageId, schema)

	//get priorPage and tailPage of this mode1Page
	priorPageId, _ := mode1Page.DpGetPriorPageId()
	nextPageId, _ := mode1Page.DpGetNextPageId()
	priorPage, _ := se.GetDataPage(priorPageId, schema)
	nextPage, _ := se.GetDataPage(nextPageId, schema)

	//delete mode1Page from this table
	if pageId == tailPageId { //this mode1Page is the tailPage of this table
		//update tailPageId
		tailPageId = priorPageId

		priorPage.DpSetNextPageId(priorPageId)
		transaction.InsertDataPage(priorPage)

		mode1Page.DpSetPriorPageId(pageId)
		mode1Page.DpSetNextPageId(pageId)
		transaction.InsertDataPage(mode1Page)
	} else { //this mode1Page is not tailPage, no need to update tailPageId
		priorPage.DpSetNextPageId(nextPageId)
		transaction.InsertDataPage(priorPage)

		nextPage.DpSetPriorPageId(priorPageId)
		transaction.InsertDataPage(nextPage)

		mode1Page.DpSetPriorPageId(pageId)
		mode1Page.DpSetNextPageId(pageId)
		transaction.InsertDataPage(mode1Page)
	}

	//loop and delete all mode2 pages in this group
	linkNextPageId, _ := mode1Page.DpGetLinkNextPageId()
	mode1Page.DpSetLinkNextPageId(pageId)
	transaction.InsertDataPage(mode1Page)
	ktm.InsertVacantDataPageId(pageId)
	var currentGroupPage *storage.DataPage
	for {
		currentGroupPage, _ = se.GetDataPage(linkNextPageId, schema)

		currentGroupPage.DpSetLinkPrePageId(currentGroupPage.DpGetPageId())
		currentGroupPage.DpSetLinkNextPageId(currentGroupPage.DpGetPageId())
		transaction.InsertDataPage(currentGroupPage)

		//return pageId
		ktm.InsertVacantDataPageId(currentGroupPage.DpGetPageId())

		//update linkNextPageId
		linkNextPageId, _ = currentGroupPage.DpGetLinkNextPageId()

		if currentGroupPage.DpGetPageId() == linkNextPageId {
			break
		}
	}

	//update k_table
	ktm.Update_k_table(tableId, tailPageId, lastTupleId, tupleNum)
}

//insert a tuple into a table
//assign a new tupleId according to k_table
//if this tuple is too long and can not be pushed in one page, push it into a group of pages in mode1 and mode2
func (tm *TableManipulator) InsertTupleIntoTable(tableId uint32, tuple *container.Tuple) {
	ktm := GetKeytableManager()
	se := storage.GetStorageEngine()
	transaction := storage.GetTransaction()

	//get schema
	schema, _ := ktm.Query_k_tableId_schema_FromTableId(tableId)

	//get tailPageId, lastTupleId, tupleNum
	_, tailPageId, lastTupleId, tupleNum, _ := ktm.Query_k_table(tableId)

	//update tupleId of this inputed tuple
	tuple.TupleSetTupleId(lastTupleId + 1)

	//if this tuple is over large, create a group of mode1 and mode2 pages to hold it
	if tuple.TupleSizeInBytes() > utility.DEFAULT_PAGE_SIZE-32 {
		tupleBytes, _ := tuple.TupleToBytes()
		newPageMode1 := tm.NewTailPageMode12GroupToTable(tableId, tupleBytes)
		tailPageId = newPageMode1.DpGetPageId()
	} else {
		oldTailPage, _ := se.GetDataPage(tailPageId, schema)
		if oldTailPage.DataPageMode() == 0 && oldTailPage.DpVacantByteNum() > tuple.TupleSizeInBytes() { //current tailPage can hold this tuple
			oldTailPage.InsertTuple(tuple)
			transaction.InsertDataPage(oldTailPage)
		} else {
			newPageMode0 := tm.NewTailPageMode0ToTable(tableId)
			newPageMode0.InsertTuple(tuple)
			tailPageId = newPageMode0.DpGetPageId()
		}
	}

	//update k_table
	ktm.Update_k_table(tableId, tailPageId, lastTupleId+1, tupleNum+1)
}

//delete the tuple according to its tupleId
//return deletedTuple
func (tm *TableManipulator) DeleteTupleFromTable(tableId uint32, tupleId uint32) *container.Tuple {
	se := storage.GetStorageEngine()
	ktm := GetKeytableManager()
	transaction := storage.GetTransaction()

	//get deletedTuple and pageId
	deletedTuple, pageId, _ := tm.QueryTupleFromTableByTupleId(tableId, tupleId)

	//get headPage
	headPageId, _, _, _, _ := ktm.Query_k_table(tableId)

	//get schema
	schema, _ := ktm.Query_k_tableId_schema_FromTableId(tableId)

	//delete the tuple from corresponding page
	deletedFromPage, _ := se.GetDataPage(pageId, schema)

	if deletedFromPage.DataPageMode() == 1 {
		tm.DeletePageMode1And2FromTable(tableId, deletedFromPage.DpGetPageId())
	} else { //mode is 0

		deletedFromPage.DpDeleteTuple(tupleId)
		transaction.InsertDataPage(deletedFromPage)

		if deletedFromPage.DpGetTupleNum() == 0 && deletedFromPage.DpGetPageId() != headPageId {
			tm.DeletePageMode0FromTable(tableId, deletedFromPage.DpGetPageId())
		}
	}

	if deletedTuple != nil {
		//update k_table, tupleNum = tupleNum - 1
		_, tailPageId, lastTupleId, tupleNum, _ := ktm.Query_k_table(tableId)
		if deletedTuple.TupleGetTupleId() == lastTupleId {
			lastTupleId--
		}
		ktm.Update_k_table(tableId, tailPageId, lastTupleId, tupleNum-1)
	}

	return deletedTuple
}

//replace the tuple with tupleId by newTuple
func (tm *TableManipulator) UpdateTupleFromTable(tableId uint32, tupleId uint32, newTuple *container.Tuple) {
	//delete the tuple
	tm.DeleteTupleFromTable(tableId, tupleId)

	//update new inserted tupleId
	newTuple.TupleSetTupleId(tupleId)

	//insert the newTuple into table
	tm.InsertTupleIntoTable(tableId, newTuple)
}

//query a tuple according to its tupleId
//return the tuple and the pageId
//throw error if no such tuple
//TODO update use index to accalarate
func (tm *TableManipulator) QueryTupleFromTableByTupleId(tableId uint32, tupleId uint32) (*container.Tuple, uint32, error) {
	se := storage.GetStorageEngine()
	ktm := GetKeytableManager()

	//get headPageId, tailPageId of this table
	headPageId, _, _, _, _ := ktm.Query_k_table(tableId)

	//get schema
	schema, _ := ktm.Query_k_tableId_schema_FromTableId(tableId)

	//loop until find the target tuple
	var targetPage *storage.DataPage
	var targetTuple *container.Tuple

	currentPage, _ := se.GetDataPage(headPageId, schema)

	for {

		if currentPage.DataPageMode() == 0 { //mode 0
			var currentTuple *container.Tuple = nil

			for i := 0; i < int(currentPage.DpGetTupleNum()); i++ {

				currentTuple, _ = currentPage.GetTupleAt(i)

				if currentTuple.TupleGetTupleId() == tupleId {
					targetPage = currentPage
					targetTuple = currentTuple
					return targetTuple, targetPage.DpGetPageId(), nil
				}
			}
		} else { //mode 1
			var data []byte
			firstPageData, _ := currentPage.DpGetData()
			data = append(data, firstPageData...)

			firstMode2PageId, _ := currentPage.DpGetLinkNextPageId()
			mode2Page, _ := se.GetDataPage(firstMode2PageId, schema)

			for {
				mode2PageData, _ := mode2Page.DpGetData()
				data = append(data, mode2PageData...)

				isListTail, _ := mode2Page.DpIsListTailPage()
				if isListTail { //reach the list tail
					break
				}
			}

			currentTuple, _ := container.NewTupleFromBytes(data, schema, tableId)
			if currentTuple.TupleGetTupleId() == tupleId {
				return currentTuple, currentPage.DpGetPageId(), nil
			}
		}

		//update currentPage to next page
		nextPageId, _ := currentPage.DpGetNextPageId()

		if currentPage.DpGetPageId() == nextPageId { //no such tuple in this table
			return nil, 0, errors.New("execution/tableManipulator.go    QueryTupleFromTableByTupleId() no such tuple")
		}

		currentPage, _ = se.GetDataPage(nextPageId, schema)
	}
}
