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

//create a new tail page in mode0 for this table, return this newly created page
func (tm *TableManipulator) NewTailPageMode0ToTable(tableId uint32) *storage.DataPage {
	ktm := GetKeytableManager()
	se := storage.GetStorageEngine()
	transaction := storage.GetTransaction()

	//get tailPageId, lastTupleId, tupleNum
	_, tailPageId, _, _, _ := ktm.Query_k_table(tableId)

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
	//ktm.Update_k_table_ForKeyTable(tableId, newTailPageId, lastTupleId, tupleNum)

	//create a new page
	newTailPage := storage.NewDataPageMode0(newTailPageId, tableId, oldTailPage.DpGetPageId(), newTailPageId)
	se.InsertDataPage(newTailPage)

	return newTailPage
}

func (tm *TableManipulator) DeletePageMode0FromTable(tableId uint32, pageId uint32) {

}

func (tm *TableManipulator) DeletePageMode1And2FromTable(tableId uint32, pageId uint32) {

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
