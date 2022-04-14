package execution

import (
	"ZetaDB/container"
	"errors"
)

type BagDifferenceIterator struct {
	nextTuple *container.Tuple

	tuplesIn1NotIn2 map[string]*bagUnionStruct

	tuplesWithIndex map[int]*bagUnionStruct
	nextIndex       int
	maxIndex        int //nextIndex should not large or equal to maxIndex

	iterator1 Iterator
	iterator2 Iterator

	hasNext bool
}

//BagDifferenceIterator constructor
func NewBagDifferenceIterator() *BagDifferenceIterator {
	bdi := &BagDifferenceIterator{
		tuplesIn1NotIn2: make(map[string]*bagUnionStruct),
		nextIndex:       0,
		maxIndex:        0,
		hasNext:         true}
	return bdi
}

//throw error if iterator1 or iterator2 is nil
func (bdi *BagDifferenceIterator) Open(iterator1 Iterator, iterator2 Iterator) error {

	//throw error if iterator1 or iterator2 is nil
	if iterator1 == nil || iterator2 == nil {
		return errors.New("bagDifferenceIterator.go    Open() input iterators invalid")
	}

	//set iterator1 and iterator2
	bdi.iterator1 = iterator1
	bdi.iterator2 = iterator2

	if !bdi.iterator1.HasNext() { //iterator1 is empty, this bdi is also empty
		bdi.hasNext = false
	} else { //iterator1 is not empty

		for { //loop until a over large tuple in iterator1 is discovered, or iterator1 is exhausted

			if !bdi.iterator1.HasNext() { //iterator1 is exhausted
				//begin to use iterator2, neglect over large tuples

				for { //loop until iterator2 exhausted

					if !bdi.iterator2.HasNext() { //create tuplesWithIndex

						bdi.tuplesWithIndex = make(map[int]*bagUnionStruct)
						i := 0
						for _, v := range bdi.tuplesIn1NotIn2 {
							bdi.tuplesWithIndex[i] = v
							i++
							bdi.maxIndex = i
						}

						if len(bdi.tuplesWithIndex) == 0 {
							bdi.hasNext = false
						} else {
							bdi.nextTuple = bdi.tuplesWithIndex[0].tuple
							if bdi.tuplesWithIndex[0].count == 1 {
								bdi.nextIndex = 1
							} else {
								bdi.tuplesWithIndex[0].count--
								bdi.nextIndex = 0
							}
						}
						return nil
					}

					tuple2, error2 := bdi.iterator2.GetNext()
					if error2 != nil {
						return error2
					}

					mapKey2, key2Err := tuple2.TupleGetMapKey()
					if key2Err == nil { //tuple2 should be used to update iterator1Tuples
						if bdi.tuplesIn1NotIn2[mapKey2] != nil {
							if bdi.tuplesIn1NotIn2[mapKey2].count == 1 { //delete this bagUnionStruct from iterator1Tuples
								delete(bdi.tuplesIn1NotIn2, mapKey2)
							} else { //decrease bagUnionStruct count value
								bdi.tuplesIn1NotIn2[mapKey2].count--
							}
						}
					}
				}
			}

			//get a tuple from iterator1
			tuple1, error1 := bdi.iterator1.GetNext()
			if error1 != nil {
				return error1
			}

			mapKey1, key1Err := tuple1.TupleGetMapKey()
			if key1Err != nil { //tuple1 is over large
				bdi.nextTuple = tuple1
				return nil
			} else { //tuple1 should be insert into tuplesIn1NotIn2
				if bdi.tuplesIn1NotIn2[mapKey1] != nil { //duplicated tuple(s) have appeared
					bdi.tuplesIn1NotIn2[mapKey1].count++
				} else {
					bdi.tuplesIn1NotIn2[mapKey1] = &bagUnionStruct{ //discover a new tuple
						tuple: tuple1,
						count: 1}
				}
			}
		}
	}

	return nil
}

