package storage

import (
	. "ZetaDB/container"
	. "ZetaDB/utility"
	"sync"
)

type storageEngine struct {
	//head pages of 17 key tables

	//k_userId_userName(userId INT PRIMARY KEY, userName VARCHAR(20))
	//head page number 0, tableId 0
	headPage_k_userId_userName *dataPage

	//k_userId_password(userId INT PRIMARY KEY, password VARCHAR(20))
	//head page number 1, tableId 1
	headPage_k_userId_password *dataPage

	//k_tableId_userId(tableId INT PRIMARY KEY, userId INT)
	//head page number 2, tableId 2
	headPage_k_tableId_userId *dataPage

	//k_assertId_userId(assertId INT PRIMARY KEY, userId INT)
	//head page number 3, tableId 3
	headPage_k_assertId_userId *dataPage

	//k_viewId_userId(viewId INT PRIMARY KEY, userId INT)
	//head page number 4, tableId 4
	headPage_k_viewId_userId *dataPage

	//k_indexId_tableId(indexId INT PRIMARY KEY, tableId INT)
	//head page number 5, tableId 5
	headPage_k_indexId_tableId *dataPage

	//k_triggerId_userId(triggerId INT PRIMARY KEY, userId INT)
	//head page number 6, tableId 6
	headPage_k_triggerId_userId *dataPage

	//k_psmId_userId(psmId INT PRIMARY KEY, userId INT)
	//head page number 7, tableId 7
	headPage_k_psmId_userId *dataPage

	//k_tableId_schema(tableId INT PRIMARY KEY, schema VARCHAR(255))
	//head page number 8, tableId 8
	headPage_k_tableId_schema *dataPage

	//k_table(tableId INT PRIMARY KEY, headPageId INT, lastTupleIndexId INT, tupleNum INT)
	//head page number 9, tableId 9
	headPage_k_table *dataPage

	//k_assert(assertId INT PRIMARY KEY, assertStmt VARCHAR(255))
	//head page number 10, tableId 10
	headPage_k_assert *dataPage

	//k_view(viewId INT PRIMARY KEY, viewStmt VARCHAR(255))
	//head page number 11, tableId 11
	headPage_k_view *dataPage

	//k_index(indexId INT PRIMARY KEY, logHeadPageId INT)
	//head page number 12, tableId 12
	headPage_k_index *dataPage

	//k_trigger(triggerId INT PRIMARY KEY, triggerStmt VARCHAR(255))
	//head page number 13, tableId 13
	headPage_k_trigger *dataPage

	//k_psm(psmId INT PRIMARY KEY, psmStmt VARCHAR(255))
	//head page number 14, tableId 14
	headPage_k_psm *dataPage

	//k_emptyPageSlot(pageId INT)
	//head page number 15, tableId 15
	headPage_k_emptyPageSlot *dataPage

	//k_log page(log INT PRIMARY KEY, logStmt VARCHAR(20))
	//head page number 16, tableId 16
	headPage_k_log *dataPage

	//buffer for data page
	dBuffer *dataBuffer

	//buffer for index page
	iBuffer *indexBuffer

	//IOManipulator
	iom *IOManipulator
}

//use GetBufferPool to get the unique bufferPool
var seInstance *storageEngine
var seOnce sync.Once

func GetStorageEngine(dfl string, ifl string, lfl string) *storageEngine {
	iom, _ := GetIOManipulator(dfl, ifl, lfl)

	seOnce.Do(func() {
		seInstance = &storageEngine{
			dBuffer: NewDataBuffer(),
			iBuffer: NewIndexBuffer(),
			iom:     iom}
	})
	return seInstance
}

//TODO
//get a data page according to its pageId
//if this page is modified, remember to swap it
func (se *storageEngine) GetDataPage(pageId uint32, schema *Schema) (*dataPage, error) {
	bytes, err := se.iom.BytesFromDataFile(pageId, DEFAULT_PAGE_SIZE)
	if err != nil {
		return nil, err
	}

	page, pageErr := NewDataPageFromBytes(bytes)
	return nil, nil
}

//TODO
//insert a newly created data page into buffer, but not swapped into disk
//remember to swap it
func (se *storageEngine) InsertDataPage(page *dataPage) error {

	//check if dataBuffer is full
	if se.dBuffer.DataBufferIsFull() { //dataBuffer is full, evict and delete one page
		evictPage, evictErr := se.dBuffer.DataBufferEvictDataPage()
		if evictErr != nil {
			return evictErr
		}

		//if the evict page is modified, swap it into disk
		if evictPage.DataPageIsModified() {
			err3 := se.SwapDataPage(evictPage.DpGetPageId())
			if err3 != nil {
				return err3
			}
		}

		//delete the evict page from buffer
		evictPage.UnmodifyDataPage()
		deleteErr := se.dBuffer.DataBufferDeleteDataPage(evictPage.DpGetPageId())
		if deleteErr != nil {
			return deleteErr
		}
	}

	//insert the page into buffer
	err4 := se.dBuffer.DataBufferInsertDataPage(page)
	if err4 != nil {
		return err4
	}

	return nil

}

