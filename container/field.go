package storage

import "errors"

//store data of a particular field within a field
//fields compose a tuple
type field struct {
	isNull bool
	data   []byte
}

//generate a field from a series of bytes
func NewFieldFromBytes(bytes []byte) (*field, error) {

	if len(bytes) == 0 {
		return nil, errors.New("bytes length invalid")
	}

	f := &field{
		isNull: false,
		data:   bytes}

	return f, nil
}

//generate a null field
func NewNullField() *field {

	f := &field{
		isNull: true}

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

//return data bytes that this field holds
//if this field is null, throw error
func (f *field) FieldToBytes() ([]byte, error) {

	if f.isNull {
		return nil, errors.New("null field")
	} else {
		return f.data, nil
	}
}

//return true if this field is null
func (f *field) FieldIsNull() bool {

	return f.isNull
}

//set this field to null, delete its data
func (f *field) SetFieldNull() {

	f.isNull = true
	f.data = nil
}
