package storage

//TODO
type dataBuffer struct {

	//page container
	buffer map[uint32]*dataBuffer

	//mapping from file pageId to buffer index
	mapper map[uint32]uint32

	//record all empty buffer slot index
	bufferSlots []uint32
}

//fetch a data page from data buffer by its pageId
func (db *dataBuffer) GetDataPageByPageId(pageId uint32) (*dataPage, error) {
	dp := &dataPage{}
	return dp, nil
}

//fetch a data page from data buffer by its buffer index
func (db *dataBuffer) GetDataPageByBufferIndex(pos uint32) (*dataPage, error) {
	dp := &dataPage{}
	return dp, nil
}

//insert a data page into the buffer, then recording the mapping relation between pageId and buffer index
func (db *dataBuffer) InsertDataPage(*dataPage) error {
	return nil
}

//convert pageId to dataBufferId
//if no such pageId in mapper, throw error
func (db *dataBuffer) PageIdToDataBufferId(pageId uint32) (uint32, error) {
	return 0, nil
}

//convert dataBufferId to pageId
//if no such dataBufferId in mapper, throw error
func (db *dataBuffer) DataBufferIdToPageId(dbId uint32) (uint32, error) {
	return 0, nil
}
