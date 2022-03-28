package storage

import (
	. "ZetaDB/utility"
	"errors"
)

type indexBuffer struct {
	//page container, key is called bufferId
	buffer map[int]*IndexPage

	//mapping from pageId to bufferId
	mapper map[uint32]int

	//bufferId is from 1 to DEFAULT_INDEX_BUFFER_SIZE
	//record all empty slot bufferId
	bufferSlots []int

	//bufferIds that are using, for evict
	bufferIds []int

	//pointer from which the next evict begin
	evictPointer int
}

//in order to create a new indexBuffer, call this function
func NewIndexBuffer() *indexBuffer {
	ib := &indexBuffer{}
	ib.buffer = make(map[int]*IndexPage)
	ib.mapper = make(map[uint32]int)
	ib.evictPointer = -1

	//initialize the bufferSlots
	for i := 1; i <= DEFAULT_INDEX_BUFFER_SIZE; i++ {
		ib.bufferSlots = append(ib.bufferSlots, i)
	}

	return ib
}

//fetch an index page from index buffer by its pageId
//throw error if index page with id pageId is not in this buffer
func (ib *indexBuffer) IndexBufferFetchPage(pageId uint32) (*IndexPage, error) {
	//throw error if index page with id pageId is not in this buffer
	if ib.mapper[pageId] == 0 {
		return nil, errors.New("pageId invalid, this page is not buffered")
	}

	//fetch bufferId from mapper
	bufferId := ib.mapper[pageId]

	return ib.buffer[bufferId], nil
}

//insert an index page into the buffer
//throw error if index page with the same pageId is already in this buffer
//throw error if this buffer is full at present
func (ib *indexBuffer) IndexBufferInsertIndexPage(page *IndexPage) error {
	currentPageId := page.IndexPageGetPageId()

	//throw error if index page with the same pageId is already in this buffer
	if ib.mapper[currentPageId] != 0 {
		return errors.New("indexPage already buffered")
	}

	//throw error if this buffer is full at present
	if ib.IndexBufferIsFull() {
		return errors.New("buffer full")
	}

	//get bufferId where this page is ready to insert
	bufferId := ib.bufferSlots[0]
	ib.bufferSlots = ib.bufferSlots[1:]

	//insert this page into buffer
	ib.buffer[bufferId] = page

	//append the bufferId into bufferIds
	ib.bufferIds = append(ib.bufferIds, bufferId)

	//update mapper
	ib.mapper[currentPageId] = bufferId

	return nil
}

//delete an index page from the buffer
//throw error if this page is not in this buffer
//throw error if this page is modified
func (ib *indexBuffer) IndexBufferDeleteIndexPage(pageId uint32) error {

	//throw error if this page is not in this buffer
	if ib.mapper[pageId] == 0 {
		return errors.New("page not buffered")
	}

	//get bufferId
	bufferId := ib.mapper[pageId]

	//throw error if this page is modified
	if ib.buffer[bufferId].IndexPageIsModified() {
		return errors.New("page modified")
	}

	//delete this page from buffer
	delete(ib.buffer, bufferId)

	//delete mapping relation from mapper
	delete(ib.mapper, pageId)

	//delete bufferId from bufferIds, update evict pointer
	for i, v := range ib.bufferIds {
		//find the bufferId
		if v == bufferId {
			//update evictPointer
			if len(ib.bufferIds) == 1 { //after deleting, this buffer is empty
				ib.evictPointer = -1
			} else if i+1 == ib.evictPointer && i == len(ib.bufferIds)-1 {
				//pointer points to the id that ready to delete, and the id is the last one
				ib.evictPointer = 1
			} else if ib.evictPointer > i+1 { //ecivtPointer is affected
				ib.evictPointer = ib.evictPointer - 1
			}

			if i == len(ib.bufferIds)-1 { //last one of bufferIds
				ib.bufferIds = ib.bufferIds[:i]
			} else if i == 0 { //first one of bufferIds
				ib.bufferIds = ib.bufferIds[1:]
			} else { //in the middle
				ib.bufferIds = append(ib.bufferIds[:i], ib.bufferIds[i+1:]...)
			}
			break
		}
	}

	//return bufferId to bufferSlots
	ib.bufferSlots = append(ib.bufferSlots, bufferId)

	return nil
}

//pick a page that is not used recently, return it, but it is not deleted from this buffer
//evict order: (unmarked,unmodified) (marked,unmodified) (unmarked, modified) (marked, modified)
//throw error if this buffer is not full
func (ib *indexBuffer) IndexBufferEvictIndexPage() (*IndexPage, error) {
	//throw error if this buffer is not full
	if !ib.IndexBufferIsFull() {
		return nil, errors.New("this buffer is not full")
	}

	//if this is the first evict since last time the buffer is empty, initialize evictPointer to 1
	if ib.evictPointer == -1 {
		ib.evictPointer = len(ib.bufferIds)
	}

	//remember turn ending position
	endPointer := ib.evictPointer

	//first turn, find out the first (unmarked,unmodified), unmark all witnessed (marked,unmodified)
	ib.evictPointer = (ib.evictPointer % len(ib.bufferIds)) + 1
	for ; endPointer != ib.evictPointer; ib.evictPointer = (ib.evictPointer % len(ib.bufferIds)) + 1 {
		page := ib.buffer[ib.evictPointer]
		if page.IndexPageIsMarked() && !page.IndexPageIsModified() {
			//if current page is (marked,unmodified), unmark it
			page.IndexPageUnMark()
		} else if !page.IndexPageIsMarked() && !page.IndexPageIsModified() {
			//find the first (unmarked,unmodified), return it
			return page, nil
		}
	}

	//second turn, find out the first (marked,unmodified)
	ib.evictPointer = (ib.evictPointer % len(ib.bufferIds)) + 1
	for ; endPointer != ib.evictPointer; ib.evictPointer = (ib.evictPointer % len(ib.bufferIds)) + 1 {
		page := ib.buffer[ib.evictPointer]
		//(marked,unmodified) are all set to (unmarked,unmodified) in the previous turn
		if !page.IndexPageIsMarked() && !page.IndexPageIsModified() {
			return page, nil
		}
	}

	//third turn, find out the first (unmarked, modified), unmark all witnessed (marked, modified)
	ib.evictPointer = (ib.evictPointer % len(ib.bufferIds)) + 1
	for ; endPointer != ib.evictPointer; ib.evictPointer = (ib.evictPointer % len(ib.bufferIds)) + 1 {
		page := ib.buffer[ib.evictPointer]
		if page.IndexPageIsMarked() {
			//if current page is (marked,modified), unmark it
			page.IndexPageUnMark()
		} else if !page.IndexPageIsMarked() {
			//find the first(unmarked,modified), return it
			return page, nil
		}
	}

	//forth turn, find out the first (marked, modified)
	ib.evictPointer = (ib.evictPointer % len(ib.bufferIds)) + 1
	for ; endPointer != ib.evictPointer; ib.evictPointer = (ib.evictPointer % len(ib.bufferIds)) + 1 {
		page := ib.buffer[ib.evictPointer]
		//(marked,modified) are all set to (unmarked,modified) in the previous turn
		if !page.IndexPageIsMarked() {
			return page, nil
		}
	}

	return nil, errors.New("evict error")
}

//test if buffer is full
func (ib *indexBuffer) IndexBufferIsFull() bool {
	return len(ib.bufferSlots) == 0
}

//test if buffer is empty
func (ib *indexBuffer) IndexBufferIsEmpty() bool {
	return len(ib.bufferSlots) == DEFAULT_INDEX_BUFFER_SIZE
}
