package storage

type IndexRecord struct {
	elementValue    []byte
	indexDataPageId uint32
	recordType      uint8
}

//create a new IndexRecord
func NewIndexRecord(elementValue []byte, indexDataPageId uint32, recordType uint8) *IndexRecord {
	newRecord := &IndexRecord{
		elementValue:    elementValue,
		indexDataPageId: indexDataPageId,
		recordType:      recordType}

	return newRecord
}

//elementValue getter
func (ir *IndexRecord) IndexRecordGetElementValue() []byte {
	return ir.elementValue
}

//indexDataPageId getter
func (ir *IndexRecord) IndexRecordGetIndexDataPageId() uint32 {
	return ir.indexDataPageId
}

//recordType getter
func (ir *IndexRecord) IndexRecordGetRecordType() uint8 {
	return ir.recordType
}
