package storage

import (
	. "ZetaDB/container"
	. "ZetaDB/utility"
	"errors"
)

/*
                             index page mode 1, internal node
   -------------------------------------------------------------------------------------
   |    indexPageId     |        mode        |     elementType    |    pointerNum      |
   -------------------------------------------------------------------------------------
   |                element 1                |                element 2                |
   -------------------------------------------------------------------------------------
   |                element 3                |                element 4                |
   -------------------------------------------------------------------------------------
   |                                    . . . . . .                                    |
   -------------------------------------------------------------------------------------
   |                element n-1              |                element n                |
   -------------------------------------------------------------------------------------
   |  pointerPageId 0   |  pointerPageId 1   |  pointerPageId 2   |  pointerPageId 3   |
   -------------------------------------------------------------------------------------
   |                                    . . . . . .                                    |
   -------------------------------------------------------------------------------------
   |  pointerPageId n-3 |  pointerPageId n-2 |  pointerPageId n-1 |  pointerPageId n   |
   -------------------------------------------------------------------------------------

	~indexPageId
		-uint32, 4 bytes
		-pageId in index file
		-physical address: indexPageId * DEFAULT_PAGE_SIZE

	~mode
		-uint32, 4 bytes
		-1: this is an internal node (could be a root node)

	~elementType
		-uint32, 4 bytes
		-type of elements in this page
		-mapping:        type           elementType           elementLength
		            CHAR                    1                      1
		            INT                     2                      4
		            INTEGER                 3                      4
		            SHORTINT                4                      2
		            FLOAT                   5                      4
		            REAL                    6                      4
		            DOUBLEPRECISION         7                      8
		            DATE                    8                      4
		            TIME                    9                      4

	~pointerNum
		-int32, 4 bytes
		-valid pointer number in this page at present

	~element X
		value of element X

	~pointerPageId X
		-uint32, 4 bytes
		-pointer to an index page whose value >= element X and < element X+1


                               index page mode 2, leaf node
   -------------------------------------------------------------------------------------
   |    indexPageId     |        mode        |     elementType    |     recordNum      |
   -------------------------------------------------------------------------------------
   |     prePageId      |     nextPageId     |              IndexRecord 0              |
   -------------------------------------------------------------------------------------
   |              IndexRecord 1              |              IndexRecord 2              |
   -------------------------------------------------------------------------------------
   |              IndexRecord 3              |              IndexRecord 4              |
   -------------------------------------------------------------------------------------
   |                                    . . . . . .                                    |
   -------------------------------------------------------------------------------------
   |             IndexRecord n-3             |             IndexRecord n-2             |
   -------------------------------------------------------------------------------------
   |             IndexRecord n-1             |              padding bytes              |
   -------------------------------------------------------------------------------------

	~mode
		2: this is a leaf node (could be a root node)

	~recordNum
		-int32, 4 bytes
		-record number within this leaf node

	~prePageId
		-uint32, 4 bytes
		-pageId of previous page

	~nextPageId
		-uint32, 4 bytes
		-pageId of next page

	~IndexRecord
		see container/indexRecord.go


                           index page mode 3, duplicated node
   -------------------------------------------------------------------------------------
   |    indexPageId     |        mode        |     elementType    |   dataPageIdNum    |
   -------------------------------------------------------------------------------------
   |     prePageId      |     nextPageId     |    dataPageId 0    |   dataPageId 1     |
   -------------------------------------------------------------------------------------
   |    dataPageId 2    |    dataPageId 3    |    dataPageId 4    |    dataPageId 5    |
   -------------------------------------------------------------------------------------
   |    dataPageId 6    |    dataPageId 7    |    dataPageId 8    |    dataPageId 9    |
   -------------------------------------------------------------------------------------
   |                                    . . . . . .                                    |
   -------------------------------------------------------------------------------------
   |    dataPageId N-1  |                         padding bytes                        |
   -------------------------------------------------------------------------------------

	~mode
		3: this is a duplicated node

	~dataPageIdNum
		-int32, 4 bytes
		-pageId number in this node

	~prePageId
		-uint32, 4 bytes
		-pageId of previous page that is duplicated node containing related tuples

	~nextPageId
		-uint32, 4 bytes
		-pageId of next page that is duplicated node containing related tuples

	~dataPageId
		-uint32, 4 bytes
		-pageId in data file which contains element x

*/

