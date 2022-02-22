package storage

type page struct {
	pageId      uint32
	tableId     uint32
	priorPageId uint32
	nextPageId  uint32
}

func (page *page) ToBitStram() {
}
