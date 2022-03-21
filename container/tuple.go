package container

import (
	. "ZetaDB/utility"
	"errors"
)

/*
                                tuple structure in disk
	---------------------------------------------------------------------------------
	|| tupleId || isNull bytes || field1 | field2 | length3 | field3 | ... |fieldN ||
	---------------------------------------------------------------------------------

	~tupleId:
		-unique within a table
		-type is uint32
		-scope: 0 - 4,294,967,295
		-space: 4 bytes

	~isNull bytes:
		-structure
			 high     byte 0     low	   high     byte 1     low
			-------------------------     -------------------------
			|0 |1 |0 |0 |1 |1 |0 |1 |     |1 |1 |0 |1 |1 |0 |0 |0 |   ... ...
			-------------------------     -------------------------
			 7  6  5  4  3  2  1  0        15 14 13 12 11 10 9  8
		-field x is at byte x/8, index x%8 from low to high
		-1 means null, 0 means not null
		-space: see method TupleSizeOfIsNullBytes()

	~fields and lengths

		~length
			-it is valid if this field is unfixed
			-it is appended in front of the field data bytes
			-int32, 4 bytes

		~null fields
			-length fixed: padding bytes are needed
			-length unfixed: padding bytes are not needed, but set length to 0
*/

type Tuple struct {
	tableId     uint32
	tupleId     uint32
	schema      *Schema
	isNullBytes []byte //should not be used except to/from bytes
	fields      []*Field
}

//generate a tuple from bytes, need to know the schema
func NewTupleFromBytes(bytes []byte, s *Schema, tableId uint32) (*Tuple, error) {

	t := &Tuple{}
	t.schema = s
	t.tableId = tableId

	//fetch tupleId
	tupleId, tupleIdErr := BytesToUint32(bytes[:4])
	if tupleIdErr != nil {
		return nil, tupleIdErr
	}
	t.tupleId = tupleId
	bytes = bytes[4:]

	//fetch isNull info
	nullBytesNum := 0
	if t.schema.GetSchemaDomainNum()%8 == 0 {
		nullBytesNum = t.schema.GetSchemaDomainNum() / 8
	} else {
		nullBytesNum = t.schema.GetSchemaDomainNum()/8 + 1
	}
	t.isNullBytes = bytes[:nullBytesNum]
	bytes = bytes[nullBytesNum:]

	//fetch fields
	for i, d := range t.schema.GetSchemaDomains() {

		//current field length
		var l int32
		if d.DomainSizeUnfixed() { //current field size unfixed, length is stored

			//convert [4]bytes into int, get length
			size, dErr := BytesToINT(bytes[:4])
			if dErr != nil {
				return nil, dErr
			}
			l = size

			//delete length in bytes
			bytes = bytes[4:]

		} else { //current field size fixed, length can be fetched from domain

			//get length from current domain
			size, dErr := d.DomainSizeInBytes()
			if dErr != nil {
				return nil, dErr
			}
			l = int32(size)
		}

		//check if this tuple is null value
		isNull, nullErr := t.TupleFieldIsNull(i)
		if nullErr != nil {
			return nil, nullErr
		}

		//create a field
		var f *Field
		if l == 0 { //VARCHAR or BITVARYING
			f = NewNullField()
		} else if isNull {
			f = NewNullField()

			//still need to drop field bytes
			bytes = bytes[l:]
		} else {
			field, fErr := NewFieldFromBytes(bytes[:l])
			if fErr != nil {
				return nil, fErr
			}
			f = field

			//delete data of current field from []byte
			bytes = bytes[l:]
		}

		//add current field into fields
		t.fields = append(t.fields, f)
	}

	return t, nil
}

