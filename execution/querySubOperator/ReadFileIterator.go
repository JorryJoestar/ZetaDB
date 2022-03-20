package execution

import (
	"ZetaDB/container"
	"ZetaDB/storage"
	"errors"
)

//this iterator is used to fetch tuples from the disk
type ReadFileIterator struct {
	currentPageId   uint32
	currentTuple    *container.Tuple
	currentTuplesId int //slice id of currentTupe in current data page tuples[], valid for mode 0
	se              *storage.StorageEngine
	schema          *container.Schema
	hasNext         bool
}

//ReadFileIterator constructor
func NewReadFileIterator(se *storage.StorageEngine, tableHeadPageId uint32, schema *container.Schema) *ReadFileIterator {

	rfi := &ReadFileIterator{
		currentPageId:   tableHeadPageId,
		currentTuplesId: 0,
		se:              se,
		schema:          schema,
		hasNext:         true}

	return rfi
}

//iterator1 & iterator2 should always be null
//throw error if iterator1 or iterator2 is not null
//throw error if the first page is mode 2
func (rfi *ReadFileIterator) Open(iterator1 *Iterator, iterator2 *Iterator) error {
	//throw error if iterator1 or iterator2 is not null
	if iterator1 != nil || iterator2 != nil {
		return errors.New("ReadFileIterator.go    Open() parameter invalid")
	}

	firstPage, err := rfi.se.GetDataPage(rfi.currentPageId, rfi.schema)
	if err != nil {
		return err
	}

	if firstPage.DataPageMode() == 0 { //mode 0 page
		if firstPage.DpGetTupleNum() != 0 { // not an empty table
			rfi.currentTuple, err = firstPage.GetTupleAt(rfi.currentTuplesId)
			if err != nil {
				return err
			}
		} else {
			rfi.hasNext = false
			rfi.currentTuple = nil
		}

	} else if firstPage.DataPageMode() == 1 { //mode 1 page, should iterate all following mode 2 page

		var data []byte
		firstPageData, _ := firstPage.DpGetData()
		data = append(data, firstPageData...)

		firstMode2PageId, _ := firstPage.DpGetLinkNextPageId()
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

		rfi.currentTuple, err = container.NewTupleFromBytes(data, rfi.schema, firstPage.DpGetTableId())
		if err != nil {
			return err
		}

		rfi.currentTuplesId = 0

	} else { //mode2 is invalid for being a head page
		return errors.New("ReadFileIterator.go    Open() page mode invalid")
	}

	return nil
}

//throw error if HasNext() is false
func (rfi *ReadFileIterator) GetNext() (*container.Tuple, error) {
	if !rfi.HasNext() {
		return nil, errors.New("ReadFileIterator.go    GetNext() hasNext false")
	}

	//save tuple to return temperary
	tupleToReturn := rfi.currentTuple

	currentPage, err := rfi.se.GetDataPage(rfi.currentPageId, rfi.schema)
	if err != nil {
		return nil, err
	}

	if currentPage.DataPageMode() == 0 { //mode 0 page

		if rfi.currentTuplesId+1 == int(currentPage.DpGetTupleNum()) { //already iterate all tuples within this page
			isTail, _ := currentPage.DpIsTailPage()
			if isTail { //already iterate all tuples within this table
				rfi.hasNext = false
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

	return tupleToReturn, nil
}

func (rfi *ReadFileIterator) HasNext() bool {
	return rfi.hasNext
}

func (rfi *ReadFileIterator) Close() error {
	return nil
}