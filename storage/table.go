package storage

type table struct {
	headPageId  uint32
	tableId     uint32
	tableSchema *schema
}