type IndexPage struct {
	marked   bool
	modified bool

	indexPageId uint32
	mode        uint32
	elementType uint32

	pointerNum    int32    //valid for mode1
	elements      [][]byte //valid for mode1, elements[0] invalid
	pointerPageId []uint32 //valid for mode 1

	prePageId  uint32 //valid for mode 2&3
	nextPageId uint32 //valid for mode 2&3

	recordNum int32          //valid for mode2
	records   []*IndexRecord //valid for mode2

	dataPageIdNum int32    //valid for mode 3
	dataPageIds   []uint32 //valid for mode 3
}

//create empty indexPage
//throw error if mode is not 1, 2 or 3
//throw error if elementType is not 1 to 9
func NewIndexPage(indexPageId uint32, mode uint32, elementType uint32) (*IndexPage, error) {
	//throw error if mode is not 1, 2 or 3
	if mode != 1 && mode != 2 && mode != 3 {
		return nil, errors.New("mode invalid")
	}

	//throw error if elementType is not 1 to 9
	if elementType < 1 || elementType > 9 {
		return nil, errors.New("elementType invalid")
	}

	ip := &IndexPage{
		marked:      true,
		modified:    true,
		indexPageId: indexPageId,
		mode:        mode,
		elementType: elementType}

	if ip.IndexPageGetMode() == 1 { //set pointerNum to 0
		ip.IndexPageSetPointerNum(0)
		maxPointerNum, _ := ip.IndexPageGetMaxPointerNum()
		ip.elements = make([][]byte, maxPointerNum)
		ip.pointerPageId = make([]uint32, maxPointerNum)
	} else if ip.IndexPageGetMode() == 2 { //set recordNum to 0, set prePageId and nextPageId to indexPageId itself
		ip.recordNum = 0
		ip.IndexPageSetPrePageId(ip.IndexPageGetPageId())
		ip.IndexPageSetNextPageId(ip.IndexPageGetPageId())
	} else if ip.IndexPageGetMode() == 3 { //set dataPageIdNum to 0, set prePageId and nextPageId to indexPageId itself
		ip.dataPageIdNum = 0
		ip.IndexPageSetPrePageId(ip.IndexPageGetPageId())
		ip.IndexPageSetNextPageId(ip.IndexPageGetPageId())
	}

	return ip, nil
}

