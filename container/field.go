package storage

//store data of a particular field within a field
//fields compose a tuple
type field struct {
	isNull bool
	data   []byte
}

//generate a field from a series of bytes
func BytesToField(bytes []byte) field {
	f := field{
		data: bytes}

	return f
}

//length of field in bytes
func (f *field) FieldLen() int {
	if f.isNull {
		return 0
	} else {
		return len(f.data)
	}
}

func (f *field) GetFieldData() []byte {
	if f.isNull {
		return nil
	} else {
		return f.data
	}
}

func (f *field) IsNull() bool {
	return f.isNull
}

func (f *field) SetNull(b bool) {
	f.isNull = b
}
