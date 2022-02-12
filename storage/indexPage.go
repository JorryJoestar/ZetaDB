package storage

type IndexPage struct {
	pageId  uint32
	tableId uint32
}

func (page *IndexPage) ToBitStram() {
}

func NewIndexPage(data []uint8) *IndexPage {
	return &IndexPage{}
}
