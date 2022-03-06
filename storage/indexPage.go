package storage

/*
                           index page mode 1, internal node
   -------------------------------------------------------------------------------------
   |    indexPageId     |        mode        |     elementType    |    elementNum      |
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
		1: this is an internal node (could be a root node)

	~elementType
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

	~elementNum
		valid element number in this page at present

	~element X
		value of element X

	~pointerPageId X
		pointer to an index page whose value >= element X-1 and < element X


                               index page mode 2, leaf node
   -------------------------------------------------------------------------------------
   |    indexPageId     |        mode        |     elementType    |    elementNum      |
   -------------------------------------------------------------------------------------
   |                element 0                | index/dataPageId 0 |   record type 0    |
   -------------------------------------------------------------------------------------
   |                element 1                | index/dataPageId 1 |   record type 1    |
   -------------------------------------------------------------------------------------
   |                                    . . . . . .                                    |
   -------------------------------------------------------------------------------------
   |                element N                | index/dataPageId N |   record type N    |
   -------------------------------------------------------------------------------------

	~mode
		2: this is a leaf node

	~index/dataPageId X
		pointer to a dataPage which contains tuple with element X
		or
		pointer to an indexPage mode 1, for multiple-key indexes
		or
		pointer to an indexPage mode 3, for duplicated tuples

	~record type X
		1: index/dataPageId X is pointer to an indexPage mode 1
		2: index/dataPageId X is pointer to a dataPage in data file
		3: index/dataPageId X is pointer to an indexPage mode 3


                           index page mode 3, duplicated node
   -------------------------------------------------------------------------------------
   |    indexPageId     |        mode        |     elementType    |    elementNum      |
   -------------------------------------------------------------------------------------
   |     element a      |    dataPageNum a   |    dataPageId a0   |   dataPageId a1    |
   -------------------------------------------------------------------------------------
   |   dataPageId a2    |     element b      |   dataPageNum b    |   dataPageId b0    |
   -------------------------------------------------------------------------------------
   |                                    . . . . . .                                    |
   -------------------------------------------------------------------------------------
   |     element x      |   dataPageNum x    |   dataPageId x0    |   dataPageId x1    |
   -------------------------------------------------------------------------------------

	~mode
		3: this is a duplicated node

	~dataPageNum x
		number of dataPages that contain element x

	~dataPageId xn
		pageId in data file which contains element x

*/

type indexPage struct{}
