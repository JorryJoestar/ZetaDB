package storage

import (
	. "ZetaDB/container"
	. "ZetaDB/utility"
)

type dataBuffer struct {

	//page container, key is called bufferId
	buffer map[uint32]*dataPage

	//mapping from pageId to bufferId
	mapper map[uint32]uint32

	//bufferId is from 1 to DEFAULT_DATA_BUFFER_SIZE
	//record all empty slot bufferId
	bufferSlots []uint32
}

//in order to create a new dataBuffer, call this function
func NewDataBuffer() *dataBuffer {
	db := &dataBuffer{}
	db.buffer = make(map[int]*dataPage)
	db.mapper = make(map[uint32]int)
	db.dataBufferSize = DEFAULT_DATA_BUFFER_SIZE
	db.dataPageSize = DEFAULT_PAGE_SIZE

	//initialize the bufferSlots
	for i := 1; i <= db.dataBufferSize; i++ {
		db.bufferSlots = append(db.bufferSlots, i)
	}

	return db
}

//fetch a data page from data buffer by its pageId
//throw error if pageId invalid
//throw error if data page with id pageId is not in this buffer
func (db *dataBuffer) DataBufferFetchPage(pageId uint32) (*dataPage, error) {

	//fetch bufferId from mapper
	bufferId, err1 := db.PageIdToBufferId(pageId)

	if err1 != nil { //this page is not in buffer, should fetch it from disk
		pageBytes, fileErr := ioM.BytesFromDataFile(pageId, db.dataPageSize)
		if fileErr != nil {
			return nil, fileErr
		}
		newPage, newPageError := NewDataPageFromBytes(pageBytes, schema)
		if newPageError != nil {
			return nil, newPageError
		}
		//db.InsertDataPage(newPage)
		return newPage, nil
	}

	//this page is in buffer
	dataPage, err2 := db.GetDataPageByBufferId(bufferId)
	//mark this page
	dataPage.MarkDataPage()

	return dataPage, err2
}

//insert a data page into the buffer
//throw error if data page with the same pageId is already in this buffer
//throw error if this buffer is full at present
func (db *dataBuffer) DataBufferInsertDataPage(*dataPage) error {
	return nil
}

//delete a data page from the buffer
//throw error if this page is not in this buffer
//throw error if this page is modified
func (db *dataBuffer) DataBufferDeleteDataPage(pageId uint32) error {
	return nil
}

//pick a page that is not used recently, return it, but it is not deleted from this buffer
//evict order: (unmarked,unmodified) (marked,unmodified) (unmarked, modified) (marked, modified)
//throw error if this buffer is empty at present
func (db *dataBuffer) DataBufferEvictDataPage() (*dataPage, error) {
	return 0
}

//test if buffer is full
func (db *dataBuffer) DataBufferIsFull() bool {
	return len(db.bufferSlots) == 0
}

//test if buffer is empty
func (db *dataBuffer) DataBufferIsEmpty() bool {
	return len(db.bufferSlots) == db.dataBufferSize
}