//create indexPage from bytes
//throw error if bytes length invalid
//throw error if mode is not 1, 2, or 3
//throw error if elementType is not in [1,9]
func NewIndexPageFromBytes(bytes []byte) (*IndexPage, error) {
	//throw error if bytes length invalid
	if len(bytes) != DEFAULT_PAGE_SIZE {
		return nil, errors.New("bytes length invalid")
	}

	//indexPageId
	indexPageId, _ := BytesToUint32(bytes[:4])
	bytes = bytes[4:]

	//mode
	mode, _ := BytesToUint32(bytes[:4])
	bytes = bytes[4:]

	////throw error if mode is not 1, 2, or 3
	if mode != 1 && mode != 2 && mode != 3 {
		return nil, errors.New("mode invalid")
	}

	//elementType
	elementType, _ := BytesToUint32(bytes[:4])
	bytes = bytes[4:]

	//throw error if elementType is not in [1,9]
	if elementType < 0 || elementType > 9 {
		return nil, errors.New("elementType invalid")
	}

	if mode == 1 { //internal node
		//pointerNum
		pointerNum, _ := BytesToINT(bytes[:4])
		bytes = bytes[4:]

		ip := &IndexPage{
			marked:      true,
			modified:    false,
			indexPageId: indexPageId,
			mode:        mode,
			elementType: elementType,
			pointerNum:  pointerNum}

		maxPointerNum, _ := ip.IndexPageGetMaxPointerNum()

		//elements
		elementLen, _ := IndexPageGetElementLength(elementType)
		var elements [][]byte
		firstElement := make([]byte, elementLen)
		elements = append(elements, firstElement)
		var i int32
		for i = 1; i < maxPointerNum; i++ {
			elements = append(elements, bytes[:elementLen])
			bytes = bytes[elementLen:]
		}

		//pointerPageId
		var pointerPageId []uint32
		var j int32
		for j = 0; j < maxPointerNum; j++ {
			pointer, _ := BytesToUint32(bytes[:4])
			pointerPageId = append(pointerPageId, pointer)
			bytes = bytes[4:]
		}
		ip.elements = elements
		ip.pointerPageId = pointerPageId

		return ip, nil
	} else if mode == 2 { //leaf node
		//recordNum
		recordNum, _ := BytesToINT(bytes[:4])
		bytes = bytes[4:]

		//prePageId
		prePageId, _ := BytesToUint32(bytes[:4])
		bytes = bytes[4:]

		//nextPageId
		nextPageId, _ := BytesToUint32(bytes[:4])
		bytes = bytes[4:]

		//records
		var records []*IndexRecord
		var i int32
		for i = 0; i < recordNum; i++ {
			elementLen, _ := IndexPageGetElementLength(elementType)
			record, _ := NewIndexRecordFromBytes(bytes[:elementLen+5], elementLen)
			bytes = bytes[elementLen+5:]
			records = append(records, record)
		}

		ip := &IndexPage{
			marked:      true,
			modified:    false,
			indexPageId: indexPageId,
			mode:        mode,
			elementType: elementType,
			recordNum:   recordNum,
			prePageId:   prePageId,
			nextPageId:  nextPageId,
			records:     records}

		return ip, nil
	} else { //duplicated node
		//dataPageIdNum
		dataPageIdNum, _ := BytesToINT(bytes[:4])
		bytes = bytes[4:]

		//prePageId
		prePageId, _ := BytesToUint32(bytes[:4])
		bytes = bytes[4:]

		//nextPageId
		nextPageId, _ := BytesToUint32(bytes[:4])
		bytes = bytes[4:]

		//dataPageIds
		var dataPageIds []uint32
		var i int32
		for i = 0; i < dataPageIdNum; i++ {
			dataPageId, _ := BytesToUint32(bytes[:4])
			bytes = bytes[4:]
			dataPageIds = append(dataPageIds, dataPageId)
		}

		ip := &IndexPage{
			marked:        true,
			modified:      false,
			indexPageId:   indexPageId,
			mode:          mode,
			elementType:   elementType,
			dataPageIdNum: dataPageIdNum,
			prePageId:     prePageId,
			nextPageId:    nextPageId,
			dataPageIds:   dataPageIds}

		return ip, nil
	}
}

