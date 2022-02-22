package storage

import "sync"

type storageEngine struct {
	fileLocation string
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

//fetch page from disk
func FetchPage(pageId uint32) *page {
	p := &page{}
	return p
}

//swap page into disk
func SwapPage(dp *page) {
	
}
