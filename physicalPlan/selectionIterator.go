package execution

import (
	"ZetaDB/assister"
	"ZetaDB/container"
	"errors"
)

type SelectionIterator struct {
	currentTuple *container.Tuple
	hasNext      bool
	condition    *container.Condition
	iterator1    Iterator
}

//SelectionIterator constructor
func NewSelectionIterator(condition *container.Condition) *SelectionIterator {
	si := &SelectionIterator{
		hasNext:   true,
		condition: condition,
	}
	return si
}

//throw error if iterator1 is nil or iterator2 is not nil
func (si *SelectionIterator) Open(iterator1 Iterator, iterator2 Iterator) error {

	//throw error if iterator1 is nil or iterator2 is not nil
	if iterator1 == nil || iterator2 != nil {
		return errors.New("selectionIterator.go    Open() parameter invalid")
	}

	//set iterator1
	si.iterator1 = iterator1

	predicator := assister.GetPredicator()
	for si.iterator1.HasNext() {
		tuple, _ := si.iterator1.GetNext()
		if predicator.CheckCondition(si.condition, tuple) {
			si.currentTuple = tuple
			return nil
		}
	}

	si.hasNext = false
	return nil
}

func (si *SelectionIterator) GetNext() (*container.Tuple, error) {
	//throw error if HasNext() returns false
	if !si.HasNext() {
		return nil, errors.New("selectionIterator.go    GetNext() hasNext false")
	}

	//save currentTuple for return
	tupleToReturn := si.currentTuple

	predicator := assister.GetPredicator()
	for si.iterator1.HasNext() {
		tuple, _ := si.iterator1.GetNext()
		if predicator.CheckCondition(si.condition, tuple) {
			si.currentTuple = tuple
			return tupleToReturn, nil
		}
	}

	//no more tuples meet condition
	si.hasNext = false
	return tupleToReturn, nil
}

func (si *SelectionIterator) HasNext() bool {
	return si.hasNext
}

func (si *SelectionIterator) Close() {
	si.currentTuple = nil
	si.iterator1 = nil
	si.condition = nil
	si.hasNext = true
}

func (si *SelectionIterator) GetSchema() *container.Schema {
	return si.iterator1.GetSchema()
}
