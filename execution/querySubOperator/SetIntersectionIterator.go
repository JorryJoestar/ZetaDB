package execution

import (
	"ZetaDB/container"
	"errors"
)

type SetIntersectionIterator struct {
	nextTuple *container.Tuple

	appearedTuples map[string]*container.Tuple

	iterator1 Iterator
	iterator2 Iterator

	hasNext bool
}

func NewSetIntersectionIterator() *SetIntersectionIterator {
	sii := &SetIntersectionIterator{
		appearedTuples: make(map[string]*container.Tuple),
		hasNext:        true}
	return sii
}

//throw error if iterator1 or iterator2 is nil
func (sii *SetIntersectionIterator) Open(iterator1 Iterator, iterator2 Iterator) error {

	//throw error if iterator1 or iterator2 is nil
	if iterator1 == nil || iterator2 == nil {
		return errors.New("setIntersectionIterator.go    Open() iterator invalid")
	}

	sii.iterator1 = iterator1
	sii.iterator2 = iterator2

	if !sii.iterator1.HasNext() { //iterator1 is empty, set hasNext to false
		sii.hasNext = false
	}

	for { //loop until an over large tuple is discovered, or first output in iterator2 is ready, or 1&2 are all exhausted
		if sii.iterator1.HasNext() { //loop iterator1

			tup1, err1 := sii.iterator1.GetNext()
			if err1 != nil {
				return err1
			}

			mapKey1, keyErr1 := tup1.TupleGetMapKey()
			if keyErr1 != nil { //this is an over large tuple
				sii.nextTuple = tup1
				break
			} else { //update appearedTuples if necessary
				if sii.appearedTuples[mapKey1] == nil {
					sii.appearedTuples[mapKey1] = tup1
				}
			}

		} else if sii.iterator2.HasNext() { //loop iterator2

			tup2, err2 := sii.iterator2.GetNext()
			if err2 != nil {
				return err2
			}

			mapKey2, keyErr2 := tup2.TupleGetMapKey()
			if keyErr2 != nil { //this is an over large tuple
				sii.nextTuple = tup2
				break
			} else { //if it appears in appearedTuples, set it as nextTuple, delete it from appearedTuples
				if sii.appearedTuples[mapKey2] != nil {
					delete(sii.appearedTuples, mapKey2)
					sii.nextTuple = tup2
					break
				}
			}

		} else { //both iterator1 & iterator2 are exhausted
			sii.hasNext = false
			break
		}
	}

	return nil
}

//throw error if HasNext() return false
func (sii *SetIntersectionIterator) GetNext() (*container.Tuple, error) {

	//throw error if HasNext() return false
	if !sii.HasNext() {
		return nil, errors.New("setUnionIterator.go    GetNext() HasNext invalid")
	}

	returnTuple := sii.nextTuple

	//update nextTuple, hasNext, appearedTuples
	for { //loop until hasNext set to false, or nextTuple updated
		if sii.iterator1.HasNext() {
			tup1, err1 := sii.iterator1.GetNext()
			if err1 != nil {
				return nil, err1
			}

			tupKey, keyErr := tup1.TupleGetMapKey()

			if keyErr == nil { // this is not an over large tuple
				if sii.appearedTuples[tupKey] == nil { //this tuple has not appeared
					sii.appearedTuples[tupKey] = tup1
				}
			} else { //this tuple is an over large tuple
				sii.nextTuple = tup1
				break
			}
		} else if sii.iterator2.HasNext() { //iterator1 is empty, iterator2 is not empty
			tup2, err2 := sii.iterator2.GetNext()
			if err2 != nil {
				return nil, err2
			}

			tupKey, keyErr := tup2.TupleGetMapKey()

			if keyErr == nil { // this is not an over large tuple
				if sii.appearedTuples[tupKey] != nil { //this tuple has not appeared
					delete(sii.appearedTuples, tupKey)
					sii.nextTuple = tup2
					break
				}
			} else { //this tuple is an over large tuple
				sii.nextTuple = tup2
				break
			}
		} else { //both iterator1 and iterator2 is empty
			sii.hasNext = false
			break
		}
	}

	return returnTuple, nil
}

func (sii *SetIntersectionIterator) HasNext() bool {
	return sii.hasNext
}

func (sii *SetIntersectionIterator) Close() {
	sii.nextTuple = nil
	sii.appearedTuples = make(map[string]*container.Tuple)
	sii.iterator1 = nil
	sii.iterator2 = nil
	sii.hasNext = true
}

func (sii *SetIntersectionIterator) GetSchema() *container.Schema {
	return sii.iterator1.GetSchema()
}
