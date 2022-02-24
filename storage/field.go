package storage

//store data of a particular field within a field
//fields compose a tuple
type field struct {
	data []byte
}

//generate a field from a series of bytes
func BytesToField(bytes []byte) field {
	f := field{
		data: bytes}

	return f
}

//length of field in bytes
func (f *field) FieldLen() int {
	return len(f.data)
}

func (f *field) GetData() []byte {
	return f.data
}
