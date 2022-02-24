package storage

type tuple struct {
	tableId uint32
	pageId  uint32
	tupleId uint32
	schema  *schema

	fields []field
}

//size of this tuple in bytes
func (t *tuple) TupleSizeInBytes() int {
	size := 0

	//for each field whose size is not fixed, length of 4 bytes is needed
	size += 4 * t.schema.UnfixedDomainNum()

	//add size of each field in this schema
	for _, v := range t.fields {
		size += v.FieldLen()
	}

	return size
}

//convert this tuple into a series of bytes, ready to push into disk
func (t *tuple) TupleToBytes() []byte {

	//slice to return
	var bytes []byte

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
		bytes = append(bytes, f.GetData()...)
	}
	return bytes
}

//generate a tuple from bytes, need to know the schema
func BytesToTuple(bytes []byte, s *schema) tuple {

	t := tuple{}
	t.schema = s

	for _, d := range t.schema.GetDomains() {

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

		//add current field into fields
		t.fields = append(t.fields, f)

		//delete data of current field from []byte
		bytes = bytes[l:]
	}

	return t
}
