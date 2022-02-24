package storage

type dataBuffer struct {

	//page container
	buffer map[int]dataBuffer
}

//
func (db *dataBuffer) GetDataPageByPageId(pageId int) dataPage {
	dp := dataPage{}
	return dp
}

func (db *dataBuffer) InsertDataPage(dataPage) {

}
