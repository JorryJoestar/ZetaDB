package storage

import (
	. "ZetaDB/container"
	. "ZetaDB/utility"
	"errors"
)

type dataPage struct {
	//mark used for evict policy
	//this value would not be saved into disk
	mark bool

	//if this page is modified since it is fetched from file
	//this value would not be saved into disk
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
	tupleNum int32

	//tuples in this page
	tuples []Tuple
}

//generate a new page from a byte slice
//TODO
func NewDataPageFromBytes(bytes []byte, schema *Schema) (*dataPage, error) {
	dp := &dataPage{}

	//set pageId
	pageIdBytes := bytes[:4]
	pageId, pIdErr := BytesToUint32(pageIdBytes)
	if pIdErr != nil {
		return nil, pIdErr
	}
	dp.DpSetPageId(pageId)
	bytes = bytes[4:]

	//set tableId
	tableIdBytes := bytes[:4]
	tableId, tIdErr := BytesToUint32(tableIdBytes)
	if tIdErr != nil {
		return nil, tIdErr
	}
	dp.DpSetDataTableId(tableId)
	bytes = bytes[4:]

	//set priorPageId
	priorPageIdBytes := bytes[:4]
	priorPageId, priErr := BytesToUint32(priorPageIdBytes)
	if priErr != nil {
		return nil, priErr
	}
	dp.DpSetPriorPageId(priorPageId)
	bytes = bytes[4:]

	//set nextPageId
	nextPageIdBytes := bytes[:4]
	nextPageId, npiErr := BytesToUint32(nextPageIdBytes)
	if npiErr != nil {
		return nil, npiErr
	}
	dp.DpSetNextPageId(nextPageId)
	bytes = bytes[4:]

	//set tupleNum
	tupleNumBytes := bytes[:4]
	tupleNum, tupErr := BytesToINT(tupleNumBytes)
	if tupErr != nil {
		return nil, tupErr
	}
	dp.DpSetTupleNum(tupleNum)
	bytes = bytes[4:]

	for i := 0; i < int(dp.DpGetTupleNum()); i++ {
		tupleLenBytes := bytes[:4]
		bytes = bytes[4:]

		tupleLen, lenErr := BytesToINT(tupleLenBytes)
		if lenErr != nil {
			return nil, lenErr
		}

		tupleDataBytes := bytes[:tupleLen]
		newTuple := BytesToTuple(tupleDataBytes, schema)
		dp.tuples = append(dp.tuples, newTuple)
	}

	return dp, nil
}

//generate a new page, should assign head values, but it has no tuple now
func NewDataPage(pageid uint32, tableid uint32, priorPageid uint32, nextPageid uint32) *dataPage {
	dp := &dataPage{
		pageId:      pageid,
		tableId:     tableid,
		priorPageId: priorPageid,
		nextPageId:  nextPageid,
		tupleNum:    0}
	return dp
}

//set mark to true
func (dataPage *dataPage) MarkDataPage() {
	dataPage.mark = true
}

//set mark to false
func (dataPage *dataPage) UnmarkDataPage() {
	dataPage.mark = false
}

//set modified to true
func (dataPage *dataPage) ModifyDataPage() {
	dataPage.modified = true
}

//set modified to false
func (dataPage *dataPage) UnmodifyDataPage() {
	dataPage.modified = false
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
	bytes = append(bytes, INTToBytes(dataPage.tupleNum)...)

	//tuples
	for _, tup := range dataPage.tuples {

		//push length of current tuple first
		tupSize := tup.TupleSizeInBytes()
		tupSizeBytes := INTToBytes(int32(tupSize))
		bytes = append(bytes, tupSizeBytes...)

		//push tuple data then
		bytes = append(bytes, tup.TupleToBytes()...)
	}

	return bytes
}

func (dataPage *dataPage) DpSizeInByte() int {
	size := 0

	//five fields in header, each 4 bytes
	size += 4 * 5

	//add size of each tuples within this page
	for _, tup := range dataPage.tuples {

		//4 bytes for tuple length
		size += 4

		//size of current tuple
		size += tup.TupleSizeInBytes()
	}

	return size
}

//return vacant byte number within this page
func (dataPage *dataPage) DpVacantByteNum() int {
	return DEFAULT_PAGE_SIZE - dataPage.DpSizeInByte()
}

func (dataPage *dataPage) InsertTuple(tup Tuple) error {

	//check if there is enough space to insert
	if tup.TupleSizeInBytes()+4 > dataPage.DpVacantByteNum() {
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
func (dataPage *dataPage) DpDeleteTuple(tupleId uint32) bool {

	dataPage.MarkDataPage()

	for i, tup := range dataPage.tuples {
		if tup.GetTupleId() == tupleId {
			oldTuples := dataPage.tuples
			dataPage.tuples = append(oldTuples[:i], oldTuples[i+1:]...)

			dataPage.ModifyDataPage()

			return true
		}
	}

	return false
}

//check if this page is a head page
func (dataPage *dataPage) DpIsHeadPage() bool {
	return dataPage.pageId == dataPage.priorPageId
}

//check if this page is a tail page
func (dataPage *dataPage) DpIsTailPage() bool {
	return dataPage.pageId == dataPage.nextPageId
}

//pageId getter
func (dataPage *dataPage) DpGetPageId() uint32 {

	dataPage.MarkDataPage()

	return dataPage.pageId
}

//pageId setter
func (dataPage *dataPage) DpSetPageId(pageId uint32) {

	dataPage.MarkDataPage()
	dataPage.ModifyDataPage()

	dataPage.pageId = pageId
}

//tableId getter
func (dataPage *dataPage) DpGetDataTableId() uint32 {

	dataPage.MarkDataPage()

	return dataPage.tableId
}

//tableId setter
func (dataPage *dataPage) DpSetDataTableId(tableId uint32) {

	dataPage.MarkDataPage()
	dataPage.ModifyDataPage()

	dataPage.tableId = tableId
}

//priorPageId getter
func (dataPage *dataPage) DpGetPriorPageId() uint32 {

	dataPage.MarkDataPage()

	return dataPage.priorPageId
}

//priorPageId setter
func (dataPage *dataPage) DpSetPriorPageId(priorPageId uint32) {

	dataPage.MarkDataPage()
	dataPage.ModifyDataPage()

	dataPage.priorPageId = priorPageId
}

//nextPageId getter
func (dataPage *dataPage) DpGetNextPageId() uint32 {

	dataPage.MarkDataPage()

	return dataPage.nextPageId
}

//nextPageId setter
func (dataPage *dataPage) DpSetNextPageId(nextPageId uint32) {

	dataPage.MarkDataPage()
	dataPage.ModifyDataPage()

	dataPage.nextPageId = nextPageId
}

//tupleNum getter
func (dataPage *dataPage) DpGetTupleNum() int32 {

	dataPage.MarkDataPage()

	return dataPage.tupleNum
}

//tupleNum setter
func (dataPage *dataPage) DpSetTupleNum(tupleNum int32) {

	dataPage.MarkDataPage()
	dataPage.ModifyDataPage()

	dataPage.tupleNum = tupleNum
}
