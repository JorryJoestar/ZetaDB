package storage

import (
	. "ZetaDB/container"
	. "ZetaDB/utility"
	"errors"
)

/*
                             data page structure (in pageMode 0)
	-------------------------------------------------------------------------------------   -
	|       pageId       |      tableId       |      pageMode      |  tupleNum/dataSize |    | header
	-------------------------------------------------------------------------------------    | part
	|    priorPageId     |     nextPageId     |   linkPrePageId    |   linkNextPageId   |    |
	-------------------------------------------------------------------------------------   -
	|      length0       |                 tuple0                  |      length1       |    |
	-------------------------------------------------------------------------------------    |
	|      tuple1        |      length2       |                  tuple2                 |    | data
	-------------------------------------------------------------------------------------    | part
	|                              tuple2                          |         ...        |    |
	-------------------------------------------------------------------------------------    |
	|      lengthN       |      tupleN        |               vacant parts              |    |
	-------------------------------------------------------------------------------------   -

	<-------------------------------------16 bytes-------------------------------------->

	~pageId
		-uint32, 4 bytes
		-unique within data.zdb
		-physical address in data.zdb is (pageId * DEFAULT_PAGE_SIZE)

	~tableId
		-uint32, 4 bytes
		-denote which table it belongs to

	~pageMode
		-uint32, 4 bytes
		-value 0:
			-mode 0
			-normal data pages, contain multiple tuples
		-value 1:
			-mode 1
			-head of a list to store a tuple larger than DEFAULT_PAGE_SIZE - header(32bytes)
		-value 2:
			-mode 2
			-non-head of a list to store a tuple larger than DEFAULT_PAGE_SIZE - header(32bytes)

	~tupleNum/dataSize
		-int32, 4 bytes
		-mode 0:
			as tupleNum, stores tuple number within this page
		-mode 1:
			invalid, because data part of mode1 page is always full, but set to DEFAULT_PAGE_SIZE - header(32bytes) for safety
		-mode 2:
			as dataSize, stores data byte number in data part

	~priorPageId
		-uint32, 4 bytes
		-invalid for mode 2
		-pageId of prior page within the page
		-if this.pageId == this.priorPageId, this page is the head page of table

	~nextPageId
		-uint32, 4 bytes
		-invalid for mode 2
		-pageId of next page within the page
		-if this.pageId == this.nextPageId, this page is the tail page of table

	~linkPrePageId
		-uint32, 4 bytes
		-invalid for mode 0
		-pageId of previous page within the list to denote a large tuple
		-if this.pageId == this.linkPrePageId, this page is the head page of list, and must be in mode 1

	~linkNextPageId
		-uint32, 4 bytes
		-invalid for mode 0
		-pageId of next page within the list to denote a large tuple
		-if this.pageId == this.linkNextPageId, this page is the tail page of list, must be in mode 1 or 2

	~data part
		-mode 0
			-lengthX
				-int32, 4 bytes
				-length of tuple X in bytes
			-tupleX
				-arbitary bytes
				-data from TupleToBytes()
		-mode 1 or 2
			-data bytes of a large tuple

	~relationship between pages

        ------------------------------------------- table --------------------------------------------->

                   priorPageId            priorPageId           priorPageId                                 |
        /-------\  <----------  /-------\ <---------- /-------\ <---------- /-------\ ... /-------\         |
    /--	| mode0 |               | mode0 |             | mode1 |             | mode0 |     | mode0 | --\     |
    |	\-------/  ---------->  \-------/ ----------> \-------/ ----------> \-------/ ... \-------/   |     |
    |       /\      nextPageId             nextPageId  /\    |  nextPageId                   /\       |     |
    \-------/                                          |     |                                \-------/     |
    priorPageId                          linkPrePageId |     |  linkNextPageId               nextPageId
                                                       |    \/                                              l
                                                      /-------\                                             i
                                                      | mode2 |                                             s
                                                      \-------/                                             t
                                           nextPageId  /\    |  nextPageId
                                                       |     |                                              |
                                         linkPrePageId |     |  linkNextPageId                              |
                                                       |    \/                                              |
                                                      /-------\                                             |
                                                      | mode2 |                                             |
                                                      \-------/                                             \/
*/

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

	//mode of this page
	pageMode uint32

	//number of tuples in this page, valid for mode 0
	tupleNum int32

	//byte number of data in this page, valid for mode 2
	dataSize int32

	//id of prior page
	//if priorPageId == pageId, this page is a head page
	//invalid for mode 2
	priorPageId uint32

	//id of next page
	//if nextPageId == pageId, this page is a tail page
	//invalid for mode 2
	nextPageId uint32

	//pageId of previous page within the list to denote a large tuple
	//invalid for mode 0
	linkPrePageId uint32

	//pageId of next page within the list to denote a large tuple
	//invalid for mode 0
	linkNextPageId uint32

	//tuples in this page, valid for mode 0
	tuples []*Tuple

	//part of data bytes of a large tuple, valid for mode 1 and 2
	data []byte
}

