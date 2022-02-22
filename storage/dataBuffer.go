package storage

import "sync"

type dataBuffer struct {
	//max amount of pages this buffer can hold
	pageMaxNumber int
	pages         []*dataPage
}

//use GetBufferPool to get the unique bufferPool
var dbInstance *dataBuffer
var dbOnce sync.Once

func GetDataBuffer() *dataBuffer {
	dbOnce.Do(func() {
		dbInstance = &dataBuffer{}
	})
	return dbInstance
}

func (db *dataBuffer) GetpageMaxNumber() int {
	return db.pageMaxNumber
}

func (db *dataBuffer) GetPage(pageId uint32) *dataPage {
	dp := &dataPage{}
	return dp
}
