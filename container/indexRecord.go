package container

import (
	. "ZetaDB/utility"
	"errors"
)

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

//create a new IndexRecord from bytes
//throw error if bytes length invalid
func NewIndexRecordFromBytes(bytes []byte, elementLen int32) (*IndexRecord, error) {
	//throw error if bytes length invalid
	if int32(len(bytes)) != elementLen+5 {
		return nil, errors.New("bytes length invalid")
	}

	elementValue := bytes[:elementLen]
	indexDataPageId, _ := BytesToUint32(bytes[elementLen : elementLen+4])
	recordType := bytes[elementLen+4]

	newRecord := &IndexRecord{
		elementValue:    elementValue,
		indexDataPageId: indexDataPageId,
		recordType:      recordType}

	return newRecord, nil
}

//convert this indexRecord into bytes
func (ir *IndexRecord) IndexRecordToBytes() []byte {
	//elementValue
	bytes := ir.elementValue

	//indexDataPageId
	bytes = append(bytes, Uint32ToBytes(ir.indexDataPageId)...)

	//recordType
	bytes = append(bytes, ir.recordType)

	return bytes
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

//return the length of this record in bytes
func (ir *IndexRecord) IndexRecordLength() int32 {
	var size int32

	//elementValue
	size = int32(len(ir.elementValue))

	//indexDataPageId
	size += 4

	//recordType
	size += 1

	return size
}
