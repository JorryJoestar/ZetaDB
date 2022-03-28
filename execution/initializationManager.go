package execution

import (
	"ZetaDB/container"
	"ZetaDB/storage"
)

/*
	- k_userId_userName(userId INT PRIMARY KEY, userName VARCHAR(20))
		head page number 0, tableId 0
	- k_userId_password(userId INT PRIMARY KEY, password VARCHAR(20))
		head page number 1, tableId 1
	- k_tableId_userId(tableId INT PRIMARY KEY, userId INT)
		head page number 2, tableId 2
	- k_assertId_userId(assertId INT PRIMARY KEY, userId INT)
		head page number 3, tableId 3
	- k_viewId_userId(viewId INT PRIMARY KEY, userId INT)
		head page number 4, tableId 4
	- k_indexId_tableId(indexId INT PRIMARY KEY, tableId INT)
		head page number 5, tableId 5
	- k_triggerId_userId(triggerId INT PRIMARY KEY, userId INT)
		head page number 6, tableId 6
	- k_psmId_userId(psmId INT PRIMARY KEY, userId INT)
		head page number 7, tableId 7
	- k_tableId_schema(tableId INT PRIMARY KEY, schema VARCHAR(255))
		head page number 8, tableId 8
	- k_table(tableId INT PRIMARY KEY, headPageId INT, lastTupleId INT, tupleNum INT)
		head page number 9, tableId 9
	- k_assert(assertId INT PRIMARY KEY, assertStmt VARCHAR(255))
		head page number 10, tableId 10
	- k_view(viewId INT PRIMARY KEY, viewStmt VARCHAR(255))
		head page number 11, tableId 11
	- k_index(indexId INT PRIMARY KEY, indexStmt VARCHAR(255), indexHeadPageId INT)
		head page number 12, tableId 12
	- k_trigger(triggerId INT PRIMARY KEY, triggerStmt VARCHAR(255))
		head page number 13, tableId 13
	- k_psm(psmId INT PRIMARY KEY, psmStmt VARCHAR(255))
		head page number 14, tableId 14
	- k_emptyDataPageSlot(pageId INT)
		head page number 15, tableId 15
	- k_emptyIndexPageSlot(pageId INT)
		head page number 16, tableId 16
*/

type InitializationManager struct {
	se              *storage.StorageEngine
	keyTableSchemas [17]*container.Schema
}

func NewInitializationManager(se *storage.StorageEngine) *InitializationManager {
	im := &InitializationManager{
		se: se}

	return im
}

func (im *InitializationManager) InitializeSystem() {
	//create p0, p1, p2, p8, p9, p15, p16
	p0 := storage.NewDataPageMode0(0, 0, 0, 0)
	p1 := storage.NewDataPageMode0(1, 1, 1, 1)
	p2 := storage.NewDataPageMode0(2, 2, 2, 2)
	p8 := storage.NewDataPageMode0(8, 8, 8, 8)
	p9 := storage.NewDataPageMode0(9, 9, 9, 9)
	p15 := storage.NewDataPageMode0(15, 15, 15, 15)
	p16 := storage.NewDataPageMode0(16, 16, 16, 16)

	//insert these pages into storageEngine
	im.se.InsertDataPage(p0)
	im.se.InsertDataPage(p1)
	im.se.InsertDataPage(p2)
	im.se.InsertDataPage(p8)
	im.se.InsertDataPage(p9)
	im.se.InsertDataPage(p15)
	im.se.InsertDataPage(p16)

	//insert tuple into page0

	//insert tuple into page1

	//insert tuple into page2

	//insert tuple into page8

	//insert tuple into page9

	//insert tuple into page15

	//insert tuple into page16

	//swap these pages into disk
	im.se.SwapDataPage(0)
	im.se.SwapDataPage(1)
	im.se.SwapDataPage(2)
	im.se.SwapDataPage(8)
	im.se.SwapDataPage(9)
	im.se.SwapDataPage(15)
	im.se.SwapDataPage(16)
}
