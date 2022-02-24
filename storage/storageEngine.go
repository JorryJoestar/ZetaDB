package storage

import "sync"

type storageEngine struct {
	dataFileLocation string
	logFileLocation  string
}

//use GetStorageEngine to get the unique storageEngine
var seInstance *storageEngine
var seOnce sync.Once

func GetStorageEngine() *storageEngine {
	seOnce.Do(func() {
		seInstance = &storageEngine{}
	})
	return seInstance
}

//fetch data page from disk
func FetchDataPage(pageId uint32) *dataPage {
	dp := &dataPage{}
	return dp
}

//swap data page into disk
func SwapLogPage(dp *dataPage) {

}

//fetch log page from disk
func FetchLogPage(pageId uint32) *dataPage {
	dp := &dataPage{}
	return dp
}

//swap log page into disk
func SwapDataPage(dp *dataPage) {

}
