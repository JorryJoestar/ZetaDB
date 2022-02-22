package storage

import "sync"

const (
	//bytes per page
	DEFAULT_PAGE_SIZE = 4096
)

type bufferPool struct {
	pageSize int
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

func (bf *bufferPool) GetPageSize() int {
	return bf.pageSize
}

func (bf *bufferPool) GetDataPage() *dataPage {
	dp := &dataPage{}
	return dp
}