//generate a new page from a byte slice
func NewDataPageFromBytes(bytes []byte, schema *Schema) (*dataPage, error) {
	dp := &dataPage{}

	//set pageId
	pageIdBytes := bytes[:4]
	pageId, pIdErr := BytesToUint32(pageIdBytes)
	if pIdErr != nil {
		return nil, pIdErr
	}
	dp.pageId = pageId
	bytes = bytes[4:]

	//set tableId
	tableIdBytes := bytes[:4]
	tableId, tIdErr := BytesToUint32(tableIdBytes)
	if tIdErr != nil {
		return nil, tIdErr
	}
	dp.tableId = tableId
	bytes = bytes[4:]

	//set pageMode
	pageModeBytes := bytes[:4]
	pageMode, pmErr := BytesToUint32(pageModeBytes)
	if pmErr != nil {
		return nil, pmErr
	}
	dp.pageMode = pageMode
	bytes = bytes[4:]

	//set tupleNum/dataSize
	if dp.pageMode == 0 { //if mode = 0, set tupleNum
		tupleNumBytes := bytes[:4]
		tupleNum, tnErr := BytesToINT(tupleNumBytes)
		if tnErr != nil {
			return nil, tnErr
		}
		dp.tupleNum = tupleNum
	} else if dp.pageMode == 2 { //else if mode = 2, set dataSize
		dataSizeBytes := bytes[:4]
		dataSize, dsErr := BytesToINT(dataSizeBytes)
		if dsErr != nil {
			return nil, dsErr
		}
		dp.dataSize = dataSize
	}
	bytes = bytes[4:] //no matter which mode, delete 4 bytes

	//if mode != 2, set priorPageId
	if dp.pageMode != 2 {
		priorPageIdBytes := bytes[:4]
		priorPageId, priErr := BytesToUint32(priorPageIdBytes)
		if priErr != nil {
			return nil, priErr
		}
		dp.DpSetPriorPageId(priorPageId)
	}
	bytes = bytes[4:] //no matter which mode, delete 4 bytes

	//if mode != 2, set nextPageId
	if dp.pageMode != 2 {
		nextPageIdBytes := bytes[:4]
		nextPageId, npiErr := BytesToUint32(nextPageIdBytes)
		if npiErr != nil {
			return nil, npiErr
		}
		dp.DpSetNextPageId(nextPageId)
	}
	bytes = bytes[4:] //no matter which mode, delete 4 bytes

	//if mode != 0, set linkPrePageId
	if dp.pageMode != 0 {
		linkPrePageIdBytes := bytes[:4]
		linkPrePageId, lppErr := BytesToUint32(linkPrePageIdBytes)
		if lppErr != nil {
			return nil, lppErr
		}
		dp.linkPrePageId = linkPrePageId

	}
	bytes = bytes[4:] //no matter which mode, delete 4 bytes

	//if mode != 0, set linkNextPageId
	if dp.pageMode != 0 {
		linkNextPageIdBytes := bytes[:4]
		linkNextPageId, lnpErr := BytesToUint32(linkNextPageIdBytes)
		if lnpErr != nil {
			return nil, lnpErr
		}
		dp.linkNextPageId = linkNextPageId
	}
	bytes = bytes[4:] //no matter which mode, delete 4 bytes

	//set fields or data
	if dp.pageMode == 0 { //if mode = 0, set fields
		for i := 0; i < int(dp.DpGetTupleNum()); i++ {
			tupleLenBytes := bytes[:4]
			bytes = bytes[4:]

			tupleLen, lenErr := BytesToINT(tupleLenBytes)
			if lenErr != nil {
				return nil, lenErr
			}

			tupleDataBytes := bytes[:tupleLen]
			newTuple, ntErr := NewTupleFromBytes(tupleDataBytes, schema, dp.DpGetTableId())
			if ntErr != nil {
				return nil, ntErr
			}
			dp.tuples = append(dp.tuples, newTuple)
		}
	} else if dp.pageMode == 1 { //else if mode = 1,set remain bytes to data
		dp.data = bytes

	} else if dp.pageMode == 2 { //else if mode = 2, set dataSize bytes to data
		dp.data = bytes[:dp.dataSize]
	}

	//set mark to true
	dp.MarkDataPage()

	//set page unmodified
	dp.UnmodifyDataPage()

	return dp, nil
}

