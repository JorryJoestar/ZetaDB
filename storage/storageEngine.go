package storage

import (
	. "ZetaDB/container"
	. "ZetaDB/utility"
	"errors"
	"fmt"
	"sync"
)

/*
							storage engine structure
	-an ioManipulator
	-a data page buffer
	-an index page buffer
	-key table head page slice
	-key tables (16)
		- k_userId_userName(userId INT PRIMARY KEY, userName VARCHAR(20))
			head page number 0, tableId 0
		- k_userId_password(userId INT PRIMARY KEY, password VARCHAR(20))
			head page number 1, tableId 1
		- k_tableId_userId(tableId INT PRIMARY KEY, userId INT)
			head page number 2, tableId 2
		- k_assertId_userId(assertId INT PRIMARY KEY, userId INT)
			head page number 3, tableId 3
		- k_viewId_userId(viewId INT PRIMARY KEY, userId INT)
			head page number 4, tableId 4
		- k_indexId_tableId(indexId INT PRIMARY KEY, tableId INT)
			head page number 5, tableId 5
		- k_triggerId_userId(triggerId INT PRIMARY KEY, userId INT)
			head page number 6, tableId 6
		- k_psmId_userId(psmId INT PRIMARY KEY, userId INT)
			head page number 7, tableId 7
		- k_tableId_schema(tableId INT PRIMARY KEY, schema VARCHAR(255))
			head page number 8, tableId 8
		- k_table(tableId INT PRIMARY KEY, headPageId INT, lastTupleIndexId INT, tupleNum INT)
			head page number 9, tableId 9
		- k_assert(assertId INT PRIMARY KEY, assertStmt VARCHAR(255))
			head page number 10, tableId 10
		- k_view(viewId INT PRIMARY KEY, viewStmt VARCHAR(255))
			head page number 11, tableId 11
		- k_index(indexId INT PRIMARY KEY, logHeadPageId INT)
			head page number 12, tableId 12
		- k_trigger(triggerId INT PRIMARY KEY, triggerStmt VARCHAR(255))
			head page number 13, tableId 13
		- k_psm(psmId INT PRIMARY KEY, psmStmt VARCHAR(255))
			head page number 14, tableId 14
		- k_emptyPageSlot(pageId INT)
			head page number 15, tableId 15
*/

type storageEngine struct {
	//stores 16 head pages
	keyTableHeadPageBuffer [16]*dataPage

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

//get a data page according to its pageId
//if this page is modified, remember to swap it
func (se *storageEngine) GetDataPage(pageId uint32, schema *Schema) (*dataPage, error) {
	if pageId <= 15 { // fetch page from keyTableHeadPageBuffer
		if se.keyTableHeadPageBuffer[pageId] == nil { //fetch it from disk
			bytes, err := se.iom.BytesFromDataFile(pageId, DEFAULT_PAGE_SIZE)
			if err != nil {
				return nil, err
			}

			se.keyTableHeadPageBuffer[pageId], err = NewDataPageFromBytes(bytes, schema)
			if err != nil {
				return nil, err
			}
		}
		return se.keyTableHeadPageBuffer[pageId], nil
	} else { // fetch page from dataBuffer
		page, err := se.dBuffer.DataBufferFetchPage(pageId)
		if err.Error() == "pageId invalid, this page is not buffered" { // fetch page from disk
			bytes, err := se.iom.BytesFromDataFile(pageId, DEFAULT_PAGE_SIZE)
			if err != nil {
				return nil, err
			}

			page, err = NewDataPageFromBytes(bytes, schema)
			if err != nil {
				return nil, err
			}

			err = se.InsertDataPage(page)
			if err != nil {
				return nil, err
			} else {
				return page, nil
			}
		} else if err != nil {
			return nil, err
		} else {
			return page, nil
		}
	}
}

//insert a newly created data page into dataBuffer/keyTableHeadPageBuffer, but not swapped into disk
//remember to swap it
func (se *storageEngine) InsertDataPage(page *dataPage) error {
	if page.DpGetPageId() <= 15 { // insert into keyTableHeadPageBuffer
		se.keyTableHeadPageBuffer[page.DpGetPageId()] = page
	} else { // insert into dataBuffer
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
	}

	return nil
}

//delete a data page according to its pageId, related page not swapped into disk
//throw error if corresponding page is a key table head page
//remember to swap related pages
func (se *storageEngine) DeleteDataPage(pageId uint32) error {
	if pageId <= 15 { //throw error
		return errors.New("pageId invalid")
	}
	return se.dBuffer.DataBufferDeleteDataPage(pageId)
}

//swap a data page into disk according to its pageId
func (se *storageEngine) SwapDataPage(pageId uint32) error {

	//get this page
	var page *dataPage
	var err error
	if pageId <= 15 { //pageId <= 15, fetch page from keyTableHeadPageBuffer
		page = se.keyTableHeadPageBuffer[pageId]
	} else { //get this page from buffer
		page, err = se.dBuffer.DataBufferFetchPage(pageId)
		if err != nil {
			return err
		}
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
	if err1 == nil {
		return page, nil
	} else if err1.Error() == "pageId invalid, this page is not buffered" { //not in buffer

		//fetch bytes from disk
		bytes, err2 := se.iom.BytesFromIndexFile(pageId*uint32(DEFAULT_PAGE_SIZE), DEFAULT_PAGE_SIZE)
		if err2 != nil {
			return nil, err2
		}

		//convert bytes into an index page
		page, err2 = NewIndexPageFromBytes(bytes)
		if err2 != nil {
			return nil, err2
		}

		//push this page into buffer
		err2 = se.InsertIndexPage(page)
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
	err2 := se.iom.BytesToIndexFile(bytes, page.IndexPageGetPageId()*uint32(DEFAULT_PAGE_SIZE))

	fmt.Printf("swap: %v\n", pageId)

	return err2
}

//fetch a log page according to its pageId from the disk
//if this page is modified, remember to swap it
func (se *storageEngine) FetchLogPage(pageId uint32) (*logPage, error) {
	bytes, bytesErr := se.iom.BytesFromLogFile(pageId*uint32(DEFAULT_PAGE_SIZE), DEFAULT_PAGE_SIZE)
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

	se.iom.BytesToLogFile(bytes, pos*uint32(DEFAULT_PAGE_SIZE))
	return nil
}

//erase data file
func (se *storageEngine) EraseDataFile() error {
	return se.iom.EmptyDataFile()
}

//erase index file
func (se *storageEngine) EraseIndexFile() error {
	return se.iom.EmptyIndexFile()
}

//erase log file
func (se *storageEngine) EraseLogFile() error {
	return se.iom.EmptyLogFile()
}
