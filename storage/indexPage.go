package storage

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
   |    indexPageId     |        mode        |     elementType    |    elementNum      |
   -------------------------------------------------------------------------------------
   |     prePageId      |     nextPageId     |                element 0                |
   -------------------------------------------------------------------------------------
   | index/dataPageId 0 |   record type 0    |                element 1                |
   -------------------------------------------------------------------------------------
   | index/dataPageId 1 |   record type 1    |                element 2                |
   -------------------------------------------------------------------------------------
   |                                    . . . . . .                                    |
   -------------------------------------------------------------------------------------
   |index/dataPageIdN-2 |   record type N-2  |                element N-1              |
   -------------------------------------------------------------------------------------
   |index/dataPageIdN-1 |   record type N-1  |              padding bytes              |
   -------------------------------------------------------------------------------------

	~mode
		2: this is a leaf node

	~prePageId
		-uint32, 4 bytes
		-pageId of previous page

	~nextPageId
		-uint32, 4 bytes
		-pageId of next page

	~index/dataPageId X
		-uint32, 4 bytes
		pointer to a dataPage which contains tuple with element X
		or
		pointer to an indexPage mode 1, for multiple-key indexes
		or
		pointer to an indexPage mode 3, for duplicated tuples

	~record type X
		-uint8, 1 byte
		1: index/dataPageId X is pointer to an indexPage mode 1
		3: index/dataPageId X is pointer to an indexPage mode 3
		4: index/dataPageId X is pointer to a dataPage in data file


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
		pageId in data file which contains element x

*/

type indexPage struct {
	marked   bool
	modified bool

	indexPageId uint32
	mode        uint32
	elementType uint32

	pointerNum    int32    //valid for mode1
	elements      [][]byte //valid for mode1&2, for mode 1, elements[0] invalid
	pointerPageId []uint32 //valid for mode 1

	elementNum       int32    //valid for mode2
	prePageId        uint32   //valid for mode 2&3
	nextPageId       uint32   //valid for mode 2&3
	indexDataPageIds []uint32 //valid for mode 2
	recordType       []uint8  //valid for mode 2

	dataPageIdNum int32    //valid for mode 3
	dataPageIds   []uint32 //valid for mode 3
}

//create indexPage from bytes
//throw error if bytes length invalid
func NewIndexPageFromBytes(bytes []byte) (*indexPage, error) {
	return nil, nil
}

//create indexPage
func NewIndexPage(indexPageId uint32, mode uint32, elementType uint32) (*indexPage, error) {
	return nil, nil
}

//convert this index page to byte slice, ready to push into disk
func (ip *indexPage) IndexPageToBytes() []byte {
	return nil
}

//search the next layer, return indexPageId of the corresponding next layer node
//throw error if this page is not an internal node
//throw error if elementValue byte length is invalid
func (ip *indexPage) IndexPageInternalSearch(elementValue []byte) (uint32, error) {
	return 0, nil
}

//search related record according to elementValue, return index/dataPageId and recordType
//throw error if this page is not a leaf node
//throw error if elementValue byte number is invalid
//throw error if no element meets requirement
func (ip *indexPage) IndexPageLeafSearch(elementValue []byte) (uint32, uint8, error) {
	return 0, 0, nil
}

//search related dataPageIds
//return nextPageId
//if there is a next duplicated page, return true, else return false
//throw error if this page is not a duplicated node
func (ip *indexPage) IndexPageDuplicatedSearch() ([]uint32, uint32, bool, error) {
	return nil, 0, false, nil
}

//indexPageId getter
func (ip *indexPage) IndexPageGetPageId() uint32 {
	return 0
}

//mode getter
func (ip *indexPage) IndexPageGetMode() uint32 {
	return 0
}

//elementType getter
func (ip *indexPage) IndexPageGetElementType() uint32 {
	return 0
}

//elementNum getter
//invalid for mode 3
func (ip *indexPage) IndexPageGetElementNum() (int32, error) {
	return 0, nil
}

//dataPageIdNum getter
//invalid for mode 1&2
func (ip *indexPage) IndexPageGetDataPageIdNum() (int32, error) {
	return 0, nil
}

//return max entry number this page can contain
//for internal node, it is max number of pointers it can have (M)
//for leaf node, it is max number of elements it can have (L)
//for duplicated node, it is max number of dataPageIds this page can hold
func (ip *indexPage) IndexPageMaxNum() (int32, error) {
	return 0, nil
}
