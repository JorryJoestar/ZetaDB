package execution

import (
	"ZetaDB/container"
	"errors"
	"sort"
)

type ProjectionIterator struct {
	proIndexs []int

	newSchema     *container.Schema
	inputIterator Iterator

	nextTuple *container.Tuple
	hasNext   bool
}

//proIndexs are indexs of projected fields, it starts from 0
func NewProjectionIterator(proIndexs []int) *ProjectionIterator {

	projectIt := &ProjectionIterator{
		proIndexs: proIndexs,
		hasNext:   true}

	return projectIt
}

//throw error if iterator1 is nil or iterator2 is not nil
//throw error if proIndexs is invalid for schema of iterator1
func (pi *ProjectionIterator) Open(iterator1 Iterator, iterator2 Iterator) error {
	//throw error if iterator1 is nil or iterator2 is not nil
	if iterator1 == nil || iterator2 != nil {
		return errors.New("projectionIterator.go    Open() iterator invalid")
	}

	//set iterator1
	pi.inputIterator = iterator1

	//throw error if proIndexs is invalid for schema of iterator1
	oldSchemaDomainNum := pi.inputIterator.GetSchema().GetSchemaDomainNum()
	for _, v := range pi.proIndexs {
		if v >= oldSchemaDomainNum {
			return errors.New("projectionIterator.go    NewProjectionIterator() proIndexs invalid")
		}
	}

	//create newSchema from proIndexs and oldSchema
	var newDomainList []*container.Domain
	sort.Ints(pi.proIndexs)
	for _, v := range pi.proIndexs {
		currentDomain, _ := pi.inputIterator.GetSchema().GetSchemaDomain(v)
		newDomainList = append(newDomainList, currentDomain)
	}

	//TODO create newConstraintList
	var newConstraintList []*container.Constraint = nil

	newSchema, err := container.NewSchema(pi.inputIterator.GetSchema().GetSchemaTableName(), newDomainList, newConstraintList)
	if err != nil {
		return err
	}
	pi.newSchema = newSchema

	//set hasNext and nextTuple
	if !pi.inputIterator.HasNext() {
		pi.hasNext = false
	} else {
		pi.hasNext = true
		oldFormTuple, err := pi.inputIterator.GetNext()
		if err != nil {
			return err
		}

		var newFields []*container.Field
		for _, v := range pi.proIndexs {
			currentFieldBytes, _ := oldFormTuple.TupleGetFieldValue(v)
			currentField, err := container.NewFieldFromBytes(currentFieldBytes)
			if err != nil {
				return err
			}
			newFields = append(newFields, currentField)
		}
		pi.nextTuple, err = container.NewTuple(oldFormTuple.TupleGetTableId(), oldFormTuple.TupleGetTupleId(), pi.newSchema, newFields)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pi *ProjectionIterator) GetNext() (*container.Tuple, error) {
	//save returnTuple
	returnTuple := pi.nextTuple

	//set hasNext and nextTuple
	if !pi.inputIterator.HasNext() {
		pi.hasNext = false
	} else {
		pi.hasNext = true
		oldFormTuple, err := pi.inputIterator.GetNext()
		if err != nil {
			return nil, err
		}

		var newFields []*container.Field
		for _, v := range pi.proIndexs {
			currentFieldBytes, _ := oldFormTuple.TupleGetFieldValue(v)
			currentField, err := container.NewFieldFromBytes(currentFieldBytes)
			if err != nil {
				return nil, err
			}
			newFields = append(newFields, currentField)
		}
		pi.nextTuple, err = container.NewTuple(oldFormTuple.TupleGetTableId(), oldFormTuple.TupleGetTupleId(), pi.newSchema, newFields)
		if err != nil {
			return nil, err
		}
	}

	return returnTuple, nil
}

func (pi *ProjectionIterator) HasNext() bool {
	return pi.hasNext
}

func (pi *ProjectionIterator) Close() {
	pi.proIndexs = nil
	pi.newSchema = nil
	pi.inputIterator = nil
	pi.nextTuple = nil
	pi.hasNext = true
}

func (pi *ProjectionIterator) GetSchema() *container.Schema {
	return pi.newSchema
}
