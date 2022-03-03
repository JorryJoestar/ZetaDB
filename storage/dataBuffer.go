package storage

import (
	. "ZetaDB/container"
	. "ZetaDB/utility"
	"errors"
)

type dataBuffer struct {

	//page container, key is bufferId
	buffer map[int]*dataPage

	//mapping from file pageId to bufferId
	mapper map[uint32]int

	//bufferId is from 1 to dataBufferSize
	//record all empty slot bufferId
	bufferSlots []int

	//fetch from systemParameter.go when this buffer is initiallized
	dataBufferSize int

	//fetch from systemParameter.go when this buffer is initiallized
	dataPageSize int
}

//in order to fetch a dataBuffer, call this function
func GetDataBuffer() *dataBuffer {
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

//fetch a data page from data buffer by its pageId, if it is not in buffer, fetch it from disk and buffer it
//TODO
func (db *dataBuffer) GetDataPageByPageId(pageId uint32, schema *Schema, ioM *IOManipulator) (*dataPage, error) {

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

//fetch a data page from data buffer by its bufferId
func (db *dataBuffer) GetDataPageByBufferId(bufferId int) (*dataPage, error) {

	page, err := db.buffer[bufferId]

	if !err { //can not find corresponding page from buffer
		return nil, errors.New("bufferId invalid")
	}

	//mark this page
	page.MarkDataPage()

	return page, nil
}

//insert a data page into the buffer, notice the page is newly created so it need to be pushed into disk
func (db *dataBuffer) InsertDataPage(*dataPage) error {
	return nil
}

//delete a data page from the buffer, also delete it from the disk
func (db *dataBuffer) DeleteDataPage() error {
	return nil
}

//test if buffer is full
func (db *dataBuffer) BufferIsFull() bool {
	return len(db.bufferSlots) == 0
}

//test if buffer is empty
func (db *dataBuffer) BufferIsEmpty() bool {
	return len(db.bufferSlots) == db.dataBufferSize
}

//TODO
//pick a page that is not used recently, if it is modified, push it into the disk, empty the slot and return bufferId
func (db *dataBuffer) EvictDataPage() int {
	return 0
}

//convert pageId to bufferId
//if no such pageId in mapper, throw error
func (db *dataBuffer) PageIdToBufferId(pageId uint32) (int, error) {
	bufferId := db.mapper[pageId]

	if bufferId == 0 { //no corresponding bufferId mapping from this pageId
		return 0, errors.New("pageId invalid")
	}

	return bufferId, nil
}

//convert bufferId to pageId
//if no such bufferId in mapper, throw error
func (db *dataBuffer) BufferIdToPageId(bufferId int) (uint32, error) {
	p := db.buffer[bufferId]

	if p == nil { //empty slot from bufferId
		return 0, errors.New("bufferId invalid")
	}

	//get pageId from this dataPage
	pageId := p.DpGetPageId()
	return pageId, nil
}