//create a new data page in mode 0, empty
func NewDataPageMode0(pageId uint32, tableId uint32, priorPageId uint32, nextPageId uint32) *dataPage {
	dp := &dataPage{
		pageId:         pageId,
		tableId:        tableId,
		pageMode:       0,
		tupleNum:       0,
		dataSize:       0,
		priorPageId:    priorPageId,
		nextPageId:     nextPageId,
		linkPrePageId:  pageId,
		linkNextPageId: pageId}

	//mark this page
	dp.MarkDataPage()

	//mark this page as modified
	dp.ModifyDataPage()

	return dp

}

//create a new data page in mode 1
func NewDataPageMode1(pageId uint32, tableId uint32, priorPageId uint32, nextPageId uint32, linkNextPageId uint32, data []byte) *dataPage {
	dp := &dataPage{
		pageId:         pageId,
		tableId:        tableId,
		pageMode:       1,
		tupleNum:       0,
		dataSize:       int32(DEFAULT_PAGE_SIZE) - 32,
		priorPageId:    priorPageId,
		nextPageId:     nextPageId,
		linkPrePageId:  pageId,
		linkNextPageId: linkNextPageId,
		data:           data}

	//mark this page
	dp.MarkDataPage()

	//mark this page as modified
	dp.ModifyDataPage()

	return dp
}

//create a new data page in mode 2
func NewDataPageMode2(pageId uint32, tableId uint32, dataSize int32, linkPrePageId uint32, linkNextPageId uint32, data []byte) *dataPage {
	dp := &dataPage{
		pageId:         pageId,
		tableId:        tableId,
		pageMode:       2,
		tupleNum:       0,
		dataSize:       dataSize,
		priorPageId:    pageId,
		nextPageId:     pageId,
		linkPrePageId:  linkPrePageId,
		linkNextPageId: linkNextPageId,
		data:           data}

	//mark this page
	dp.MarkDataPage()

	//mark this page as modified
	dp.ModifyDataPage()

	return dp
}

//convert this dataPage into byte slice, ready to insert into file
func (dataPage *dataPage) DataPageToBytes() ([]byte, error) {
	var bytes []byte

	//pageId
	bytes = append(bytes, Uint32ToBytes(dataPage.pageId)...)

	//tableId
	bytes = append(bytes, Uint32ToBytes(dataPage.tableId)...)

	//pageMode
	bytes = append(bytes, INTToBytes(int32(dataPage.pageMode))...)

	//tupleNum/dataSize
	if dataPage.pageMode == 0 { //tupleNum valid
		bytes = append(bytes, INTToBytes(dataPage.tupleNum)...)
	} else if dataPage.pageMode == 1 { //padding empty bytes
		paddingBytes := make([]byte, 4)
		bytes = append(bytes, paddingBytes...)
	} else if dataPage.pageMode == 2 { //dataSize valid
		bytes = append(bytes, INTToBytes(dataPage.dataSize)...)
	}

	//priorPageId
	bytes = append(bytes, Uint32ToBytes(dataPage.priorPageId)...)

	//nextPageId
	bytes = append(bytes, Uint32ToBytes(dataPage.nextPageId)...)

	//linkPrePageId
	bytes = append(bytes, Uint32ToBytes(dataPage.linkPrePageId)...)

	//linkNextPageId
	bytes = append(bytes, Uint32ToBytes(dataPage.linkNextPageId)...)

	//tuples or data
	if dataPage.pageMode == 0 { //tuples, need padding bytes
		for _, tup := range dataPage.tuples {

			//push length of current tuple first
			tupSize := tup.TupleSizeInBytes()
			tupSizeBytes := INTToBytes(int32(tupSize))
			bytes = append(bytes, tupSizeBytes...)

			//push tuple data then
			tupleBytes, tbErr := tup.TupleToBytes()
			if tbErr != nil {
				return nil, tbErr
			}
			bytes = append(bytes, tupleBytes...)
		}
		//padding bytes
		paddingBytes := make([]byte, DEFAULT_PAGE_SIZE-dataPage.DpSizeInByte())
		bytes = append(bytes, paddingBytes...)
	} else if dataPage.pageMode == 1 { //data
		bytes = append(bytes, dataPage.data...)
	} else if dataPage.pageMode == 2 { //data, need padding bytes
		bytes = append(bytes, dataPage.data...)
		//padding bytes
		paddingBytes := make([]byte, DEFAULT_PAGE_SIZE-dataPage.DpSizeInByte())
		bytes = append(bytes, paddingBytes...)
	}

	//mark this page
	dataPage.MarkDataPage()

	return bytes, nil
}

