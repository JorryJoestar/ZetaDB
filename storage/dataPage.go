package storage

type DataPage struct {
	pageId  uint32
	tableId uint32
}

func (page *DataPage) ToBitStram() {
}

func NewDataPage(data []uint8) *DataPage {
	return &DataPage{}
}
