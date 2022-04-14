package execution

import (
	"ZetaDB/container"
	"ZetaDB/storage"
	"errors"
)

//this iterator is used to fetch tuples from the disk sequentially
type SequentialFileReaderIterator struct {
	headPageId      uint32
	currentPageId   uint32
	currentTuple    *container.Tuple
	currentTuplesId int //slice id of currentTupe in current data page tuples[], valid for mode 0
	se              *storage.StorageEngine
	schema          *container.Schema
	hasNext         bool
}

//SequentialFileReaderIterator constructor
func NewSequentialFileReaderIterator(tableHeadPageId uint32, schema *container.Schema) *SequentialFileReaderIterator {

	rfi := &SequentialFileReaderIterator{
		headPageId:      tableHeadPageId,
		currentPageId:   tableHeadPageId,
		currentTuplesId: 0,
		se:              storage.GetStorageEngine(),
		schema:          schema,
		hasNext:         true}

	return rfi
}

//iterator1 & iterator2 should always be null
//throw error if iterator1 or iterator2 is not null
//throw error if the first page is mode 2
func (rfi *SequentialFileReaderIterator) Open(iterator1 Iterator, iterator2 Iterator) error {
	//throw error if iterator1 or iterator2 is not null
	if iterator1 != nil || iterator2 != nil {
		return errors.New("ReadFileIterator.go    Open() parameter invalid")
	}

	firstPage, err := rfi.se.GetDataPage(rfi.currentPageId, rfi.schema)
	if err != nil {
		return err
	}

	if firstPage.DpGetTupleNum() != 0 {
		rfi.currentTuple, err = firstPage.GetTupleAt(rfi.currentTuplesId)
		if err != nil {
			return err
		}
	} else {
		isTailPage, _ := firstPage.DpIsTailPage()
		if isTailPage { //only one empty page in this table
			rfi.hasNext = false
			rfi.currentTuple = nil
		} else { //check next page
			secondPageId, _ := firstPage.DpGetNextPageId()
			nextPage, _ := rfi.se.GetDataPage(secondPageId, rfi.schema)
			if nextPage.DataPageMode() == 0 {
				if nextPage.DpGetTupleNum() != 0 {
					rfi.currentTuple, err = nextPage.GetTupleAt(rfi.currentTuplesId)
					if err != nil {
						return err
					}
				} else {
					rfi.hasNext = false
					rfi.currentTuple = nil
				}
			} else {
				var data []byte
				firstPageData, _ := nextPage.DpGetData()
				data = append(data, firstPageData...)

				firstMode2PageId, _ := nextPage.DpGetLinkNextPageId()
				mode2Page, err := rfi.se.GetDataPage(firstMode2PageId, rfi.schema)
				if err != nil {
					return err
				}

				for {
					mode2PageData, _ := mode2Page.DpGetData()
					data = append(data, mode2PageData...)

					isListTail, _ := mode2Page.DpIsListTailPage()
					if isListTail { //reach the list tail
						break
					}
				}

				rfi.currentPageId = secondPageId
				rfi.currentTuple, err = container.NewTupleFromBytes(data, rfi.schema, firstPage.DpGetTableId())
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

//throw error if HasNext() is false
func (rfi *SequentialFileReaderIterator) GetNext() (*container.Tuple, error) {
	if !rfi.HasNext() {
		return nil, errors.New("ReadFileIterator.go    GetNext() hasNext false")
	}

	//save tuple to return temperary
	tupleToReturn := rfi.currentTuple

	currentPage, err := rfi.se.GetDataPage(rfi.currentPageId, rfi.schema)
	if err != nil {
		return nil, err
	}

	//update currentPageId, currentTuplesId
	//if no more tuples, set hasNext to false
	if currentPage.DataPageMode() == 0 { //mode 0 page

		if rfi.currentTuplesId+1 == int(currentPage.DpGetTupleNum()) { //already iterate all tuples within this page
			isTail, _ := currentPage.DpIsTailPage()
			if isTail { //already iterate all tuples within this table
				rfi.hasNext = false
			} else { //update currentPageId & currentTuplesId, ready for next iterate turn
				rfi.currentPageId, _ = currentPage.DpGetNextPageId()
				rfi.currentTuplesId = 0
			}

		} else { // update currentTuplesId for next iterate turn
			rfi.currentTuplesId++
		}

	} else if currentPage.DataPageMode() == 1 { //mode 1 page

		isTail, _ := currentPage.DpIsTailPage()
		if isTail { //already iterate all tuples within this table
			rfi.hasNext = false
		} else { //update currentPageId for next iterate turn
			rfi.currentPageId, _ = currentPage.DpGetNextPageId()
			rfi.currentTuplesId = 0
		}

	} else {
		return nil, errors.New("ReadFileIterator.go    GetNext() page mode invalid")
	}

	//update currentTuple
	currentPage, err = rfi.se.GetDataPage(rfi.currentPageId, rfi.schema)
	if err != nil {
		return nil, err
	}
	if currentPage.DataPageMode() == 0 {

		if currentPage.DpGetTupleNum() != 0 { // not an empty table
			rfi.currentTuple, err = currentPage.GetTupleAt(rfi.currentTuplesId)
			if err != nil {
				return nil, err
			}
		} else {
			rfi.hasNext = false
			rfi.currentTuple = nil
		}

	} else if currentPage.DataPageMode() == 1 {
		var data []byte
		firstPageData, _ := currentPage.DpGetData()
		data = append(data, firstPageData...)

		firstMode2PageId, _ := currentPage.DpGetLinkNextPageId()
		mode2Page, err := rfi.se.GetDataPage(firstMode2PageId, rfi.schema)
		if err != nil {
			return nil, err
		}

		for {
			mode2PageData, _ := mode2Page.DpGetData()
			data = append(data, mode2PageData...)

			isListTail, _ := mode2Page.DpIsListTailPage()
			if isListTail { //reach the list tail
				break
			}
		}

		rfi.currentTuple, err = container.NewTupleFromBytes(data, rfi.schema, currentPage.DpGetTableId())
		if err != nil {
			return nil, err
		}
	} else { //mode 2, throw error
		return nil, errors.New("ReadFileIterator.go    GetNext() page mode invalid")
	}

	return tupleToReturn, nil
}

//if HasNext returns false, it is invalid to call GetNext()
func (rfi *SequentialFileReaderIterator) HasNext() bool {
	return rfi.hasNext
}

//initialize this SequentialFileReaderIterator
func (rfi *SequentialFileReaderIterator) Close() {
	rfi.currentPageId = rfi.headPageId
	rfi.currentTuple = nil
	rfi.currentTuplesId = 0
	rfi.hasNext = true
}

func (rfi *SequentialFileReaderIterator) GetSchema() *container.Schema {
	return rfi.schema
}
