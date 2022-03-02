package storage

import (
	. "ZetaDB/container"
	. "ZetaDB/utility"
	"errors"
)

type dataPage struct {
	//mark used for evict policy
	mark bool

	//if this page is modified since it is fetched from file
	modified bool

	//slot number in data.zdb
	pageId uint32

	//id of table which this page belongs to
	tableId uint32

	//id of prior page
	//if priorPageId == pageId, this page is a head page
	priorPageId uint32

	//id of next page
	//if nextPageId == pageId, this page is a tail page
	nextPageId uint32

	//number of tuples in this page
	tupleNum uint32

	//tuples in this page
	tuples []Tuple
}

//generate a new page from a byte slice
//TODO
func NewDataPageFromBytes(bytes []byte) (*dataPage, error) {
	return nil, nil
}

//set mark to true
func (dataPage *dataPage) MarkDataPage() {
	dataPage.mark = true
}

//set mark to false
func (dataPage *dataPage) UnmarkDataPage() {
	dataPage.mark = false
}

//convert this dataPage into byte slice, ready to insert into file
func (dataPage *dataPage) DataPageToBytes() []byte {
	var bytes []byte

	//pageId
	bytes = append(bytes, Uint32ToBytes(dataPage.pageId)...)

	//tableId
	bytes = append(bytes, Uint32ToBytes(dataPage.tableId)...)

	//priorPageId
	bytes = append(bytes, Uint32ToBytes(dataPage.priorPageId)...)

	//nextPageId
	bytes = append(bytes, Uint32ToBytes(dataPage.nextPageId)...)

	//tupleNum
	bytes = append(bytes, Uint32ToBytes(dataPage.tupleNum)...)

	//tuples
	for _, tup := range dataPage.tuples {
		bytes = append(bytes, tup.TupleToBytes()...)
	}

	return bytes
}

func (dataPage *dataPage) SizeInByte() int {
	size := 0

	//five fields in header, each 4 bytes
	size += 4 * 5

	//add size of each tuples within this page
	for _, tup := range dataPage.tuples {
		size += tup.TupleSizeInBytes()
	}

	return size
}

//return vacant byte number within this page
func (dataPage *dataPage) VacantByteNum() int {
	return DEFAULT_DATAPAGE_SIZE - dataPage.SizeInByte()
}

func (dataPage *dataPage) InsertTuple(tup Tuple) error {

	//check if there is enough space to insert
	if tup.TupleSizeInBytes() > dataPage.VacantByteNum() {
		return errors.New("not enough space to insert this tuple into this dataPage")
	}

	//change pageId of this tuple
	tup.SetPageId(dataPage.pageId)

	//insert into tuples
	dataPage.tuples = append(dataPage.tuples, tup)

	//marked as modified
	dataPage.modified = true

	return nil
}

//delete a tuple from this page according to its tupleId
func (dataPage *dataPage) DeleteTuple(tupleId uint32) bool {
	for i, tup := range dataPage.tuples {
		if tup.GetTupleId() == tupleId {
			oldTuples := dataPage.tuples
			dataPage.tuples = append(oldTuples[:i], oldTuples[i+1:]...)
			return true
		}
	}
	return false
}

//check if this page is a head page
func (dataPage *dataPage) IsHeadPage() bool {
	return dataPage.pageId == dataPage.priorPageId
}

//check if this page is a tail page
func (dataPage *dataPage) IsTailPage() bool {
	return dataPage.pageId == dataPage.nextPageId
}

func (dataPage *dataPage) GetPageId() uint32 {
	return dataPage.pageId
}

func (dataPage *dataPage) GetTableId() uint32 {
	return dataPage.tableId
}

func (dataPage *dataPage) GetPriorPageId() uint32 {
	return dataPage.priorPageId
}

func (dataPage *dataPage) GetNextPageId() uint32 {
	return dataPage.nextPageId
}
