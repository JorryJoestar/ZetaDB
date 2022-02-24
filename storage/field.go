package storage

import "errors"

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
	return len(f.data)
}

func (f *field) GetFieldData() []byte {
	return f.data
}

//convert field data into CHAR (string)
func (f *field) FieldToChar() (string, error) {
	if f.isNull {
		return "", errors.New("storage.field.go: field data is NULL")
	}

	if f.FieldLen() != 1 {
		return "", errors.New("storage.field.go: field data length mismatch")
	}

	return string(f.data), nil
}

func (f *field) IsNull() bool {
	return f.isNull
}
