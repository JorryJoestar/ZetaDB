package execution

import (
	"ZetaDB/container"
	"ZetaDB/storage"
	"errors"
)

type TableManipulator struct {
	se *storage.StorageEngine
}

//DataPageManipulator generator
func GetTableManipulator() *TableManipulator {
	dpm := &TableManipulator{
		se: storage.GetStorageEngine()}
	return dpm
}

//throw error if pageMode is 2
func (dpm *TableManipulator) GetDataHeadPageId(pageId uint32, schema *container.Schema) (uint32, error) {
	currentDataPage, err := dpm.se.GetDataPage(pageId, schema)
	if err != nil {
		return 0, err
	}

	//throw error if pageMode is 2
	if currentDataPage.DataPageMode() == 2 {
		return 0, errors.New("execution/dataPageManipulator.go    GetDataHeadPageId() mode invalid")
	}

	//loop until reach headPage
	isHeadPage, _ := currentDataPage.DpIsHeadPage()
	for !isHeadPage {
		pageId, _ = currentDataPage.DpGetPriorPageId()
		currentDataPage, err = dpm.se.GetDataPage(pageId, schema)
		if err != nil {
			return 0, err
		}
		isHeadPage, _ = currentDataPage.DpIsHeadPage()
	}

	return pageId, nil
}

//throw error if pageMode is 2
func (dpm *TableManipulator) GetDataTailPageId(pageId uint32, schema *container.Schema) (uint32, error) {
	currentDataPage, err := dpm.se.GetDataPage(pageId, schema)
	if err != nil {
		return 0, err
	}

	//throw error if pageMode is 2
	if currentDataPage.DataPageMode() == 2 {
		return 0, errors.New("execution/dataPageManipulator.go    GetDataTailPageId() mode invalid")
	}

	//loop until reach tailPage
	isTailPage, _ := currentDataPage.DpIsTailPage()
	for !isTailPage {
		pageId, _ = currentDataPage.DpGetNextPageId()
		currentDataPage, err = dpm.se.GetDataPage(pageId, schema)
		if err != nil {
			return 0, err
		}
		isTailPage, _ = currentDataPage.DpIsTailPage()
	}

	return pageId, nil
}

//delete a page
//if page is mode 0, just delete itself
//if page is mode 1/2, delete all pages in mode 1/2 related to it
func (dpm *TableManipulator) DeleteDataPage(pageId uint32, schema *container.Schema) error {
	//get this page
	currentPage, err := dpm.se.GetDataPage(pageId, schema)
	if err != nil {
		return err
	}

	if currentPage.DataPageMode() == 0 { //simply delete this page
		//TODO
	} else { //delete a series of pages in mode 1/2
		//TODO
	}

	return nil
}

/* //insert a tuple into the end of this table
//if this tuple is too large to hold within a single page in mode0, break it into several part and insert into mode1/2 page series
//if no enough to hold this tuple in the last page, create a new page in mode0 to hold it
func (dpm *TableManipulator) InsertTupleIntoTable(tableId uint32, tuple *container.Tuple) error {
	//get table schema
	ktm := GetKeytableManager()
	se := storage.GetStorageEngine()
	tableSchema, err := ktm.Query_k_tableId_schema_FromTableId(tableId)
	if err != nil {
		return err
	}

	//get tailPageId, lastTupleId, tupleNum
	_, tailPageId, lastTupleId, tupleNum, err := ktm.Query_k_table_FromTableId(tableId)
	if err != nil {
		return err
	}

	//get the tail page of this table
	lastPage, err := se.GetDataPage(tailPageId, tableSchema)
	if err != nil {
		return err
	}

	//check if lastPage can hold this tuple, insert tuple and update k_table & k_emptyDataPageSlot
	tupleByteLengh := tuple.TupleSizeInBytes()
	remainByteLength := lastPage.DpVacantByteNum()
	if tupleByteLengh <= remainByteLength { //page can hold this tuple

	} else if tupleByteLengh > utility.DEFAULT_PAGE_SIZE-32 { // tuple large than a singe page

	} else { //need a new page to hold tuple

	}

	return nil
} */

func (dpm *TableManipulator) DeleteTupleFromTable(tupleId uint32, tableId uint32, pageId uint32) {
}

func (dpm *TableManipulator) UpdateTupleFromTable(tupleId uint32, tableId uint32, pageId uint32, newTuple *container.Tuple) {
}

//TODO use index to accalarate
func (dpm *TableManipulator) QueryTupleFromTable() {}