//convert this index page to byte slice, ready to push into disk
func (ip *IndexPage) IndexPageToBytes() []byte {
	if ip.IndexPageGetMode() == 1 { //internal node
		//indexPageId
		bytes := Uint32ToBytes(ip.IndexPageGetPageId())

		//mode
		bytes = append(bytes, Uint32ToBytes(ip.IndexPageGetMode())...)

		//elementType
		bytes = append(bytes, Uint32ToBytes(ip.IndexPageGetElementType())...)

		//pointerNum
		bytes = append(bytes, INTToBytes(ip.pointerNum)...)

		//elements
		maxPointerNum, _ := ip.IndexPageGetMaxPointerNum()
		var i int32
		for i = 1; i < maxPointerNum; i++ {
			bytes = append(bytes, ip.elements[i]...)
		}

		//pointerPageId
		for i = 0; i < maxPointerNum; i++ {
			bytes = append(bytes, Uint32ToBytes(ip.pointerPageId[i])...)
		}

		//paddings
		paddingLen := DEFAULT_PAGE_SIZE - len(bytes)
		for i = 0; i < int32(paddingLen); i++ {
			bytes = append(bytes, byte(0))
		}

		return bytes
	} else if ip.IndexPageGetMode() == 2 { //leaf ndoe
		//indexPageId
		bytes := Uint32ToBytes(ip.IndexPageGetPageId())

		//mode
		bytes = append(bytes, Uint32ToBytes(ip.IndexPageGetMode())...)

		//elementType
		bytes = append(bytes, Uint32ToBytes(ip.IndexPageGetElementType())...)

		//recordNum
		bytes = append(bytes, INTToBytes(ip.recordNum)...)

		//prePageId
		bytes = append(bytes, Uint32ToBytes(ip.prePageId)...)

		//nextPageId
		bytes = append(bytes, Uint32ToBytes(ip.nextPageId)...)

		//records
		var i int32
		for i = 0; i < ip.recordNum; i++ {
			bytes = append(bytes, ip.records[i].IndexRecordToBytes()...)
		}

		//padding bytes
		paddingLen := DEFAULT_PAGE_SIZE - len(bytes)
		for i = 0; i < int32(paddingLen); i++ {
			bytes = append(bytes, byte(0))
		}

		return bytes
	} else { //duplicated node
		//indexPageId
		bytes := Uint32ToBytes(ip.IndexPageGetPageId())

		//mode
		bytes = append(bytes, Uint32ToBytes(ip.IndexPageGetMode())...)

		//elementType
		bytes = append(bytes, Uint32ToBytes(ip.IndexPageGetElementType())...)

		//dataPageIdNum
		bytes = append(bytes, INTToBytes(ip.dataPageIdNum)...)

		//prePageId
		bytes = append(bytes, Uint32ToBytes(ip.prePageId)...)

		//nextPageId
		bytes = append(bytes, Uint32ToBytes(ip.nextPageId)...)

		//dataPageIds
		var i int32
		for i = 0; i < ip.dataPageIdNum; i++ {
			bytes = append(bytes, Uint32ToBytes(ip.dataPageIds[i])...)
		}

		//padding bytes
		paddingLen := DEFAULT_PAGE_SIZE - len(bytes)
		for i = 0; i < int32(paddingLen); i++ {
			bytes = append(bytes, byte(0))
		}

		return bytes

	}
}

//indexPageId getter
func (ip *IndexPage) IndexPageGetPageId() uint32 {
	ip.IndexPageMark()
	return ip.indexPageId
}

//mode getter
func (ip *IndexPage) IndexPageGetMode() uint32 {
	ip.IndexPageMark()
	return ip.mode
}

//elementType getter
func (ip *IndexPage) IndexPageGetElementType() uint32 {
	ip.IndexPageMark()
	return ip.elementType
}

//pointerNum getter
//valid for mode 1
//throw error, if mode is not 1
func (ip *IndexPage) IndexPageGetPointerNum() (int32, error) {
	//throw error, if mode is not 1
	if ip.IndexPageGetMode() != 1 {
		return 0, errors.New("mode invalid")
	}

	ip.IndexPageMark()

	return ip.pointerNum, nil
}

//pointerNum setter
//valid for mode 1
//throw error, if mode is not 1
//throw error, if pointerNum > IndexPageGetMaxPointerNum()
func (ip *IndexPage) IndexPageSetPointerNum(n int32) error {
	//throw error, if mode is not 1
	if ip.IndexPageGetMode() != 1 {
		return errors.New("mode invalid")
	}

	//throw error, if pointerNum > IndexPageGetMaxPointerNum()
	maxPointerNum, _ := ip.IndexPageGetMaxPointerNum()
	if n > maxPointerNum {
		return errors.New("pointerNum over large")
	}

	ip.IndexPageMark()
	ip.IndexPageModify()

	ip.pointerNum = n
	return nil
}

