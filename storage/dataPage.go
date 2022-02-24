package storage

type dataPage struct {
	//slot number in data.zdb
	pageId uint32

	//id of table which this page belongs to
	tableId uint32

	//id of prior page
	//if priorPageId == pageId, this page is a head page
	priorPageId uint32

	//id of next page
	//if nextPageId == pageId, this page is a tail page
	nextPageId uint32

	//the id of first tuple in this page
	//tuple id is unique within a table
	beginTupleId uint32

	//number of tuples in this page
	tupleNum uint32

	//tuples in this page
	tuples []tuple
}

func (dataPage *dataPage) ToBytes() {}

func (dataPage *dataPage) SizeInByte() int { return 0 }

func (dataPage *dataPage) VacantByteNum() int { return 0 }

func (dataPage *dataPage) InsertTuple() {}

func (dataPage *dataPage) DeleteTuple() {}

func (dataPage *dataPage) IsHeadPage() {}

func (dataPage *dataPage) IsTailPage() {}

func (dataPage *dataPage) GetPageId() {}

func (dataPage *dataPage) GetTableId() {}

func (dataPage *dataPage) GetPriorPageId() {}

func (dataPage *dataPage) GetNextPageId() {}
