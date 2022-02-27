package storage

import (
	. "ZetaDB/utility"
	"sync"
)

type bufferPool struct {

	//head pages of 16 key tables, should be loaded into bufferpool when this bufferpool is initialised //
	//k_userId_userName(userId INT PRIMARY KEY, userName VARCHAR(20))
	//head page number 0, tableId 0
	headPage_k_userId_userName dataPage

	//k_userId_password(userId INT PRIMARY KEY, password VARCHAR(20))
	//head page number 1, tableId 1
	headPage_k_userId_password dataPage

	//k_tableId_userId(tableId INT PRIMARY KEY, userId INT)
	//head page number 2, tableId 2
	headPage_k_tableId_userId dataPage

	//k_assertId_userId(assertId INT PRIMARY KEY, userId INT)
	//head page number 3, tableId 3
	headPage_k_assertId_userId dataPage

	//k_viewId_userId(viewId INT PRIMARY KEY, userId INT)
	//head page number 4, tableId 4
	headPage_k_viewId_userId dataPage

	//k_indexId_tableId(indexId INT PRIMARY KEY, tableId INT)
	//head page number 5, tableId 5
	headPage_k_indexId_tableId dataPage

	//k_triggerId_userId(triggerId INT PRIMARY KEY, userId INT)
	//head page number 6, tableId 6
	headPage_k_triggerId_userId dataPage

	//k_psmId_userId(psmId INT PRIMARY KEY, userId INT)
	//head page number 7, tableId 7
	headPage_k_psmId_userId dataPage

	//k_tableId_schema(tableId INT PRIMARY KEY, schema VARCHAR(255))
	//head page number 8, tableId 8
	headPage_k_tableId_schema dataPage

	//k_table(tableId INT PRIMARY KEY, headPageId INT, tupleIndexId INT, tupleNum INT)
	//head page number 9, tableId 9
	headPage_k_table dataPage

	//k_assert(assertId INT PRIMARY KEY, assertStmt VARCHAR(255))
	//head page number 10, tableId 10
	headPage_k_assert dataPage

	//k_view(viewId INT PRIMARY KEY, viewStmt VARCHAR(255))
	//head page number 11, tableId 11
	headPage_k_view dataPage

	//k_index(indexId INT PRIMARY KEY, logHeadPageId INT)
	//head page number 12, tableId 12
	headPage_k_index dataPage

	//k_trigger(triggerId INT PRIMARY KEY, triggerStmt VARCHAR(255))
	//head page number 13, tableId 13
	headPage_k_trigger dataPage

	//k_psm(psmId INT PRIMARY KEY, psmStmt VARCHAR(255))
	//head page number 14, tableId 14
	headPage_k_psm dataPage

	//k_emptyPageSlot(pageId INT)
	//head page number 15, tableId 15
	headPage_k_emptyPageSlot dataPage

	//k_log page(log INT PRIMARY KEY, logStmt VARCHAR(20))
	//head page number 16, tableId 16
	headPage_k_log dataPage

	//bytes per page in this bufferpool
	dataPageSize int

	buffer []dataPage
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

	//assign data page size
	bp.dataPageSize = DEFAULT_DATAPAGE_SIZE
}

//pageSize getter
func (bf *bufferPool) GetDataPageSize() int {
	return bf.dataPageSize
}

//fetch the corresponding page from buffer, if the page is not buffered, fetch it from disk
func (bf *bufferPool) GetDataPage(pageId uint32) *dataPage {
	p := &dataPage{}
	return p
}

//swap a page to get empty slot in buffer
func (bf *bufferPool) EvictDataPage() {
}