//TODO
//delete a data page according to its pageId, related page not swapped into disk
//remember to swap related pages
func (se *storageEngine) DeleteDataPage(pageId uint32) error {
	return se.dBuffer.DataBufferDeleteDataPage(pageId)
}

//TODO
//swap a data page into disk according to its pageId
func (se *storageEngine) SwapDataPage(pageId uint32) error {
	//get this page from buffer
	page, err := se.dBuffer.DataBufferFetchPage(pageId)
	if err != nil {
		return err
	}

	//convert this page into bytes
	bytes, err2 := page.DataPageToBytes()
	if err2 != nil {
		return err2
	}

	//push bytes into disk
	err3 := se.iom.BytesToDataFile(bytes, page.DpGetPageId())

	if err3 == nil { //delete page from buffer
		page.UnmodifyDataPage()
		err4 := se.dBuffer.DataBufferDeleteDataPage(page.DpGetPageId())
		if err4 != nil {
			return err4
		}
	}

	return err3
}

//get an index page according to its pageId from index buffer
//if this page is not buffered, fetch it from disk and push it into buffer
//if buffer is full, evict a page
//if this page is modified, remember to swap it
func (se *storageEngine) GetIndexPage(pageId uint32) (*indexPage, error) {
	page, err1 := se.iBuffer.IndexBufferFetchPage(pageId)

	if err1.Error() == "pageId invalid, this page is not buffered" { //not in buffer

		//fetch bytes from disk
		bytes, err2 := se.iom.BytesFromIndexFile(pageId, DEFAULT_PAGE_SIZE)
		if err2 != nil {
			return nil, err2
		}

		//convert bytes into an index page
		page, err2 = NewIndexPageFromBytes(bytes)
		if err2 != nil {
			return nil, err2
		}

	} else if err1 != nil {
		return nil, err1
	}

	return page, nil
}

//insert a newly created index page into buffer, but not swapped into disk
//remember to swap it
func (se *storageEngine) InsertIndexPage(page *indexPage) error {

	//check if indexBuffer is full
	if se.iBuffer.IndexBufferIsFull() { //indexBuffer is full, evict and delete one page
		evictPage, evictErr := se.iBuffer.IndexBufferEvictIndexPage()
		if evictErr != nil {
			return evictErr
		}

		//if the evict page is modified, swap it into disk
		if evictPage.IndexPageIsModified() {
			err3 := se.SwapIndexPage(evictPage.IndexPageGetPageId())
			if err3 != nil {
				return err3
			}

		}

		//delete the evict page from buffer
		evictPage.IndexPageUnModify()
		deleteErr := se.iBuffer.IndexBufferDeleteIndexPage(evictPage.IndexPageGetPageId())
		if deleteErr != nil {
			return deleteErr
		}

	}

	//insert the page into buffer
	err4 := se.iBuffer.IndexBufferInsertIndexPage(page)
	if err4 != nil {
		return err4
	}

	return nil
}

//delete an index page according to its pageId, related page not swapped into disk
//remember to swap related pages
func (se *storageEngine) DeleteIndexPage(pageId uint32) error {
	return se.iBuffer.IndexBufferDeleteIndexPage(pageId)
}

//swap an index page into disk according to its pageId
func (se *storageEngine) SwapIndexPage(pageId uint32) error {
	//get this page from buffer
	page, err := se.iBuffer.IndexBufferFetchPage(pageId)
	if err != nil {
		return err
	}

	//convert this page into bytes
	bytes := page.IndexPageToBytes()

	//push bytes into disk
	err2 := se.iom.BytesToIndexFile(bytes, page.IndexPageGetPageId())

	if err2 == nil { //delete page from buffer
		page.IndexPageUnModify()
		err3 := se.iBuffer.IndexBufferDeleteIndexPage(page.IndexPageGetPageId())
		if err3 != nil {
			return err3
		}
	}

	return err2
}

//fetch a log page according to its pageId from the disk
//if this page is modified, remember to swap it
func (se *storageEngine) FetchLogPage(pageId uint32) (*logPage, error) {
	bytes, bytesErr := se.iom.BytesFromLogFile(pageId, DEFAULT_PAGE_SIZE)
	if bytesErr != nil {
		return nil, bytesErr
	}

	page, pageErr := NewLogPageFromBytes(bytes)
	if pageErr != nil {
		return nil, pageErr
	}

	return page, nil
}

//swap a log page into disk according to its pageId
func (se *storageEngine) SwapLogPage(page *logPage) error {

	bytes := page.LogPageToBytes()
	pos := page.LogPageGetLogPageId()

	se.iom.BytesToLogFile(bytes, pos)
	return nil
}
