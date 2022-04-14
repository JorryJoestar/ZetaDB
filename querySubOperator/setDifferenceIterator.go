package execution

import (
	"ZetaDB/container"
	"errors"
)

type SetDifferenceIterator struct {
	nextTuple *container.Tuple

	tuplesIn1NotIn2 map[string]*container.Tuple

	tuplesWithIndex map[int]*container.Tuple
	nextIndex       int
	maxIndex        int //nextIndex should not large or equal to maxIndex

	iterator1 Iterator
	iterator2 Iterator

	hasNext bool
}

func NewSetDifferenceIterator() *SetDifferenceIterator {
	sdi := &SetDifferenceIterator{
		tuplesIn1NotIn2: make(map[string]*container.Tuple),
		nextIndex:       0,
		maxIndex:        0,
		hasNext:         true}
	return sdi
}

//throw error if iterator1 or iterator2 is nil
func (sdi *SetDifferenceIterator) Open(iterator1 Iterator, iterator2 Iterator) error {

	//throw error if iterator1 or iterator2 is nil
	if iterator1 == nil || iterator2 == nil {
		return errors.New("setDifferenceIterator.go    Open() input iterators invalid")
	}

	//set iterator1 and iterator2
	sdi.iterator1 = iterator1
	sdi.iterator2 = iterator2

	if !sdi.iterator1.HasNext() { //iterator1 is empty, set hasNext to false
		sdi.hasNext = false
	}

	//loop until an over large tuple is discovered, or tuplesWithIndex is ready, or 1&2 are all exhausted
	for {
		if sdi.iterator1.HasNext() { //loop iterator1

			tup1, err1 := sdi.iterator1.GetNext()
			if err1 != nil {
				return err1
			}

			mapKey1, keyErr1 := tup1.TupleGetMapKey()
			if keyErr1 != nil { //an over large tuple
				sdi.nextTuple = tup1
				break
			} else { //update appearedTuples if necessary
				if sdi.tuplesIn1NotIn2[mapKey1] == nil {
					sdi.tuplesIn1NotIn2[mapKey1] = tup1
				}
			}

		} else if sdi.iterator2.HasNext() { //loop iterator2

			tup2, err2 := sdi.iterator2.GetNext()
			if err2 != nil {
				return err2
			}

			mapKey2, keyErr2 := tup2.TupleGetMapKey()
			if keyErr2 != nil { //an over large tuple
				sdi.nextTuple = tup2
				break
			} else {
				if sdi.tuplesIn1NotIn2[mapKey2] != nil {
					delete(sdi.tuplesIn1NotIn2, mapKey2)
				}
			}

		} else { //iterator1 & iterator2 are exhausted, set tuplesWithIndex
			if len(sdi.tuplesIn1NotIn2) == 0 { //this setDifferenceIt is empty
				sdi.hasNext = false
				break
			} else {
				sdi.tuplesWithIndex = make(map[int]*container.Tuple)
				i := 0
				for _, v := range sdi.tuplesIn1NotIn2 {
					sdi.tuplesWithIndex[i] = v
					i++
					sdi.maxIndex = i
				}

				sdi.nextIndex = 1
				sdi.nextTuple = sdi.tuplesWithIndex[0]
				break
			}
		}
	}

	return nil
}

//throw error if HasNext() return false
func (sdi *SetDifferenceIterator) GetNext() (*container.Tuple, error) {

	//throw error if HasNext() return false
	if !sdi.HasNext() {
		return nil, errors.New("bagDifferenceIterator.go    GetNext() hasNext false")
	}

	returnTuple := sdi.nextTuple

	for {
		if sdi.iterator1.HasNext() { //loop iterator1

			tup1, err1 := sdi.iterator1.GetNext()
			if err1 != nil {
				return nil, err1
			}

			mapKey1, keyErr1 := tup1.TupleGetMapKey()
			if keyErr1 != nil { //an over large tuple
				sdi.nextTuple = tup1
				break
			} else { //update appearedTuples if necessary
				if sdi.tuplesIn1NotIn2[mapKey1] == nil {
					sdi.tuplesIn1NotIn2[mapKey1] = tup1
				}
			}

		} else if sdi.iterator2.HasNext() { //loop iterator2

			tup2, err2 := sdi.iterator2.GetNext()
			if err2 != nil {
				return nil, err2
			}

			mapKey2, keyErr2 := tup2.TupleGetMapKey()
			if keyErr2 != nil { //an over large tuple
				sdi.nextTuple = tup2
				break
			} else {
				if sdi.tuplesIn1NotIn2[mapKey2] != nil {
					delete(sdi.tuplesIn1NotIn2, mapKey2)
				}
			}
		} else { //iterator1 & iterator2 are exhausted, set tuplesWithIndex
			if sdi.tuplesWithIndex != nil { //tuplesWithIndex is already generated
				if sdi.nextIndex == sdi.maxIndex { //tuplesWithIndex exhausted
					sdi.hasNext = false
					break
				} else {
					sdi.nextTuple = sdi.tuplesWithIndex[sdi.nextIndex]
					sdi.nextIndex++
					break
				}
			} else { //tuplesWithIndex has not been generated
				if len(sdi.tuplesIn1NotIn2) == 0 { //this setDifferenceIt is empty
					sdi.hasNext = false
					break
				} else {
					sdi.tuplesWithIndex = make(map[int]*container.Tuple)
					i := 0
					for _, v := range sdi.tuplesIn1NotIn2 {
						sdi.tuplesWithIndex[i] = v
						i++
						sdi.maxIndex = i
					}

					sdi.nextIndex = 1
					sdi.nextTuple = sdi.tuplesWithIndex[0]
					break
				}
			}
		}
	}

	return returnTuple, nil
}

func (sdi *SetDifferenceIterator) HasNext() bool {
	return sdi.hasNext
}

func (sdi *SetDifferenceIterator) Close() {
	sdi.nextTuple = nil
	sdi.tuplesIn1NotIn2 = make(map[string]*container.Tuple)
	sdi.tuplesWithIndex = nil
	sdi.iterator1 = nil
	sdi.iterator2 = nil
	sdi.hasNext = true
	sdi.nextIndex = 0
	sdi.maxIndex = 0
}

func (sdi *SetDifferenceIterator) GetSchema() *container.Schema {
	return sdi.iterator1.GetSchema()
}
