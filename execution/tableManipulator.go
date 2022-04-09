package execution

import (
	"ZetaDB/container"
	"ZetaDB/storage"
	"ZetaDB/utility"
)

type TableManipulator struct {
}

//DataPageManipulator generator
func GetTableManipulator() *TableManipulator {
	tm := &TableManipulator{}
	return tm
}

//create a new tail page in mode0 for this table, return this newly created page
//TODO unchecked
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
//TODO unchecked
func (tm *TableManipulator) NewTailPageMode12GroupToTable(tableId uint32, data []byte) {
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
	}

	//update k_table
	ktm.Update_k_table(tableId, mode1PageId, lastTupleId+1, tupleNum+1)
}

//delete a mode 0 page from this table
//update k_table if tuple number or tailPageId is changed
//TODO unchecked
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

	//update tupleNum
	tupleNum -= deletedPage.DpGetTupleNum()

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
//TODO unchecked
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
	currentGroupPage := mode1Page
	for {
		linkNextPageId, _ := currentGroupPage.DpGetLinkNextPageId()
		if currentGroupPage.DpGetPageId() == linkNextPageId {
			break
		}

		currentGroupPage.DpSetLinkPrePageId(currentGroupPage.DpGetPageId())
		currentGroupPage.DpSetLinkNextPageId(currentGroupPage.DpGetPageId())
		transaction.InsertDataPage(currentGroupPage)

		//return pageId
		ktm.InsertVacantDataPageId(currentGroupPage.DpGetPageId())

		currentGroupPage, _ = se.GetDataPage(linkNextPageId, schema)
	}

	//update k_table
	ktm.Update_k_table(tableId, tailPageId, lastTupleId, tupleNum-1)
}

func (tm *TableManipulator) InsertTupleIntoTable(tableId uint32, tuple *container.Tuple) {

}

func (tm *TableManipulator) DeleteTupleFromTable(tableId uint32, pageId uint32, tupleId uint32) {

}

func (tm *TableManipulator) UpdateTupleFromTable(tableId uint32, tupleId uint32, pageId uint32, newTuple *container.Tuple) {
}

//query a tuple according to its tupleId
//return the tuple and the pageId
//TODO use index to accalarate
func (tm *TableManipulator) QueryTupleFromTable(tableId uint32, tupleId uint32) (*container.Tuple, uint32) {
	return nil, 0
}