//return mode
func (dataPage *dataPage) DataPageMode() uint32 {

	//mark this page
	dataPage.MarkDataPage()

	return dataPage.pageMode
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

func (dataPage *dataPage) DpSizeInByte() int {
	size := 0

	//8 header fields, each 4 bytes
	size += 32

	//tuples or data
	if dataPage.pageMode == 0 { //tuples
		//add size of each tuples within this page
		for _, tup := range dataPage.tuples {

			//4 bytes for tuple length
			size += 4

			//size of current tuple
			size += tup.TupleSizeInBytes()
		}
	} else if dataPage.pageMode == 1 { //data, full page
		size = DEFAULT_PAGE_SIZE
	} else if dataPage.pageMode == 2 { //data
		size += int(dataPage.dataSize)
	}

	dataPage.MarkDataPage()

	return size
}

//return vacant byte number within this page
func (dataPage *dataPage) DpVacantByteNum() int {
	dataPage.MarkDataPage()

	return DEFAULT_PAGE_SIZE - dataPage.DpSizeInByte()
}

//insert a tuple into this page
//throw error if mode is not 0
//throw error if no enough space to store this tuple
func (dataPage *dataPage) InsertTuple(tup *Tuple) error {

	//throw error if mode is not 0
	if dataPage.pageMode != 0 {
		return errors.New("invalid page mode")
	}

	//check if there is enough space to insert
	if tup.TupleSizeInBytes()+4 > dataPage.DpVacantByteNum() {
		return errors.New("not enough space to insert this tuple into this dataPage")
	}

	//insert into tuples
	dataPage.tuples = append(dataPage.tuples, tup)

	//alter tuple number
	dataPage.tupleNum++

	//mark this page
	dataPage.MarkDataPage()

	//marked as modified
	dataPage.ModifyDataPage()

	return nil
}

//delete a tuple from this page according to its tupleId
//throw error if mode is not 0
//throw error if no corresponding tupleId within this page
func (dataPage *dataPage) DpDeleteTuple(tupleId uint32) error {

	//throw error if mode is not 0
	if dataPage.pageMode != 0 {
		return errors.New("invalid page mode")
	}

	for i, tup := range dataPage.tuples {
		if tup.TupleGetTupleId() == tupleId {

			if i == len(dataPage.tuples)-1 { //delete the last tuple
				dataPage.tuples = dataPage.tuples[:i]
			} else if i == 0 { //delete the first tuple
				dataPage.tuples = dataPage.tuples[1:]
			} else { // in the middle
				oldTuples := dataPage.tuples
				dataPage.tuples = append(oldTuples[:i], oldTuples[i+1:]...)
			}

			//tuple number - 1
			dataPage.tupleNum--

			//mark this page
			dataPage.MarkDataPage()

			//mark this page as modified
			dataPage.ModifyDataPage()
			return nil
		}
	}

	//throw error if no corresponding tupleId within this page
	return errors.New("tupleId invalid")
}

//check if this page is a head page
//throw error if mode is 2
func (dataPage *dataPage) DpIsHeadPage() (bool, error) {

	//if mode is 2, this method is invalid
	if dataPage.pageMode == 2 {
		return false, errors.New("invalid page mode")
	}

	dataPage.MarkDataPage()

	return dataPage.pageId == dataPage.priorPageId, nil
}

//check if this page is a tail page
//throw error if mode is 2
func (dataPage *dataPage) DpIsTailPage() (bool, error) {
	//if mode is 2, this method is invalid
	if dataPage.pageMode == 2 {
		return false, errors.New("invalid page mode")
	}

	dataPage.MarkDataPage()

	return dataPage.pageId == dataPage.nextPageId, nil
}

//check if this page is a list head page
//throw error if mode is 0
func (dataPage *dataPage) DpIsListHeadPage() (bool, error) {
	//throw error if mode is 0
	if dataPage.pageMode == 0 {
		return false, errors.New("invalid page mode")
	}

	dataPage.MarkDataPage()

	return dataPage.linkPrePageId == dataPage.pageId, nil
}

//check if this page is a list tail page
//throw error if mode is 0
func (dataPage *dataPage) DpIsListTailPage() (bool, error) {
	//throw error if mode is 0
	if dataPage.pageMode == 0 {
		return false, errors.New("invalid page mode")
	}

	dataPage.MarkDataPage()

	return dataPage.linkNextPageId == dataPage.pageId, nil
}

//pageId getter
func (dataPage *dataPage) DpGetPageId() uint32 {
	dataPage.MarkDataPage()

	return dataPage.pageId
}

//tableId getter
func (dataPage *dataPage) DpGetTableId() uint32 {
	dataPage.MarkDataPage()

	return dataPage.tableId
}

//priorPageId getter
//throw error if mode is 2
func (dataPage *dataPage) DpGetPriorPageId() (uint32, error) {
	//throw error if mode is 2
	if dataPage.pageMode == 2 {
		return dataPage.pageId, errors.New("invalid page mode")
	}

	dataPage.MarkDataPage()

	return dataPage.priorPageId, nil
}

//priorPageId setter
//throw error if mode is 2
func (dataPage *dataPage) DpSetPriorPageId(priorPageId uint32) error {
	//throw error if mode is 2
	if dataPage.pageMode == 2 {
		return errors.New("invalid page mode")
	}

	dataPage.MarkDataPage()
	dataPage.ModifyDataPage()

	dataPage.priorPageId = priorPageId

	return nil
}

//nextPageId getter
//throw error if mode is 2
func (dataPage *dataPage) DpGetNextPageId() (uint32, error) {
	//throw error if mode is 2
	if dataPage.pageMode == 2 {
		return dataPage.pageId, errors.New("invalid page mode")
	}

	dataPage.MarkDataPage()

	return dataPage.nextPageId, nil
}

//nextPageId setter
//throw error if mode is 2
func (dataPage *dataPage) DpSetNextPageId(nextPageId uint32) error {
	//throw error if mode is 2
	if dataPage.pageMode == 2 {
		return errors.New("invalid page mode")
	}

	dataPage.MarkDataPage()
	dataPage.ModifyDataPage()

	dataPage.nextPageId = nextPageId

	return nil
}

//linkPrePageId getter
//throw error if mode is 0
func (dataPage *dataPage) DpGetLinkPrePageId() (uint32, error) {
	//throw error if mode is 0
	if dataPage.pageMode == 0 {
		return dataPage.pageId, errors.New("invalid page mode")
	}

	dataPage.MarkDataPage()

	return dataPage.linkPrePageId, nil
}

//linkPrePageId setter
//throw error if mode is 0
func (dataPage *dataPage) DpSetLinkPrePageId(linkPrePageId uint32) error {
	//throw error if mode is 0
	if dataPage.pageMode == 0 {
		return errors.New("invalid page mode")
	}

	dataPage.MarkDataPage()
	dataPage.ModifyDataPage()

	dataPage.linkPrePageId = linkPrePageId

	return nil
}

//linkNextPageId getter
//throw error if mode is 0
func (dataPage *dataPage) DpGetLinkNextPageId() (uint32, error) {
	//throw error if mode is 0
	if dataPage.pageMode == 0 {
		return dataPage.pageId, errors.New("invalid page mode")
	}

	dataPage.MarkDataPage()

	return dataPage.linkNextPageId, nil
}

//linkNextPageId setter
//throw error if mode is 0
func (dataPage *dataPage) DpSetLinkNextPageId(linkNextPageId uint32) error {
	//throw error if mode is 0
	if dataPage.pageMode == 0 {
		return errors.New("invalid page mode")
	}

	dataPage.MarkDataPage()
	dataPage.ModifyDataPage()

	dataPage.linkNextPageId = linkNextPageId
	return nil
}

//tupleNum getter
func (dataPage *dataPage) DpGetTupleNum() int32 {

	dataPage.MarkDataPage()

	return dataPage.tupleNum
}
