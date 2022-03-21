package utility

const (
	DEFAULT_PAGE_SIZE         int = 4096
	DEFAULT_DATA_BUFFER_SIZE  int = 5   //page number in dataBuffer
	DEFAULT_INDEX_BUFFER_SIZE int = 500 //page number in indexBuffer

	DEFAULT_DATAFILE_LOCATION  string = "/Users/jorryjoestar/Documents/go/src/ZetaDB/file/data.zdb"
	DEFAULT_INDEXFILE_LOCATION string = "/Users/jorryjoestar/Documents/go/src/ZetaDB/file/index.zdb"
	DEFAULT_LOGFILE_LOCATION   string = "/Users/jorryjoestar/Documents/go/src/ZetaDB/file/log.zdb"

	//if a tuple has a field whose length in bytes is over this value, it is invalid to generate map key for this value
	DEFAULT_TUPLE_SINGAL_FIELD_OVER_LONG_LENGTH uint16 = 30

	//if a tuple total length in bytes is over this value, it is invalid to generate map key for this value
	//total length is the sum of all singal field
	DEFAULT_TUPLE_TOTAL_OVER_LONG_LENGTH uint16 = 100
)
