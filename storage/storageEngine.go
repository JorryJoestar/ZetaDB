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
	-key tables (17)
*/

type StorageEngine struct {
	//stores 17 head pages
	keyTableHeadPageBuffer [17]*DataPage

	//buffer for data page
	dBuffer *dataBuffer

	//buffer for index page
	iBuffer *indexBuffer

	//IOManipulator
	iom *IOManipulator
}

//use GetStorageEngine to get the unique StorageEngine
var seInstance *StorageEngine
var seOnce sync.Once

func GetStorageEngine() *StorageEngine {
	iom, _ := GetIOManipulator()

	seOnce.Do(func() {
		seInstance = &StorageEngine{
			dBuffer: NewDataBuffer(),
			iBuffer: NewIndexBuffer(),
			iom:     iom}
	})
	return seInstance
}

//get a data page according to its pageId
//if this page is not buffered, fetch it from disk and push it into buffer
//if buffer is full, evict a page
//if this page is modified, remember to swap it
func (se *StorageEngine) GetDataPage(pageId uint32, schema *Schema) (*DataPage, error) {
	if pageId <= 16 { // fetch page from keyTableHeadPageBuffer
		if se.keyTableHeadPageBuffer[pageId] == nil { //fetch it from disk
			bytes, err := se.iom.BytesFromDataFile(pageId*uint32(DEFAULT_PAGE_SIZE), DEFAULT_PAGE_SIZE)
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

		if err == nil {
			return page, nil
		} else if err.Error() == "pageId invalid, this page is not buffered" { // fetch page from disk
			bytes, err := se.iom.BytesFromDataFile(pageId*uint32(DEFAULT_PAGE_SIZE), DEFAULT_PAGE_SIZE)
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
		} else {
			return nil, err
		}
	}
}

//insert a newly created data page into dataBuffer/keyTableHeadPageBuffer, but not swapped into disk
//remember to swap it
func (se *StorageEngine) InsertDataPage(page *DataPage) error {
	if page.DpGetPageId() <= 16 { // insert into keyTableHeadPageBuffer
		se.keyTableHeadPageBuffer[page.DpGetPageId()] = page
	} else { // insert into dataBuffer
		//check if dataBuffer is full
		if se.dBuffer.DataBufferIsFull() { //dataBuffer is full, evict and delete one page
			evictPage, evictErr := se.dBuffer.DataBufferEvictDataPage()
			if evictErr != nil {
				return evictErr
			}

			//if the evict page is modified, swap it into disk
			if evictPage.DataPageIsModified() { //TODO swapDataPage to SwapDataPage
				err3 := se.swapDataPage(evictPage.DpGetPageId())
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

//swap a data page into disk according to its pageId
func (se *StorageEngine) swapDataPage(pageId uint32) error {

	fmt.Printf("swap page %v\n", pageId)

	//get this page
	var page *DataPage
	var err error
	if pageId <= 16 { //pageId <= 16, fetch page from keyTableHeadPageBuffer
		page = se.keyTableHeadPageBuffer[pageId]
		if page == nil { //throw error if this page is not buffered
			return errors.New("key page not buffered")
		}
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
	err3 := se.iom.BytesToDataFile(bytes, page.DpGetPageId()*uint32(DEFAULT_PAGE_SIZE))

	//set page to unmodified
	page.UnmodifyDataPage()

	return err3
}

//throw error if dataBytes length invalid
func (se *StorageEngine) dataPageBytesToDataFile(dataBytes []byte, pageId uint32) error {
	//throw error if dataBytes length invalid
	if len(dataBytes) != DEFAULT_PAGE_SIZE {
		return errors.New("storage/storageEngine.go    dataPageBytesToDataFile() dataBytes length invalid")
	}

	//push bytes into disk
	err := se.iom.BytesToDataFile(dataBytes, pageId*uint32(DEFAULT_PAGE_SIZE))
	if err != nil {
		return err
	}

	return nil
}

//get an index page according to its pageId from index buffer
//if this page is not buffered, fetch it from disk and push it into buffer
//if buffer is full, evict a page
//if this page is modified, remember to swap it
func (se *StorageEngine) GetIndexPage(pageId uint32) (*IndexPage, error) {
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
func (se *StorageEngine) InsertIndexPage(page *IndexPage) error {

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

//throw error if indexBytes length invalid
func (se *StorageEngine) indexPageBytesToIndexFile(indexBytes []byte, pageId uint32) error {
	//throw error if dataBytes length invalid
	if len(indexBytes) != DEFAULT_PAGE_SIZE {
		return errors.New("storage/storageEngine.go    indexPageBytesToDataFile() indexBytes length invalid")
	}

	//push bytes into disk
	err := se.iom.BytesToIndexFile(indexBytes, pageId*uint32(DEFAULT_PAGE_SIZE))
	if err != nil {
		return err
	}

	return nil
}

//swap an index page into disk according to its pageId
func (se *StorageEngine) SwapIndexPage(pageId uint32) error {
	//get this page from buffer
	page, err := se.iBuffer.IndexBufferFetchPage(pageId)
	if err != nil {
		return err
	}

	//convert this page into bytes
	bytes := page.IndexPageToBytes()

	//push bytes into disk
	err2 := se.iom.BytesToIndexFile(bytes, page.IndexPageGetPageId()*uint32(DEFAULT_PAGE_SIZE))

	return err2
}

//fetch a log page according to its pageId from the disk
//if this page is modified, remember to swap it
func (se *StorageEngine) FetchLogPage(pageId uint32) (*logPage, error) {
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
func (se *StorageEngine) SwapLogPage(page *logPage) error {

	bytes := page.LogPageToBytes()
	pos := page.LogPageGetLogPageId()

	se.iom.BytesToLogFile(bytes, pos*uint32(DEFAULT_PAGE_SIZE))
	return nil
}

//set the first page of logFile
//log flag true: first byte is 0b11111111
//log flag false: first byte is 0b00000000
func (se *StorageEngine) setLogFlag(flag bool) {
	var bytes []byte
	var trueByte byte = 0b11111111
	var falseByte byte = 0b00000000

	if flag {
		bytes = append(bytes, trueByte)
	} else {
		bytes = append(bytes, falseByte)
	}

	//padding bytes
	for i := 0; i < DEFAULT_PAGE_SIZE-1; i++ {
		bytes = append(bytes, 0)
	}

	se.swapPageIntoLogFile(bytes, 0)
}

//TODO
func (se *StorageEngine) SetLogFlag(flag bool) {
	se.setLogFlag(flag)
}

//check log flag
func (se *StorageEngine) getLogFlag() (bool, error) {
	var trueByte byte = 0b11111111
	returnBytes, err := se.getPageBytesFromLogFile(0)
	if err != nil {
		return false, err
	}

	if returnBytes[0] == trueByte {
		return true, nil
	} else {
		return false, nil
	}

}

//push pageBytes into pos position in the log file
//throw error if input byte slice length invalid
func (se *StorageEngine) swapPageIntoLogFile(pageBytes []byte, pos uint32) error {
	//push byte slice into pos position in the log file
	if len(pageBytes) != DEFAULT_PAGE_SIZE {
		return errors.New("storage/storageEngine.go    swapPageIntoLogFile() pageBytes length invalid")
	}

	se.iom.BytesToLogFile(pageBytes, pos*uint32(DEFAULT_PAGE_SIZE))

	return nil
}

//get DEFAULT_PAGE_SIZE bytes from pos page
func (se *StorageEngine) getPageBytesFromLogFile(tableId uint32) ([]byte, error) {
	returnBytes, err := se.iom.BytesFromLogFile(tableId*uint32(DEFAULT_PAGE_SIZE), DEFAULT_PAGE_SIZE)
	if err != nil {
		return nil, err
	}
	return returnBytes, nil
}

//erase data file
func (se *StorageEngine) EraseDataFile() error {
	return se.iom.EmptyDataFile()
}

//erase index file
func (se *StorageEngine) EraseIndexFile() error {
	return se.iom.EmptyIndexFile()
}

//erase log file
func (se *StorageEngine) EraseLogFile() error {
	return se.iom.EmptyLogFile()
}
