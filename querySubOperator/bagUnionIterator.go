package execution

import (
	"ZetaDB/container"
	"errors"
)

type BagUnionIterator struct {
	currentTuple *container.Tuple
	iterator1    Iterator
	iterator2    Iterator
	hasNext      bool
}

//BagUnionIterator constructor
func NewBagUnionIterator() *BagUnionIterator {
	bui := &BagUnionIterator{
		hasNext: true}
	return bui
}

//throw error if iterator1 or iterator2 is nil
func (bui *BagUnionIterator) Open(iterator1 Iterator, iterator2 Iterator) error {

	//throw error if iterator1 or iterator2 is nil
	if iterator1 == nil || iterator2 == nil {
		return errors.New("bagUnionIterator.go    Open() iterators invalid")
	}

	bui.iterator1 = iterator1
	bui.iterator2 = iterator2

	if !bui.iterator1.HasNext() && !bui.iterator2.HasNext() { //both two iterators are empty
		bui.hasNext = false
	} else if bui.iterator1.HasNext() { //iterator1 is not empty, first iterate it
		var err error
		bui.currentTuple, err = bui.iterator1.GetNext()
		if err != nil {
			return err
		}
	} else { //iterator1 is empty, iterate iterator2
		var err error
		bui.currentTuple, err = bui.iterator2.GetNext()
		if err != nil {
			return err
		}
	}

	return nil
}

//throw error if HasNext() returns false
func (bui *BagUnionIterator) GetNext() (*container.Tuple, error) {

	//throw error if HasNext() returns false
	if !bui.HasNext() {
		return nil, errors.New("bagUnionIterator.go    GetNext() hasNext false")
	}

	//save currentTuple for return
	tupleToReturn := bui.currentTuple

	//update currentTuple, hasNext
	if !bui.iterator1.HasNext() && !bui.iterator2.HasNext() { //both two iterators are empty
		bui.hasNext = false
	} else if bui.iterator1.HasNext() { //iterator1 is not empty, first iterate it
		var err error
		bui.currentTuple, err = bui.iterator1.GetNext()
		if err != nil {
			return nil, err
		}
	} else { //iterator1 is empty, iterate iterator2
		var err error
		bui.currentTuple, err = bui.iterator2.GetNext()
		if err != nil {
			return nil, err
		}
	}

	return tupleToReturn, nil
}

func (bui *BagUnionIterator) HasNext() bool {
	return bui.hasNext
}

func (bui *BagUnionIterator) Close() {
	bui.currentTuple = nil
	bui.iterator1 = nil
	bui.iterator2 = nil
	bui.hasNext = true
}

func (bui *BagUnionIterator) GetSchema() *container.Schema {
	return bui.iterator1.GetSchema()
}
