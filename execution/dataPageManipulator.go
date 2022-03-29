package execution

import (
	"ZetaDB/container"
	"ZetaDB/storage"
	"errors"
)

type DataPageManipulator struct {
	se *storage.StorageEngine
}

//DataPageManipulator generator
func NewDataPageManipulator() *DataPageManipulator {
	dpm := &DataPageManipulator{
		se: storage.GetStorageEngine()}
	return dpm
}

//throw error if pageMode is 2
func (dpm *DataPageManipulator) GetDataHeadPageId(pageId uint32, schema *container.Schema) (uint32, error) {
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
func (dpm *DataPageManipulator) GetDataTailPageId(pageId uint32, schema *container.Schema) (uint32, error) {
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
func (dpm *DataPageManipulator) DeleteDataPage(pageId uint32, schema *container.Schema) error {
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
func (dpm *DataPageManipulator) InsertTupleIntoTable(tableId uint32, tuple *container.Tuple) error {
	//get table schema
	ktm := GetKeytableManager()
	tableSchema, err := ktm.Query_k_tableId_schema_FromTableId(tableId)
	if err != nil {
		return err
	}


	tailPageId := dpm.GetDataTailPageId()
	return nil
} */