//return max pointer number this page can contain
//throw error if mode is not 1
func (ip *IndexPage) IndexPageGetMaxPointerNum() (int32, error) {
	//throw error if mode is not 1
	if ip.IndexPageGetMode() != 1 {
		return 0, errors.New("mode invalid")
	}

	elementLen, _ := IndexPageGetElementLength(ip.elementType)
	maxNum := (int32(DEFAULT_PAGE_SIZE) - 16 + elementLen) /
		(elementLen + 4)

	ip.IndexPageMark()

	return maxNum, nil

}

//element getter
//throw error if i is not in [1, IndexPageGetMaxPointerNum()-1]
//throw error if mode is not 1
func (ip *IndexPage) IndexPageGetElementAt(i int32) ([]byte, error) {
	//throw error if mode is not 1
	if ip.IndexPageGetMode() != 1 {
		return nil, errors.New("mode invalid")
	}

	//throw error if i is not in [1, maxPointerNum-1]
	maxPointerNum, _ := ip.IndexPageGetMaxPointerNum()
	if i < 1 || i > maxPointerNum-1 {
		return nil, errors.New("i value invalid")
	}

	ip.IndexPageMark()

	return ip.elements[i], nil
}

//element setter
//throw error if i is not in [1, IndexPageGetMaxPointerNum()-1]
//throw error if mode is not 1
//throw error if bytes length invalid
func (ip *IndexPage) IndexPageSetElementAt(i int32, bytes []byte) error {
	//throw error if mode is not 1
	if ip.IndexPageGetMode() != 1 {
		return errors.New("mode invalid")
	}

	//throw error if i is not in [1, maxPointerNum-1]
	maxPointerNum, _ := ip.IndexPageGetMaxPointerNum()
	if i < 1 || i > maxPointerNum-1 {
		return errors.New("i value invalid")
	}

	ip.IndexPageMark()
	ip.IndexPageModify()

	ip.elements[i] = bytes
	return nil
}

//pointerPageId getter
//throw error if mode is not 1
//throw error if i is not in [0, IndexPageGetMaxPointerNum()-1]
func (ip *IndexPage) IndexPageGetPointerPageIdAt(i int32) (uint32, error) {
	//throw error if mode is not 1
	if ip.IndexPageGetMode() != 1 {
		return 0, errors.New("mode invalid")
	}

	//throw error if i is not in [0, maxPointerNum-1]
	maxPointerNum, _ := ip.IndexPageGetMaxPointerNum()
	if i < 0 || i > maxPointerNum-1 {
		return 0, errors.New("i value invalid")
	}

	ip.IndexPageMark()

	return ip.pointerPageId[i], nil
}

//pointerPageId setter
//throw error if mode is not 1
//throw error if i is not in [0, IndexPageGetMaxPointerNum()-1]
func (ip *IndexPage) IndexPageSetPointerPageIdAt(i int32, pageId uint32) error {
	//throw error if mode is not 1
	if ip.IndexPageGetMode() != 1 {
		return errors.New("mode invalid")
	}

	//throw error if i is not in [0, maxPointerNum-1]
	maxPointerNum, _ := ip.IndexPageGetMaxPointerNum()
	if i < 0 || i > maxPointerNum-1 {
		return errors.New("i value invalid")
	}

	ip.IndexPageMark()
	ip.IndexPageModify()

	ip.pointerPageId[i] = pageId
	return nil
}

//recordNum getter
//throw error if mode is not 2
func (ip *IndexPage) IndexPageGetRecordNum() (int32, error) {
	//throw error if mode is not 2
	if ip.IndexPageGetMode() != 2 {
		return 0, errors.New("mode invalid")
	}

	ip.IndexPageMark()

	return ip.recordNum, nil
}

//return max record number this page can contain
//throw error if mode is not 2
func (ip *IndexPage) IndexPageGetMaxRecordNum() (int32, error) {
	//throw error if mode is not 2
	if ip.IndexPageGetMode() != 2 {
		return 0, errors.New("mode invalid")
	}

	ip.IndexPageMark()

	elementLen, _ := IndexPageGetElementLength(ip.elementType)
	recordLen := elementLen + 4 + 1
	maxRecordNum := (int32(DEFAULT_PAGE_SIZE) - 24) / recordLen
	return maxRecordNum, nil
}

