package storage

import (
	. "ZetaDB/utility"
	"errors"
)

type dataBuffer struct {
	//page container, key is called bufferId
	buffer map[int]*dataPage

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
	db.buffer = make(map[int]*dataPage)
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
func (db *dataBuffer) DataBufferFetchPage(pageId uint32) (*dataPage, error) {
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
func (db *dataBuffer) DataBufferInsertDataPage(page *dataPage) error {
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

	//delete bufferId from bufferIds
	for i, v := range db.bufferIds {
		//find the bufferId
		if v == bufferId {
			if i == len(db.bufferIds)-1 {
				db.bufferIds = db.bufferIds[:i]
			} else if i == 0 {
				db.bufferIds = db.bufferIds[1:]
			} else {
				db.bufferIds = append(db.bufferIds[:i], db.bufferIds[i+1:]...)
			}
		}
	}

	//return bufferId to bufferSlots
	db.bufferSlots = append(db.bufferSlots, bufferId)

	return nil
}

//pick a page that is not used recently, return it, but it is not deleted from this buffer
//evict order: (unmarked,unmodified) (marked,unmodified) (unmarked, modified) (marked, modified)
//throw error if this buffer is not full
func (db *dataBuffer) DataBufferEvictDataPage() (*dataPage, error) {
	//throw error if this buffer is not full
	if !db.DataBufferIsFull() {
		return nil, errors.New("this buferr is not full")
	}

	//if this is the first evict since system boost, initialize evictPointer to 0
	if db.evictPointer == -1 {
		db.evictPointer = 0
	}

	//remember turn ending position
	oldPointer := db.evictPointer

	//first turn, find out the first (unmarked,unmodified), unmark all witnessed (marked,unmodified)
	db.evictPointer = (db.evictPointer + 1) % DEFAULT_DATA_BUFFER_SIZE
	for ; oldPointer != db.evictPointer; db.evictPointer = (db.evictPointer + 1) % DEFAULT_DATA_BUFFER_SIZE {
		if db.buffer[db.evictPointer].DataPageIsMarked() && !db.buffer[db.evictPointer].DataPageIsModified() {
			//if current page is (marked,unmodified), unmark it
			db.buffer[db.evictPointer].MarkDataPage()
		} else if !db.buffer[db.evictPointer].DataPageIsMarked() && !db.buffer[db.evictPointer].DataPageIsModified() {
			//find the first (unmarked,unmodified), return it
			return db.buffer[db.evictPointer], nil
		}
	}

	//second turn, find out the first (marked,unmodified)
	db.evictPointer = (db.evictPointer + 1) % DEFAULT_DATA_BUFFER_SIZE
	for ; oldPointer != db.evictPointer; db.evictPointer = (db.evictPointer + 1) % DEFAULT_DATA_BUFFER_SIZE {
		//(marked,unmodified) are all set to (unmarked,unmodified) in the previous turn
		if !db.buffer[db.evictPointer].DataPageIsMarked() && !db.buffer[db.evictPointer].DataPageIsModified() {
			return db.buffer[db.evictPointer], nil
		}
	}

	//third turn, find out the first (unmarked, modified), unmark all witnessed (marked, modified)
	db.evictPointer = (db.evictPointer + 1) % DEFAULT_DATA_BUFFER_SIZE
	for ; oldPointer != db.evictPointer; db.evictPointer = (db.evictPointer + 1) % DEFAULT_DATA_BUFFER_SIZE {
		if db.buffer[db.evictPointer].DataPageIsMarked() {
			//if current page is (marked,modified), unmark it
			db.buffer[db.evictPointer].MarkDataPage()
		} else if !db.buffer[db.evictPointer].DataPageIsMarked() {
			//find the first(unmarked,modified), return it
			return db.buffer[db.evictPointer], nil
		}
	}

	//forth turn, find out the first (marked, modified)
	db.evictPointer = (db.evictPointer + 1) % DEFAULT_DATA_BUFFER_SIZE
	for ; oldPointer != db.evictPointer; db.evictPointer = (db.evictPointer + 1) % DEFAULT_DATA_BUFFER_SIZE {
		//(marked,modified) are all set to (unmarked,modified) in the previous turn
		if !db.buffer[db.evictPointer].DataPageIsMarked() {
			return db.buffer[db.evictPointer], nil
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
