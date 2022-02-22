package storage

import "sync"

const (
	//bytes per page
	DEFAULT_PAGE_SIZE = 4096
)

type bufferPool struct {
	pageSize int
	buffer   []*page
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

//fetch the corresponding page from buffer, if the page is not buffered, fetch it from disk
func (bf *bufferPool) GetPage(pageId uint32) *page {
	p := &page{}
	return p
}

//swap a page to get empty slot in buffer
func (bf *bufferPool) EvictDataPage() {
}