//prePageId getter
//throw error if mode is not 2 or 3
func (ip *IndexPage) IndexPageGetPrePageId() (uint32, error) {
	//throw error if mode is not 2 or 3
	if ip.IndexPageGetMode() == 1 {
		return 0, errors.New("mode invalid")
	}

	ip.IndexPageMark()

	return ip.prePageId, nil
}

//prePageId setter
//throw error if mode is not 2 or 3
func (ip *IndexPage) IndexPageSetPrePageId(pageId uint32) error {
	//throw error if mode is not 2 or 3
	if ip.IndexPageGetMode() == 1 {
		return errors.New("mode invalid")
	}

	ip.IndexPageMark()
	ip.IndexPageModify()

	ip.prePageId = pageId
	return nil
}

//nextPageId getter
//throw error if mode is not 2 or 3
func (ip *IndexPage) IndexPageGetNextPageId() (uint32, error) {
	//throw error if mode is not 2 or 3
	if ip.IndexPageGetMode() == 1 {
		return 0, errors.New("mode invalid")
	}

	ip.IndexPageMark()

	return ip.nextPageId, nil
}

//nextPageId setter
//throw error if mode is not 2 or 3
func (ip *IndexPage) IndexPageSetNextPageId(pageId uint32) error {
	//throw error if mode is not 2 or 3
	if ip.IndexPageGetMode() == 1 {
		return errors.New("mode invalid")
	}

	ip.IndexPageMark()
	ip.IndexPageModify()

	ip.nextPageId = pageId
	return nil
}

//IndexRecord getter
//throw error if mode is not 2
//throw error if i is not in [0, recordNum-1]
func (ip *IndexPage) IndexPageGetIndexRecordAt(i int32) (*IndexRecord, error) {
	//throw error if mode is not 2
	if ip.IndexPageGetMode() != 2 {
		return nil, errors.New("mode invalid")
	}

	//throw error if i is not in [0, recordNum-1]
	recordNum, _ := ip.IndexPageGetRecordNum()
	if i < 0 || i > recordNum-1 {
		return nil, errors.New("i value invalid")
	}

	ip.IndexPageMark()

	return ip.records[i], nil
}

//insert a new index record into the end of records, length of records increase by 1
//throw error if mode is not 2
//throw error if this page is full
func (ip *IndexPage) IndexPageInsertNewIndexRecord(ir *IndexRecord) error {
	//throw error if mode is not 2
	if ip.IndexPageGetMode() != 2 {
		return errors.New("mode invalid")
	}

	//throw error if this page is full
	recordNum, _ := ip.IndexPageGetRecordNum()
	maxRecordNum, _ := ip.IndexPageGetMaxRecordNum()
	if recordNum == maxRecordNum {
		return errors.New("this page is full")
	}

	ip.records = append(ip.records, ir)
	ip.recordNum++
	return nil
}

//IndexRecord setter
//throw error if mode is not 2
//throw error if i is not in [0, recordNum-1]
func (ip *IndexPage) IndexPageSetIndexRecordAt(i int32, ir *IndexRecord) error {
	//throw error if mode is not 2
	if ip.IndexPageGetMode() != 2 {
		return errors.New("mode invalid")
	}

	//throw error if i is not in [0, recordNum-1]
	recordNum, _ := ip.IndexPageGetRecordNum()
	if i < 0 || i > recordNum-1 {
		return errors.New("i value invalid")
	}

	ip.IndexPageMark()
	ip.IndexPageModify()

	ip.records[i] = ir
	return nil
}