//throw error if HasNext() return false
func (bdi *BagDifferenceIterator) GetNext() (*container.Tuple, error) {

	//throw error if HasNext() return false
	if !bdi.HasNext() {
		return nil, errors.New("bagDifferenceIterator.go    GetNext() hasNext false")
	}

	returnTuple := bdi.nextTuple

	if bdi.tuplesWithIndex != nil { //no more over large tuples

		if bdi.nextIndex == bdi.maxIndex { //tuplesWithIndex exhausted
			bdi.hasNext = false
		} else {
			bdi.nextTuple = bdi.tuplesWithIndex[bdi.nextIndex].tuple
			if bdi.tuplesWithIndex[bdi.nextIndex].count == 1 {
				bdi.nextIndex++
			} else {
				bdi.tuplesWithIndex[bdi.nextIndex].count--
			}
		}
	} else { //still some over large tuples in iterator1

		for { //loop until a over large tuple in iterator1 is discovered, or iterator1 is exhausted

			if !bdi.iterator1.HasNext() { //iterator1 is exhausted
				//begin to use iterator2, neglect over large tuples

				for { //loop until iterator2 exhausted
					if !bdi.iterator2.HasNext() { //create tuplesWithIndex

						bdi.tuplesWithIndex = make(map[int]*bagUnionStruct)
						i := 0
						for _, v := range bdi.tuplesIn1NotIn2 {
							bdi.tuplesWithIndex[i] = v
							i++
							bdi.maxIndex = i
						}

						if len(bdi.tuplesWithIndex) == 0 {
							bdi.hasNext = false
						} else {
							bdi.nextTuple = bdi.tuplesWithIndex[0].tuple
							if bdi.tuplesWithIndex[0].count == 1 {
								bdi.nextIndex = 1
							} else {
								bdi.tuplesWithIndex[0].count--
							}
						}
						return returnTuple, nil
					}

					tuple2, error2 := bdi.iterator2.GetNext()
					if error2 != nil {
						return nil, error2
					}

					mapKey2, key2Err := tuple2.TupleGetMapKey()
					if key2Err == nil { //tuple2 should be used to update iterator1Tuples
						if bdi.tuplesIn1NotIn2[mapKey2] != nil {
							if bdi.tuplesIn1NotIn2[mapKey2].count == 1 { //delete this bagUnionStruct from iterator1Tuples
								delete(bdi.tuplesIn1NotIn2, mapKey2)
							} else { //decrease bagUnionStruct count value
								bdi.tuplesIn1NotIn2[mapKey2].count--
							}
						}
					}
				}
			}

			//get a tuple from iterator1
			tuple1, error1 := bdi.iterator1.GetNext()
			if error1 != nil {
				return nil, error1
			}

			mapKey1, key1Err := tuple1.TupleGetMapKey()
			if key1Err != nil { //tuple1 is over large
				bdi.nextTuple = tuple1
				break
			} else { //tuple1 should be insert into iterator1Tuples
				if bdi.tuplesIn1NotIn2[mapKey1] != nil { //duplicated tuple(s) have appeared
					bdi.tuplesIn1NotIn2[mapKey1].count++
				} else {
					bdi.tuplesIn1NotIn2[mapKey1] = &bagUnionStruct{ //discover a new tuple
						tuple: tuple1,
						count: 1}
				}
			}
		}
	}

	return returnTuple, nil
}

func (bdi *BagDifferenceIterator) HasNext() bool {
	return bdi.hasNext
}

func (bdi *BagDifferenceIterator) Close() {
	bdi.nextTuple = nil
	bdi.tuplesIn1NotIn2 = make(map[string]*bagUnionStruct)
	bdi.tuplesWithIndex = nil
	bdi.iterator1 = nil
	bdi.iterator2 = nil
	bdi.hasNext = true
	bdi.nextIndex = 0
	bdi.maxIndex = 0
}

func (bdi *BagDifferenceIterator) GetSchema() *container.Schema {
	return bdi.iterator1.GetSchema()
}

//this structure is used to save info about a tuple and its appearing counts in iterator1
type bagUnionStruct struct {
	tuple *container.Tuple
	count int
}
