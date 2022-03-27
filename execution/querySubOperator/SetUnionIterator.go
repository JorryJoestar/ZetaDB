package execution

import (
	"ZetaDB/container"
	"errors"
)

type SetUnionIterator struct {
	nextTuple *container.Tuple

	appearedTuples map[string]*container.Tuple

	iterator1 Iterator
	iterator2 Iterator

	hasNext bool
}

//SetUnionIterator constructor
func NewSetUnionIterator() *SetUnionIterator {
	sui := &SetUnionIterator{
		appearedTuples: make(map[string]*container.Tuple),
		hasNext:        true}
	return sui
}

//throw error if iterator1 or iterator2 is nil
func (sui *SetUnionIterator) Open(iterator1 Iterator, iterator2 Iterator) error {

	//throw error if iterator1 or iterator2 is nil
	if iterator1 == nil || iterator2 == nil {
		return errors.New("setUnionIterator.go    Open() iterator invalid")
	}

	sui.iterator1 = iterator1
	sui.iterator2 = iterator2

	if sui.iterator1.HasNext() {
		tup1, err1 := sui.iterator1.GetNext()
		if err1 != nil {
			return err1
		}

		sui.nextTuple = tup1
		tupKey, keyErr := tup1.TupleGetMapKey()

		if keyErr == nil { // this is not an over large tuple
			//insert this tuple into appearedTuples
			sui.appearedTuples[tupKey] = tup1
		}
	} else if sui.iterator2.HasNext() { //iterator1 is empty, iterator2 is not empty
		tup2, err2 := sui.iterator2.GetNext()
		if err2 != nil {
			return err2
		}

		sui.nextTuple = tup2
		tupKey, keyErr := tup2.TupleGetMapKey()

		if keyErr == nil { // this is not an over large tuple
			//insert this tuple into appearedTuples
			sui.appearedTuples[tupKey] = tup2
		}
	} else { //both iterator1 and iterator2 is empty
		sui.hasNext = false
	}

	return nil
}

//throw error if HasNext() return false
func (sui *SetUnionIterator) GetNext() (*container.Tuple, error) {

	//throw error if HasNext() return false
	if !sui.HasNext() {
		return nil, errors.New("setUnionIterator.go    GetNext() HasNext invalid")
	}

	returnTuple := sui.nextTuple

	//update nextTuple, hasNext, appearedTuples
	for { //loop until hasNext set to false, or nextTuple updated
		if sui.iterator1.HasNext() {
			tup1, err1 := sui.iterator1.GetNext()
			if err1 != nil {
				return nil, err1
			}

			tupKey, keyErr := tup1.TupleGetMapKey()

			if keyErr == nil { // this is not an over large tuple
				if sui.appearedTuples[tupKey] == nil { //this tuple has not appeared
					sui.appearedTuples[tupKey] = tup1
					sui.nextTuple = tup1
					break
				}
			} else { //this tuple is an over large tuple
				sui.nextTuple = tup1
				break
			}
		} else if sui.iterator2.HasNext() { //iterator1 is empty, iterator2 is not empty
			tup2, err2 := sui.iterator2.GetNext()
			if err2 != nil {
				return nil, err2
			}

			tupKey, keyErr := tup2.TupleGetMapKey()

			if keyErr == nil { // this is not an over large tuple
				if sui.appearedTuples[tupKey] == nil { //this tuple has not appeared
					sui.appearedTuples[tupKey] = tup2
					sui.nextTuple = tup2
					break
				}
			} else { //this tuple is an over large tuple
				sui.nextTuple = tup2
				break
			}
		} else { //both iterator1 and iterator2 is empty
			sui.hasNext = false
			break
		}
	}

	return returnTuple, nil
}

func (sui *SetUnionIterator) HasNext() bool {
	return sui.hasNext
}

func (sui *SetUnionIterator) Close() {
	sui.nextTuple = nil
	sui.appearedTuples = make(map[string]*container.Tuple)
	sui.iterator1 = nil
	sui.iterator2 = nil
	sui.hasNext = true
}

func (sui *SetUnionIterator) GetSchema() *container.Schema {
	return sui.iterator1.GetSchema()
}
