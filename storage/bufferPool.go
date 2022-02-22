package storage

import "sync"

const (
	//bytes per page
	DEFAULT_PAGE_SIZE = 4096
)

type bufferPool struct {
	pageSize    int
	dataBuffer  *dataBuffer
	logBuffer   *logBuffer
	indexBuffer *indexBuffer
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

//initialise bufferpool before use after system boot
func InitialiseBufferPool() {
	bp := GetBufferPool()

	//assign page size
	bp.pageSize = DEFAULT_PAGE_SIZE
}

//pageSize getter
func (bf *bufferPool) GetPageSize() int {
	return bf.pageSize
}

//fetch the corresponding page from dataBuffer, if the page is not buffered, fetch it from disk
func (bf *bufferPool) GetDataPage(pageId uint32) *dataPage {
	dp := &dataPage{}
	return dp
}

//swap a page to get empty slot in dataBuffer
func (bf *bufferPool) EvictDataPage() {
}

//fetch the corresponding page from logBuffer, if the page is not buffered, fetch it from disk
func (bf *bufferPool) GetLogPage(pageId uint32) *logPage {
	lp := &logPage{}
	return lp
}

//swap a page to get empty slot in logBuffer
func (bf *bufferPool) EvictLogPage() {
}

//fetch the corresponding page from indexBuffer, if the page is not buffered, fetch it from disk
func (bf *bufferPool) GetIndexPage(pageId uint32) *indexPage {
	ip := &indexPage{}
	return ip
}

//swap a page to get empty slot in indexBuffer
func (bf *bufferPool) EvictIndexPage() {
}
