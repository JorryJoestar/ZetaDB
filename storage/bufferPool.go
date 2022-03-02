package storage

import (
	"sync"
)

type bufferPool struct {
	//container for holding key pages, they are head pages of key tables
	keyPages *keyPageManager

	//buffer for data page
	dBuffer *dataBuffer

	//buffer for index page
	iBuffer *indexBuffer

	//buffer for log page
	lBuffer *logBuffer
}

//use GetBufferPool to get the unique bufferPool
var bpInstance *bufferPool
var bpOnce sync.Once

func GetBufferPool() *bufferPool {
	bpOnce.Do(func() {
		bpInstance = &bufferPool{}
	})
	return bpInstance
}

//TODO
//get a data page according to its pageId
//if this page is modified, remember to swap it
func (bp *bufferPool) GetDataPage(pageId uint32) (*dataPage, error) {
	return nil, nil
}

//TODO
//insert a newly created data page into buffer, but not swapped into disk
//remember to swap it
func (bp *bufferPool) InsertDataPage(*dataPage) error {
	return nil
}

//TODO
//delete a data page according to its pageId, related page not swapped into disk
//remember to swap related pages
func (bp *bufferPool) DeleteDataPage(pageId uint32) error {
	return nil
}

//TODO
//swap a data page into disk according to its pageId
func (bp *bufferPool) SwapDataPage(pageId uint32) error {
	return nil
}

//TODO
//get an index page according to its pageId
//if this page is modified, remember to swap it
func (bp *bufferPool) GetIndexPage(pageId uint32) (*indexPage, error) {
	return nil, nil
}

//TODO
//insert a newly created index page into buffer, but not swapped into disk
//remember to swap it
func (bp *bufferPool) InsertIndexPage(*indexPage) error {
	return nil
}

//TODO
//delete an index page according to its pageId, related page not swapped into disk
//remember to swap related pages
func (bp *bufferPool) DeleteIndexPage(pageId uint32) error {
	return nil
}

//TODO
//swap an index page into disk according to its pageId
func (bp *bufferPool) SwapIndexPage(pageId uint32) error {
	return nil
}

//TODO
//get a log page according to its pageId
//if this page is modified, remember to swap it
func (bp *bufferPool) GetLogPage(pageId uint32) (*logPage, error) {
	return nil, nil
}

//TODO
//insert a newly created log page into buffer, but not swapped into disk
//remember to swap it
func (bp *bufferPool) InsertLogPage(*logPage) error {
	return nil
}

//TODO
//delete a log page according to its pageId, related page not swapped into disk
//remember to swap related pages
func (bp *bufferPool) DeleteLogPage(pageId uint32) error {
	return nil
}

//TODO
//swap a log page into disk according to its pageId
func (bp *bufferPool) SwapLogPage(pageId uint32) error {
	return nil
}
