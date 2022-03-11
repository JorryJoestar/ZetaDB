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

type indexPage struct {
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
func NewIndexPage(indexPageId uint32, mode uint32, elementType uint32) (*indexPage, error) {
	//throw error if mode is not 1, 2 or 3
	if mode != 1 && mode != 2 && mode != 3 {
		return nil, errors.New("mode invalid")
	}

	//throw error if elementType is not 1 to 9
	if elementType < 1 || elementType > 9 {
		return nil, errors.New("elementType invalid")
	}

	ip := &indexPage{
		marked:      false,
		modified:    false,
		indexPageId: indexPageId,
		mode:        mode,
		elementType: elementType}

	if ip.IndexPageGetMode() == 1 { //set pointerNum to 0
		ip.IndexPageSetPointerNum(0)
	} else if ip.IndexPageGetMode() == 2 { //set recordNum to 0, set prePageId and nextPageId to indexPageId itself
		ip.IndexPageSetRecordNum(0)
		ip.IndexPageSetPrePageId(ip.IndexPageGetPageId())
		ip.IndexPageSetNextPageId(ip.IndexPageGetPageId())
	} else if ip.IndexPageGetMode() == 3 { //set dataPageIdNum to 0, set prePageId and nextPageId to indexPageId itself
		ip.IndexPageSetDataPageIdNum(0)
		ip.IndexPageSetPrePageId(ip.IndexPageGetPageId())
		ip.IndexPageSetNextPageId(ip.IndexPageGetPageId())
	}

	return ip, nil
}

//create indexPage from bytes
//throw error if bytes length invalid
//TODO
func NewIndexPageFromBytes(bytes []byte) (*indexPage, error) {
	return nil, nil
}

//convert this index page to byte slice, ready to push into disk
//TODO
func (ip *indexPage) IndexPageToBytes() []byte {
	return nil
}

//indexPageId getter
func (ip *indexPage) IndexPageGetPageId() uint32 {
	return ip.indexPageId
}

//mode getter
func (ip *indexPage) IndexPageGetMode() uint32 {
	return ip.mode
}

//elementType getter
func (ip *indexPage) IndexPageGetElementType() uint32 {
	return ip.elementType
}

//pointerNum getter
//valid for mode 1
//throw error, if mode is not 1
func (ip *indexPage) IndexPageGetPointerNum() (int32, error) {
	//throw error, if mode is not 1
	if ip.IndexPageGetMode() != 1 {
		return 0, errors.New("mode invalid")
	}

	return ip.pointerNum, nil
}

//pointerNum setter
//valid for mode 1
//throw error, if mode is not 1
//throw error, if pointerNum > IndexPageGetMaxPointerNum()
func (ip *indexPage) IndexPageSetPointerNum(n int32) error {
	//throw error, if mode is not 1
	if ip.IndexPageGetMode() != 1 {
		return errors.New("mode invalid")
	}

	//throw error, if pointerNum > IndexPageGetMaxPointerNum()
	maxPointerNum, _ := ip.IndexPageGetMaxPointerNum()
	if n > maxPointerNum {
		return errors.New("pointerNum over large")
	}

	ip.pointerNum = n
	return nil
}

//return max pointer number this page can contain
//throw error if mode is not 1
func (ip *indexPage) IndexPageGetMaxPointerNum() (int32, error) {
	//throw error if mode is not 1
	if ip.IndexPageGetMode() != 1 {
		return 0, errors.New("mode invalid")
	}

	maxNum := (int32(DEFAULT_PAGE_SIZE) - 16 + ip.IndexPageGetElementLength()) /
		(ip.IndexPageGetElementLength() + 4)

	return maxNum, nil

}

//element getter
//throw error if i is not in [1, pointerNum-1]
//throw error if mode is not 1
func (ip *indexPage) IndexPageGetElementAt(i int32) ([]byte, error) {
	//throw error if mode is not 1
	if ip.IndexPageGetMode() != 1 {
		return nil, errors.New("mode invalid")
	}

	//throw error if i is not in [1, pointerNum-1]
	pointerNum, _ := ip.IndexPageGetPointerNum()
	if i < 1 || i > pointerNum-1 {
		return nil, errors.New("i value invalid")
	}

	return ip.elements[i], nil
}

//element setter
//throw error if i is not in [1, IndexPageGetMaxPointerNum()-1]
//throw error if mode is not 1
//throw error if bytes length invalid
func (ip *indexPage) IndexPageSetElementAt(i int32, bytes []byte) error {
	//throw error if mode is not 1
	if ip.IndexPageGetMode() != 1 {
		return errors.New("mode invalid")
	}

	//throw error if i is not in [1, pointerNum-1]
	pointerNum, _ := ip.IndexPageGetPointerNum()
	if i < 1 || i > pointerNum-1 {
		return errors.New("i value invalid")
	}

	ip.elements[i] = bytes
	return nil
}

//pointerPageId getter
//throw error if mode is not 1
//throw error if i is not in [0, pointerNum-1]
func (ip *indexPage) IndexPageGetPointerPageIdAt(i int32) (uint32, error) {
	//throw error if mode is not 1
	if ip.IndexPageGetMode() != 1 {
		return 0, errors.New("mode invalid")
	}

	//throw error if i is not in [0, pointerNum-1]
	if i < 0 || i > ip.pointerNum-1 {
		return 0, errors.New("i value invalid")
	}

	return ip.pointerPageId[i], nil
}

//pointerPageId setter
//throw error if mode is not 1
//throw error if i is not in [0, pointerNum-1]
func (ip *indexPage) IndexPageSetPointerPageIdAt(i int32, pageId uint32) error {
	//throw error if mode is not 1
	if ip.IndexPageGetMode() != 1 {
		return errors.New("mode invalid")
	}

	//throw error if i is not in [0, pointerNum-1]
	if i < 0 || i > ip.pointerNum-1 {
		return errors.New("i value invalid")
	}

	ip.pointerPageId[i] = pageId
	return nil
}

//recordNum getter
//throw error if mode is not 2
func (ip *indexPage) IndexPageGetRecordNum() (int32, error) {
	//throw error if mode is not 2
	if ip.IndexPageGetMode() != 2 {
		return 0, errors.New("mode invalid")
	}

	return ip.recordNum, nil
}

//recordNum setter
//throw error if mode is not 2
func (ip *indexPage) IndexPageSetRecordNum(n int32) error {
	//throw error if mode is not 2
	if ip.IndexPageGetMode() != 2 {
		return errors.New("mode invalid")
	}

	ip.recordNum = n
	return nil
}

//return max record number this page can contain
//throw error if mode is not 2
func (ip *indexPage) IndexPageGetMaxRecordNum() (int32, error) {
	//throw error if mode is not 2
	if ip.IndexPageGetMode() != 2 {
		return 0, errors.New("mode invalid")
	}

	elementLen := ip.IndexPageGetElementLength()
	recordLen := elementLen + 32 + 8
	maxRecordNum := (int32(DEFAULT_PAGE_SIZE) - 24) / recordLen
	return maxRecordNum, nil
}

//prePageId getter
//throw error if mode is not 2 or 3
func (ip *indexPage) IndexPageGetPrePageId() (uint32, error) {
	//throw error if mode is not 2 or 3
	if ip.IndexPageGetMode() == 1 {
		return 0, errors.New("mode invalid")
	}

	return ip.prePageId, nil
}

//prePageId setter
//throw error if mode is not 2 or 3
func (ip *indexPage) IndexPageSetPrePageId(pageId uint32) error {
	//throw error if mode is not 2 or 3
	if ip.IndexPageGetMode() == 1 {
		return errors.New("mode invalid")
	}

	ip.prePageId = pageId
	return nil
}

//nextPageId getter
//throw error if mode is not 2 or 3
func (ip *indexPage) IndexPageGetNextPageId() (uint32, error) {
	//throw error if mode is not 2 or 3
	if ip.IndexPageGetMode() == 1 {
		return 0, errors.New("mode invalid")
	}

	return ip.nextPageId, nil
}

//nextPageId setter
//throw error if mode is not 2 or 3
func (ip *indexPage) IndexPageSetNextPageId(pageId uint32) error {
	//throw error if mode is not 2 or 3
	if ip.IndexPageGetMode() == 1 {
		return errors.New("mode invalid")
	}

	ip.nextPageId = pageId
	return nil
}

//IndexRecord getter
//throw error if mode is not 2
//throw error if i is not in [0, recordNum-1]
func (ip *indexPage) IndexPageGetIndexRecordAt(i int32) (*IndexRecord, error) {
	//throw error if mode is not 2
	if ip.IndexPageGetMode() != 2 {
		return nil, errors.New("mode invalid")
	}

	//throw error if i is not in [0, recordNum-1]
	recordNum, _ := ip.IndexPageGetRecordNum()
	if i < 0 || i > recordNum-1 {
		return nil, errors.New("i value invalid")
	}

	return ip.records[i], nil
}

//IndexRecord setter
//throw error if mode is not 2
//throw error if i is not in [0, recordNum-1]
func (ip *indexPage) IndexPageSetIndexRecordAt(i int32, ir *IndexRecord) error {
	//throw error if mode is not 2
	if ip.IndexPageGetMode() != 2 {
		return errors.New("mode invalid")
	}

	//throw error if i is not in [0, recordNum-1]
	recordNum, _ := ip.IndexPageGetRecordNum()
	if i < 0 || i > recordNum-1 {
		return errors.New("i value invalid")
	}

	ip.records[i] = ir
	return nil
}

//dataPageIdNum getter
//throw error if mode is not 3
func (ip *indexPage) IndexPageGetDataPageIdNum() (int32, error) {
	//throw error if mode is not 3
	if ip.IndexPageGetMode() != 3 {
		return 0, errors.New("mode invalid")
	}

	return ip.dataPageIdNum, nil
}

//dataPageIdNum setter
//throw error if mode is not 3
func (ip *indexPage) IndexPageSetDataPageIdNum(num int32) error {
	//throw error if mode is not 3
	if ip.IndexPageGetMode() != 3 {
		return errors.New("mode invalid")
	}

	ip.dataPageIdNum = num
	return nil
}

//return max dataPageId number this page can contain
//throw error if mode is not 3
func (ip *indexPage) IndexPageGetMaxDataPageIdNum() (int32, error) {
	//throw error if mode is not 3
	if ip.IndexPageGetMode() != 3 {
		return 0, errors.New("mode invalid")
	}

	maxNum := (int32(DEFAULT_PAGE_SIZE) - 24) / 4
	return maxNum, nil
}

//dataPageId getter
//throw error if mode is not 3
//throw error if i is not in [0, dataPageIdNum-1]
func (ip *indexPage) IndexPageGetDataPageIdAt(i int32) (uint32, error) {
	//throw error if mode is not 3
	if ip.IndexPageGetMode() != 3 {
		return 0, errors.New("mode invalid")
	}

	//throw error if i is not in [0, dataPageIdNum-1]
	dataPageIdNum, _ := ip.IndexPageGetDataPageIdNum()
	if i < 0 || i > dataPageIdNum-1 {
		return 0, errors.New("i value invalid")
	}

	return ip.dataPageIds[i], nil
}

//dataPageId setter
//throw error if modd is not 3
//throw error if i is not in [0, IndexPageGetMaxDataPageIdNum()-1]
func (ip *indexPage) IndexPageSetDataPageIdAt(i int32, pageId uint32) error {
	//throw error if mode is not 3
	if ip.IndexPageGetMode() != 3 {
		return errors.New("mode invalid")
	}

	//throw error if i is not in [0, dataPageIdNum-1]
	dataPageIdNum, _ := ip.IndexPageGetDataPageIdNum()
	if i < 0 || i > dataPageIdNum-1 {
		return errors.New("i value invalid")
	}

	ip.dataPageIds[i] = pageId
	return nil
}

//mark this index page
func (ip *indexPage) IndexPageMark() {
	ip.marked = true
}

//unmark this index page
func (ip *indexPage) IndexPageUnMark() {
	ip.marked = false
}

//marked getter
//if this page is marked, return true
func (ip *indexPage) IndexPageIsMarked() bool {
	return ip.marked
}

//set modified to true
func (ip *indexPage) IndexPageModify() {
	ip.modified = true
}

//set modified to false
func (ip *indexPage) IndexPageUnModify() {
	ip.modified = false
}

//modified getter
//if this page is modified, return true
func (ip *indexPage) IndexPageIsModified() bool {
	return ip.modified
}

//element length getter
func (ip *indexPage) IndexPageGetElementLength() int32 {
	typeValue := ip.IndexPageGetElementType()
	if typeValue == 1 {
		return 1
	} else if typeValue == 4 {
		return 2
	} else if typeValue == 7 {
		return 8
	} else {
		return 4
	}
}
