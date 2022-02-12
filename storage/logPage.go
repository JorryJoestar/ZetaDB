package storage

type LogPage struct {
	pageId  uint32
	tableId uint32
}

func (page *LogPage) ToBitStram() {
}

func NewLogPage(data []uint8) *LogPage {
	return &LogPage{}
}
