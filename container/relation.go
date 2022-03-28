package container

type Relation struct {
	headPageId  uint32
	tableId     uint32
	tableSchema Schema
}
