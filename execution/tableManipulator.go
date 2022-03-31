package execution

import (
	"ZetaDB/container"
	"ZetaDB/storage"
)

type TableManipulator struct {
}

//DataPageManipulator generator
func GetTableManipulator() *TableManipulator {
	tm := &TableManipulator{}
	return tm
}

//delete a page
//if page is mode 0, just delete itself
//if page is mode 1/2, delete all pages in mode 1/2 related to it
func (tm *TableManipulator) DeleteDataPage(pageId uint32, schema *container.Schema) error {
	se := storage.GetStorageEngine()
	//get this page
	currentPage, err := se.GetDataPage(pageId, schema)
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
func (tm *TableManipulator) InsertTupleIntoTable(tableId uint32, tuple *container.Tuple) error {
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

func (tm *TableManipulator) DeleteTupleFromTable(tupleId uint32, tableId uint32, pageId uint32) {
}

func (tm *TableManipulator) UpdateTupleFromTable(tupleId uint32, tableId uint32, pageId uint32, newTuple *container.Tuple) {
}

//TODO use index to accalarate
func (tm *TableManipulator) QueryTupleFromTable() {}
