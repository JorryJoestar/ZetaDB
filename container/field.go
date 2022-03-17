package container

import "errors"

//store data of a particular field within a field
//fields compose a tuple
type Field struct {
	isNull bool
	data   []byte
}

//generate a field from a series of bytes
//throw error if bytes length is 0
func NewFieldFromBytes(bytes []byte) (*Field, error) {

	if len(bytes) == 0 {
		return nil, errors.New("bytes length invalid")
	}

	f := &Field{
		isNull: false,
		data:   bytes}

	return f, nil
}

//generate a null field
func NewNullField() *Field {

	f := &Field{
		isNull: true}

	return f

}

//length of field in bytes
func (f *Field) FieldLen() int {

	if f.isNull {
		return 0
	} else {
		return len(f.data)
	}
}

//return data bytes that this field holds
//throw error if this field is null
func (f *Field) FieldToBytes() ([]byte, error) {

	if f.isNull {
		return nil, errors.New("null field")
	} else {
		return f.data, nil
	}
}

//return true if this field is null
func (f *Field) FieldIsNull() bool {

	return f.isNull
}

//set this field to null, delete its data
func (f *Field) SetFieldNull() {

	f.isNull = true
	f.data = nil
}
