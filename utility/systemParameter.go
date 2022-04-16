package utility

const (
	DEFAULT_PAGE_SIZE         int = 4096
	DEFAULT_DATA_BUFFER_SIZE  int = 500 //page number in dataBuffer
	DEFAULT_INDEX_BUFFER_SIZE int = 500 //page number in indexBuffer

	DEFAULT_DATAFILE_LOCATION string = "/Users/jorryjoestar/Documents/go/src/ZetaDB/file/data.zdb"
	//"/root/zeta/file/data.zdb"
	//"/Users/jorryjoestar/Documents/go/src/ZetaDB/file/data.zdb"
	DEFAULT_INDEXFILE_LOCATION string = "/Users/jorryjoestar/Documents/go/src/ZetaDB/file/index.zdb"
	//"/root/zeta/file/index.zdb"
	//"/Users/jorryjoestar/Documents/go/src/ZetaDB/file/index.zdb"
	DEFAULT_LOGFILE_LOCATION string = "/Users/jorryjoestar/Documents/go/src/ZetaDB/file/log.zdb"
	//"/root/zeta/file/log.zdb"
	//"/Users/jorryjoestar/Documents/go/src/ZetaDB/file/log.zdb"

	//if a tuple has a field whose length in bytes is over this value, it is invalid to generate map key for this value
	DEFAULT_TUPLE_SINGAL_FIELD_OVER_LONG_LENGTH uint16 = 30

	//if a tuple total length in bytes is over this value, it is invalid to generate map key for this value
	//total length is the sum of all singal field
	DEFAULT_TUPLE_TOTAL_OVER_LONG_LENGTH uint16 = 100

	//default administrator info
	DEFAULT_ADMINISTRATOR_NAME     string = "Woozie"
	DEFAULT_ADMINISTRATOR_USER_ID  int    = 0
	DEFAULT_ADMINISTRATOR_PASSWORD string = "4319633"

	DEFAULT_SERVER_ADDRESS = "localhost:40320"
	//"localhost:40320"
	//"153.92.210.106:40320"

	DEFAULT_REQUEST_CHANNEL_CAPACITY = 100
)

var (
	DEFAULT_KEYTABLES_SCHEMA = [17]string{
		"CREATE TABLE k_userId_userName (userId INT PRIMARY KEY, userName VARCHAR(20));",
		"CREATE TABLE k_userId_password (userId INT PRIMARY KEY, passwords VARCHAR(20));",
		"CREATE TABLE k_tableId_userId (tableId INT PRIMARY KEY, userId INT);",
		"CREATE TABLE k_assertId_userId (assertId INT PRIMARY KEY, userId INT);",
		"CREATE TABLE k_viewId_userId (viewId INT PRIMARY KEY, userId INT);",
		"CREATE TABLE k_indexId_tableId (indexId INT PRIMARY KEY, tableId INT);",
		"CREATE TABLE k_triggerId_userId (triggerId INT PRIMARY KEY, userId INT);",
		"CREATE TABLE k_psmId_userId (psmId INT PRIMARY KEY, userId INT);",
		"CREATE TABLE k_tableId_schema (tableId INT PRIMARY KEY, schema VARCHAR(255));",
		"CREATE TABLE k_table (tableId INT PRIMARY KEY, headPageId INT, tailPageId INT, lastTupleId INT, tupleNum INT);",
		"CREATE TABLE k_assert (assertId INT PRIMARY KEY, assertStmt VARCHAR(255));",
		"CREATE TABLE k_view (viewId INT PRIMARY KEY, viewStmt VARCHAR(255));",
		"CREATE TABLE k_index (indexId INT PRIMARY KEY, indexStmt VARCHAR(255), indexHeadPageId INT);",
		"CREATE TABLE k_trigger (triggerId INT PRIMARY KEY, triggerStmt VARCHAR(255));",
		"CREATE TABLE k_psm (psmId INT PRIMARY KEY, psmStmt VARCHAR(255));",
		"CREATE TABLE k_emptyDataPageSlot (pageId INT);",
		"CREATE TABLE k_emptyIndexPageSlot (pageId INT);",
	}
)