//create a new tuple
//throw error if fields slice length is different from schema domain numbers
//throw error if field length in byte different from the corresponding domain capacity
func NewTuple(tableId uint32, tupleId uint32, schema *Schema, fields []*Field) (*Tuple, error) {

	//throw error for fields slice length and schema domain number unmatching
	if schema.GetSchemaDomainNum() != len(fields) {
		return nil, errors.New("fields slice length and schema domain number unmatching")
	}

	//throw error for field length unmatching with the corresponding domain capacity
	domains := schema.GetSchemaDomains()
	for i, domain := range domains {
		length, err := domain.DomainSizeInBytes()
		if err == nil && !fields[i].FieldIsNull() { //if err != nil, domain size is unfixed
			if length != fields[i].FieldLen() {
				return nil, errors.New("field length unmatching with the corresponding domain capacity")
			}
		}
	}

	newTuple := &Tuple{
		tableId: tableId,
		tupleId: tupleId,
		schema:  schema,
		fields:  fields}

	//update the isNullBytes
	nullBytesNum := newTuple.TupleSizeOfIsNullBytes()

	newTuple.isNullBytes = make([]byte, nullBytesNum)
	for i := 0; i < nullBytesNum; i++ {
		newTuple.isNullBytes[i] = 0
	}

	for i, field := range newTuple.TupleGetFields() {
		if field.FieldIsNull() {
			newTuple.TupleSetFieldNull(i)
		}
	}

	return newTuple, nil
}

//convert this tuple into a series of bytes, ready to push into disk
func (t *Tuple) TupleToBytes() ([]byte, error) {

	//slice to return
	var bytes []byte

	//convert tupleId into []byte
	tupleIdBytes := Uint32ToBytes(t.TupleGetTupleId())

	//append tupleId into result slice
	bytes = append(bytes, tupleIdBytes...)

	//append isNullBytes into result slice
	bytes = append(bytes, t.isNullBytes...)

	for i, d := range t.schema.GetSchemaDomains() {

		//get corresponding field
		f := t.fields[i]

		if d.DomainSizeUnfixed() { //current domain size unfixed, length should be stored

			//length of field in byte, store it before field data
			l := f.FieldLen()

			//convert int l into []byte length for store in litte-endian
			length := INTToBytes(int32(l))

			//append length into result slice
			bytes = append(bytes, length...)

		}

		//append field data into result slice
		fieldBytes, fErr := f.FieldToBytes()
		if fErr != nil { //this field is a null field
			if !d.DomainSizeUnfixed() { //padding empty bytes for fixed length null field
				l, _ := d.DomainSizeInBytes()
				fieldBytes = make([]byte, l)
				for i := 0; i < l; i++ {
					fieldBytes[i] = 0
				}
				bytes = append(bytes, fieldBytes...)
			}
			//if this field is not fixed in length, do nothing
		} else {
			bytes = append(bytes, fieldBytes...)
		}

	}
	return bytes, nil
}

//tupleId getter
func (t *Tuple) TupleGetTupleId() uint32 {
	return t.tupleId
}

//tupleId setter
func (t *Tuple) TupleSetTupleId(tupId uint32) {
	t.tupleId = tupId
}

//tableId getter
func (t *Tuple) TupleGetTableId() uint32 {
	return t.tableId
}

//get a field according to its index
//throw error if index invalid
func (t *Tuple) TupleGetFieldValue(index int) ([]byte, error) {

	//throw error if index invalid
	if index >= t.schema.GetSchemaDomainNum() {
		return nil, errors.New("index invalid")
	}

	return t.fields[index].data, nil
}

//set value of a field according to its index
//throw error if index invalid
//throw error if input value bytes length is 0
//throw error if input value bytes length unmatch corresponding domain
func (t *Tuple) TupleSetFieldValue(data []byte, index int) error {

	//throw error if index invalid
	if index >= t.schema.GetSchemaDomainNum() {
		return errors.New("index invalid")
	}

	//throw error if input value bytes length is 0
	if len(data) == 0 {
		return errors.New("data length invalid")
	}

	//throw error if input value bytes length unmatch corresponding domain
	domain, _ := t.schema.GetSchemaDomain(index)
	if !domain.DomainSizeUnfixed() { //domain size is fixed
		domainSize, _ := domain.DomainSizeInBytes()
		if domainSize != len(data) {
			return errors.New("data length invalid")
		}
	}

	t.fields[index].data = data

	isNull, _ := t.TupleFieldIsNull(index)
	if isNull {
		t.fields[index].isNull = false
		byteIndex := index / 8
		inByteIndex := index % 8
		var mask byte = 1 << inByteIndex
		t.isNullBytes[byteIndex] = t.isNullBytes[byteIndex] - mask
	}
	return nil
}

