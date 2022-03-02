package storage

type Tuple struct {
	tableId uint32
	pageId  uint32
	tupleId uint32
	schema  *schema
	fields  []field
}

//size of this tuple in bytes
func (t *Tuple) TupleSizeInBytes() int {
	size := 4 //tupleId occupies 4 bytes

	//byte number for isNull info
	if t.schema.GetDomainNum()%8 == 0 {
		size += t.schema.GetDomainNum() / 8
	} else {
		size += t.schema.GetDomainNum()/8 + 1
	}

	//for each field whose size is not fixed, length of 4 bytes is needed
	size += 4 * t.schema.UnfixedDomainNum()

	//add size of each field in this schema
	for _, v := range t.fields {
		size += v.FieldLen()
	}

	return size
}

//convert this tuple into a series of bytes, ready to push into disk
func (t *Tuple) TupleToBytes() []byte {

	//slice to return
	var bytes []byte

	//convert tupleId into []byte for store in litte-endian
	var tupleIdBytes []byte
	tupleIdBytes[0] = uint8(t.tupleId)
	tupleIdBytes[1] = uint8(t.tupleId >> 8)
	tupleIdBytes[2] = uint8(t.tupleId >> 16)
	tupleIdBytes[3] = uint8(t.tupleId >> 24)

	//append tupleId into result slice
	bytes = append(bytes, tupleIdBytes...)

	//generate slice to keep null information
	var nullBytes []byte
	var nullByte byte = 0
	for i, f := range t.fields {
		if f.IsNull() { // push 1 into this byte
			nullByte += 1 << (i % 8)
		} else { //push 0 into this byte
			//do nothing
		}
		if (i+1)%8 == 0 { //nullByte full, should be push into nullBytes
			nullBytes = append(nullBytes, nullByte)
			nullByte = 0
		}
	}
	if len(t.fields)%8 != 0 {
		nullBytes = append(nullBytes, nullByte)
	}

	//append nullBytes into result slice
	bytes = append(bytes, nullBytes...)

	for i, d := range t.schema.GetDomains() {

		//get corresponding field
		f := t.fields[i]

		if d.DomainSizeUnfixed() { //current domain size unfixed, length should be stored

			//length of field in byte, store it before field data
			l := f.FieldLen()

			//convert int l into []byte length for store in litte-endian
			var length []byte
			length[0] = uint8(l)
			length[1] = uint8(l >> 8)
			length[2] = uint8(l >> 16)
			length[3] = uint8(l >> 24)

			//append length into result slice
			bytes = append(bytes, length...)

		}

		//append field data into result slice
		bytes = append(bytes, f.GetFieldData()...)
	}
	return bytes
}

//generate a tuple from bytes, need to know the schema
func BytesToTuple(bytes []byte, s *schema) Tuple {

	t := Tuple{}
	t.schema = s

	//fetch tupleId
	tupleId := uint32(bytes[0]) + uint32(bytes[1])<<8 + uint32(bytes[2])<<16 + uint32(bytes[3])<<24
	t.tupleId = tupleId
	bytes = bytes[4:]

	//fetch isNull info
	nullBytesNum := 0
	if t.schema.GetDomainNum()%8 == 0 {
		nullBytesNum = t.schema.GetDomainNum() / 8
	} else {
		nullBytesNum = t.schema.GetDomainNum()/8 + 1
	}
	nullBytes := bytes[:nullBytesNum]
	bytes = bytes[nullBytesNum:]

	for i, d := range t.schema.GetDomains() {

		//current field length
		var l int
		if d.DomainSizeUnfixed() { //current field size unfixed, length is stored

			//convert [4]bytes into int, get length
			l = int(bytes[0]) + (int(bytes[1]) << 8) + (int(bytes[2]) << 16) + (int(bytes[3]) << 24)

			//delete length in bytes
			bytes = bytes[4:]

		} else { //current field size fixed, length can be fetched from domain

			//get length from current domain
			l = d.DomainSizeInBytes()
		}

		//generate field from the first l bytes
		f := BytesToField(bytes[:l])

		//update f isNull
		if nullBytes[i/8]>>(i%8) == 1 {
			f.isNull = true
		} else {
			f.isNull = false
		}

		//add current field into fields
		t.fields = append(t.fields, f)

		//delete data of current field from []byte
		bytes = bytes[l:]
	}

	return t
}

func (t *Tuple) GetTupleId() uint32 {
	return t.tupleId
}

func (t *Tuple) SetPageId(id uint32) {
	t.pageId = id
}
