package storage

type dataBuffer struct {
	//max amount of pages this buffer can hold
	pageMaxNumber int
	pages         []*dataPage
	bufferMapper  []int
}

func (db *dataBuffer) GetpageMaxNumber() int {
	return db.pageMaxNumber
}

func (db *dataBuffer) GetPage(pageId uint32) *dataPage {
	dp := &dataPage{}
	return dp
}