//get fields
func (t *Tuple) TupleGetFields() []*Field {
	return t.fields
}

//return true if this field is null
func (t *Tuple) TupleFieldIsNull(index int) (bool, error) {

	byteIndex := index / 8
	inByteIndex := index % 8
	var mask byte = 1 << inByteIndex
	isNull := t.isNullBytes[byteIndex] & mask
	if isNull == 0 { //not null
		return false, nil
	} else { //is null
		return true, nil
	}
}

//set a field to null
//throw error if field is already null
func (t *Tuple) TupleSetFieldNull(index int) error {

	//throw error if field is already null
	isNull, _ := t.TupleFieldIsNull(index)
	if isNull {
		return errors.New("field is already null")
	}

	byteIndex := index / 8
	inByteIndex := index % 8
	var mask byte = 1 << inByteIndex
	t.isNullBytes[byteIndex] += mask

	return nil
}

//return size of isNullBytes
func (t *Tuple) TupleSizeOfIsNullBytes() int {
	size := 0

	if t.schema.GetSchemaDomainNum()%8 == 0 {
		size += t.schema.GetSchemaDomainNum() / 8
	} else {
		size += t.schema.GetSchemaDomainNum()/8 + 1
	}

	return size
}

//size of this tuple in bytes
func (t *Tuple) TupleSizeInBytes() int {

	//tupleId occupies 4 bytes
	size := 4

	//byte number for isNull info
	size += t.TupleSizeOfIsNullBytes()

	//for each field whose size is not fixed, length of 4 bytes is needed
	size += 4 * t.schema.UnfixedDomainNum()

	//add size of each field in this schema
	for i, v := range t.fields {
		if v.FieldIsNull() {
			d, _ := t.schema.GetSchemaDomain(i)
			if !d.DomainSizeUnfixed() { //this field size is fixed but is null
				l, _ := d.DomainSizeInBytes()
				size += l
			}
		} else {
			size += v.FieldLen()
		}
	}

	return size
}

//return true if size of this tuple is fixed
func (t *Tuple) TupleSizeFixed() bool {
	return t.schema.UnfixedDomainNum() == 0
}

//generate a key for map, it is determined by all fields content
//if fields of two tuples are identical, they must have the same map key
//throw error if any single field length in bytes is over DEFAULT_TUPLE_SINGAL_FIELD_OVER_LONG_LENGTH
//throw error if total length of this tuple is over DEFAULT_TUPLE_TOTAL_OVER_LONG_LENGTH
/*

	map key structure:
		field0Length + field0InHex + field1Length + field1InHex + ... + field(n-1)Length + field(n-1)InHex
	-if a field is null, the corresponding length is 0 and no fieldInHex included
	-field0Length is uint16 in hex form

*/
func (t *Tuple) TupleGetMapKey() (string, error) {

	var currentFieldLen uint16 = 0
	var totalLen uint16 = 0
	keyString := ""

	for _, field := range t.TupleGetFields() {

		fieldData, nullErr := field.FieldToBytes()

		if nullErr != nil { //this is a null field

			currentFieldLen = 0
			keyString += BytesToHexString(Uint16ToBytes(currentFieldLen))

		} else { //this is  not a null field

			if len(fieldData) > int(UINT16_MAX) { //field length is greater than max range of uint16
				return "", errors.New("tuple.go    TupleGetMapKey() field over long")

			}

			if uint16(len(fieldData)) > DEFAULT_TUPLE_SINGAL_FIELD_OVER_LONG_LENGTH { //this field is too large, throw error
				return "", errors.New("tuple.go    TupleGetMapKey() field over long")
			}

			currentFieldLen = uint16(len(fieldData))
			totalLen += currentFieldLen

			if totalLen > DEFAULT_TUPLE_TOTAL_OVER_LONG_LENGTH { //this tuple is too large, thorw error
				return "", errors.New("tuple.go    TupleGetMapKey() tuple over long")
			}

			keyString += BytesToHexString(Uint16ToBytes(currentFieldLen))
			keyString += BytesToHexString(fieldData)

		}
	}

	return keyString, nil
}