//dataPageIdNum getter
//throw error if mode is not 3
func (ip *IndexPage) IndexPageGetDataPageIdNum() (int32, error) {
	//throw error if mode is not 3
	if ip.IndexPageGetMode() != 3 {
		return 0, errors.New("mode invalid")
	}

	ip.IndexPageMark()

	return ip.dataPageIdNum, nil
}

//return max dataPageId number this page can contain
//throw error if mode is not 3
func (ip *IndexPage) IndexPageGetMaxDataPageIdNum() (int32, error) {
	//throw error if mode is not 3
	if ip.IndexPageGetMode() != 3 {
		return 0, errors.New("mode invalid")
	}

	ip.IndexPageMark()

	maxNum := (int32(DEFAULT_PAGE_SIZE) - 24) / 4
	return maxNum, nil
}

//dataPageId getter
//throw error if mode is not 3
//throw error if i is not in [0, dataPageIdNum-1]
func (ip *IndexPage) IndexPageGetDataPageIdAt(i int32) (uint32, error) {
	//throw error if mode is not 3
	if ip.IndexPageGetMode() != 3 {
		return 0, errors.New("mode invalid")
	}

	//throw error if i is not in [0, dataPageIdNum-1]
	dataPageIdNum, _ := ip.IndexPageGetDataPageIdNum()
	if i < 0 || i > dataPageIdNum-1 {
		return 0, errors.New("i value invalid")
	}

	ip.IndexPageMark()

	return ip.dataPageIds[i], nil
}

//dataPageId setter
//throw error if mode is not 3
//throw error if i is not in [0, IndexPageGetMaxDataPageIdNum()-1]
func (ip *IndexPage) IndexPageSetDataPageIdAt(i int32, pageId uint32) error {
	//throw error if mode is not 3
	if ip.IndexPageGetMode() != 3 {
		return errors.New("mode invalid")
	}

	//throw error if i is not in [0, dataPageIdNum-1]
	dataPageIdNum, _ := ip.IndexPageGetDataPageIdNum()
	if i < 0 || i > dataPageIdNum-1 {
		return errors.New("i value invalid")
	}

	ip.IndexPageMark()
	ip.IndexPageModify()

	ip.dataPageIds[i] = pageId
	return nil
}

//insert a dataPageId into this page
//throw error if mode is not 3
//throw error if this page is full
func (ip *IndexPage) IndexPageInsertNewDataPageId(pageId uint32) error {
	//throw error if mode is not 3
	if ip.IndexPageGetMode() != 3 {
		return errors.New("mode invalid")
	}

	//throw error if this page is full
	dataPageIdNum, _ := ip.IndexPageGetDataPageIdNum()
	maxDataPageIdNum, _ := ip.IndexPageGetMaxDataPageIdNum()
	if dataPageIdNum == maxDataPageIdNum {
		return errors.New("this page is full")
	}

	ip.dataPageIds = append(ip.dataPageIds, pageId)
	ip.dataPageIdNum++

	return nil
}

//mark this index page
func (ip *IndexPage) IndexPageMark() {
	ip.marked = true
}

//unmark this index page
func (ip *IndexPage) IndexPageUnMark() {
	ip.marked = false
}

//marked getter
//if this page is marked, return true
func (ip *IndexPage) IndexPageIsMarked() bool {
	return ip.marked
}

//set modified to true
func (ip *IndexPage) IndexPageModify() {
	ip.modified = true
}

//set modified to false
func (ip *IndexPage) IndexPageUnModify() {
	ip.modified = false
}

//modified getter
//if this page is modified, return true
func (ip *IndexPage) IndexPageIsModified() bool {
	return ip.modified
}

//element length getter
//throw error if elementType is not in [1,9]
func IndexPageGetElementLength(elementType uint32) (int32, error) {
	//throw error if elementType is not in [1,9]
	if elementType < 1 || elementType > 9 {
		return 0, errors.New("elementType invalid")
	}

	if elementType == 1 {
		return 1, nil
	} else if elementType == 4 {
		return 2, nil
	} else if elementType == 7 {
		return 8, nil
	} else {
		return 4, nil
	}
}
