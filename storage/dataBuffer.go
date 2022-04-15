package storage

import (
	. "ZetaDB/utility"
	"errors"
)

type dataBuffer struct {
	//page container, key is called bufferId
	buffer map[int]*DataPage

	//mapping from pageId to bufferId
	mapper map[uint32]int

	//bufferId is from 1 to DEFAULT_DATA_BUFFER_SIZE
	//record all empty slot bufferId
	bufferSlots []int

	//bufferIds that are using, for evict
	bufferIds []int

	//pointer from which the next evict begin
	evictPointer int
}

//in order to create a new dataBuffer, call this function
func NewDataBuffer() *dataBuffer {
	db := &dataBuffer{}
	db.buffer = make(map[int]*DataPage)
	db.mapper = make(map[uint32]int)
	db.evictPointer = -1

	//initialize the bufferSlots
	for i := 1; i <= DEFAULT_DATA_BUFFER_SIZE; i++ {
		db.bufferSlots = append(db.bufferSlots, i)
	}

	return db
}

//fetch a data page from data buffer by its pageId
//throw error if data page with id pageId is not in this buffer
func (db *dataBuffer) DataBufferFetchPage(pageId uint32) (*DataPage, error) {
	//throw error if data page with id pageId is not in this buffer
	if db.mapper[pageId] == 0 {
		return nil, errors.New("pageId invalid, this page is not buffered")
	}

	//fetch bufferId from mapper
	bufferId := db.mapper[pageId]

	return db.buffer[bufferId], nil
}

//insert a data page into the buffer
//throw error if data page with the same pageId is already in this buffer
//throw error if this buffer is full at present
func (db *dataBuffer) DataBufferInsertDataPage(page *DataPage) error {
	currentPageId := page.DpGetPageId()

	//throw error if data page with the same pageId is already in this buffer
	if db.mapper[currentPageId] != 0 {
		return errors.New("dataPage already buffered")
	}

	//throw error if this buffer is full at present
	if db.DataBufferIsFull() {
		return errors.New("buffer full")
	}

	//get bufferId where this page is ready to insert
	bufferId := db.bufferSlots[0]
	db.bufferSlots = db.bufferSlots[1:]

	//insert this page into buffer
	db.buffer[bufferId] = page

	//append the bufferId into bufferIds
	db.bufferIds = append(db.bufferIds, bufferId)

	//update mapper
	db.mapper[currentPageId] = bufferId

	return nil
}

//delete a data page from the buffer
//throw error if this page is not in this buffer
//throw error if this page is modified
func (db *dataBuffer) DataBufferDeleteDataPage(pageId uint32) error {
	//throw error if this page is not in this buffer
	if db.mapper[pageId] == 0 {
		return errors.New("page not buffered")
	}

	//get bufferId
	bufferId := db.mapper[pageId]

	//throw error if this page is modified
	if db.buffer[bufferId].DataPageIsModified() {
		return errors.New("page modified")
	}

	//delete this page from buffer
	delete(db.buffer, bufferId)

	//delete mapping relation from mapper
	delete(db.mapper, pageId)

	//delete bufferId from bufferIds, update evict pointer
	for i, v := range db.bufferIds {
		//find the bufferId
		if v == bufferId {
			//update evictPointer
			if len(db.bufferIds) == 1 { //after deleting, this buffer is empty
				db.evictPointer = -1
			} else if i+1 == db.evictPointer && i == len(db.bufferIds)-1 {
				//pointer points to the id that ready to delete, and the id is the last one
				db.evictPointer = 1
			} else if db.evictPointer > i+1 { //ecivtPointer is affected
				db.evictPointer = db.evictPointer - 1
			}

			if i == len(db.bufferIds)-1 { //last one of bufferIds
				db.bufferIds = db.bufferIds[:i]
			} else if i == 0 { //first one of bufferIds
				db.bufferIds = db.bufferIds[1:]
			} else { //in the middle
				db.bufferIds = append(db.bufferIds[:i], db.bufferIds[i+1:]...)
			}
			break
		}
	}

	//return bufferId to bufferSlots
	db.bufferSlots = append(db.bufferSlots, bufferId)

	return nil
}

//pick a page that is not used recently, return it, but it is not deleted from this buffer
//evict order: (unmarked,unmodified) (marked,unmodified) (unmarked, modified) (marked, modified)
//throw error if this buffer is not full
func (db *dataBuffer) DataBufferEvictDataPage() (*DataPage, error) {
	//throw error if this buffer is not full
	if !db.DataBufferIsFull() {
		return nil, errors.New("this buffer is not full")
	}

	//if this is the first evict since last time the buffer is empty, initialize evictPointer to 1
	if db.evictPointer == -1 {
		db.evictPointer = len(db.bufferIds)
	}

	//remember turn ending position
	endPointer := db.evictPointer

	//first turn, find out the first (unmarked,unmodified), unmark all witnessed (marked,unmodified)
	db.evictPointer = (db.evictPointer % len(db.bufferIds)) + 1
	for ; endPointer != db.evictPointer; db.evictPointer = (db.evictPointer % len(db.bufferIds)) + 1 {
		page := db.buffer[db.evictPointer]
		if page.DataPageIsMarked() && !page.DataPageIsModified() {
			//if current page is (marked,unmodified), unmark it
			page.UnmarkDataPage()
		} else if !page.DataPageIsMarked() && !page.DataPageIsModified() {
			//find the first (unmarked,unmodified), return it
			return page, nil
		}
	}

	//second turn, find out the first (marked,unmodified)
	db.evictPointer = (db.evictPointer % len(db.bufferIds)) + 1
	for ; endPointer != db.evictPointer; db.evictPointer = (db.evictPointer % len(db.bufferIds)) + 1 {
		page := db.buffer[db.evictPointer]
		//(marked,unmodified) are all set to (unmarked,unmodified) in the previous turn
		if !page.DataPageIsMarked() && !page.DataPageIsModified() {
			return page, nil
		}
	}

	//third turn, find out the first (unmarked, modified), unmark all witnessed (marked, modified)
	db.evictPointer = (db.evictPointer % len(db.bufferIds)) + 1
	for ; endPointer != db.evictPointer; db.evictPointer = (db.evictPointer % len(db.bufferIds)) + 1 {
		page := db.buffer[db.evictPointer]
		if page.DataPageIsMarked() {
			//if current page is (marked,modified), unmark it
			page.UnmarkDataPage()
		} else if !page.DataPageIsMarked() {
			//find the first(unmarked,modified), return it
			return page, nil
		}
	}

	//forth turn, find out the first (marked, modified)
	db.evictPointer = (db.evictPointer % len(db.bufferIds)) + 1
	for ; endPointer != db.evictPointer; db.evictPointer = (db.evictPointer % len(db.bufferIds)) + 1 {
		page := db.buffer[db.evictPointer]
		//(marked,modified) are all set to (unmarked,modified) in the previous turn
		if !page.DataPageIsMarked() {
			return page, nil
		}
	}

	return nil, errors.New("evict error")
}

//test if buffer is full
func (db *dataBuffer) DataBufferIsFull() bool {
	return len(db.bufferSlots) == 0
}

//test if buffer is empty
func (db *dataBuffer) DataBufferIsEmpty() bool {
	return len(db.bufferSlots) == DEFAULT_DATA_BUFFER_SIZE
}
