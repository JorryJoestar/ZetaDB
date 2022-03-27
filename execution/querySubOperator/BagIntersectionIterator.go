package execution

import (
	"ZetaDB/container"
	"errors"
)

type BagIntersectionIterator struct {
	nextTuple *container.Tuple

	tuplesIn1 map[string]*bagIntersectionStruct

	iterator1 Iterator
	iterator2 Iterator

	hasNext bool
}

//BagIntersectionIterator constructor
func NewBagIntersectionIterator() *BagIntersectionIterator {
	bii := &BagIntersectionIterator{
		tuplesIn1: make(map[string]*bagIntersectionStruct),
		hasNext:   true}
	return bii
}

//throw error if iterator1 or iterator2 is nil
func (bii *BagIntersectionIterator) Open(iterator1 Iterator, iterator2 Iterator) error {

	//throw error if iterator1 or iterator2 is nil
	if iterator1 == nil || iterator2 == nil {
		return errors.New("bagIntersectionIterator.go    Open() input iterators invalid")
	}

	bii.iterator1 = iterator1
	bii.iterator2 = iterator2

	if !bii.iterator1.HasNext() { //iterator1 is empty, set hasNext to false
		bii.hasNext = false
	} else { //iterator1 is not empty

		for { //loop until a over large tuple in iterator1 is discovered, or iterator1 is exhausted

			if !bii.iterator1.HasNext() { //iterator1 is exhausted
				//begin to use iterator2, until discover a over large tuple or iterator2 is exhausted

				for {
					if !bii.iterator2.HasNext() { //iterator2 is exhausted
						bii.hasNext = false
						return nil
					}

					tuple2, error2 := bii.iterator2.GetNext()
					if error2 != nil {
						return error2
					}

					mapKey2, key2Err := tuple2.TupleGetMapKey()
					if key2Err != nil { // this is an over large tuple,set it as nextTuple
						bii.nextTuple = tuple2
						return nil
					} else { //check if it is in tuplesIn1, if it is in set it as nextTuple and update tuplesIn1
						if bii.tuplesIn1[mapKey2] != nil {

							//update tuplesIn1
							if bii.tuplesIn1[mapKey2].count == 1 { //delete this bagUnionStruct from iterator1Tuples
								delete(bii.tuplesIn1, mapKey2)
							} else { //decrease bagUnionStruct count value
								bii.tuplesIn1[mapKey2].count--
							}

							//set nextTuple
							bii.nextTuple = tuple2
							return nil
						}
					}
				}

			}

			//get a tuple from iterator1
			tuple1, error1 := bii.iterator1.GetNext()
			if error1 != nil {
				return error1
			}

			mapKey1, key1Err := tuple1.TupleGetMapKey()
			if key1Err != nil { //tuple1 is over large
				bii.nextTuple = tuple1
				return nil
			} else { //tuple1 should be insert into tuplesIn1
				if bii.tuplesIn1[mapKey1] != nil { //duplicated tuple(s) have appeared
					bii.tuplesIn1[mapKey1].count++
				} else {
					bii.tuplesIn1[mapKey1] = &bagIntersectionStruct{ //discover a new tuple
						tuple: tuple1,
						count: 1}
				}
			}
		}

	}

	return nil
}

//throw error if HasNext() return false
func (bii *BagIntersectionIterator) GetNext() (*container.Tuple, error) {

	//throw error if HasNext() return false
	if !bii.HasNext() {
		return nil, errors.New("bagIntersectionIterator.go    GetNext() hasNext false")
	}

	if !bii.iterator2.HasNext() { //iterator2 is exhausted
		bii.hasNext = false
	}

	returnTuple := bii.nextTuple

	for { //loop until iterator1 is exhausted or an over large tuple is discovered

		if !bii.iterator1.HasNext() { //iterator1 exhausted
			break
		}

		//get a tuple from iterator1
		tuple1, error1 := bii.iterator1.GetNext()
		if error1 != nil {
			return nil, error1
		}

		mapKey1, key1Err := tuple1.TupleGetMapKey()
		if key1Err != nil { //tuple1 is over large
			bii.nextTuple = tuple1
			return returnTuple, nil
		} else { //tuple1 should be insert into tuplesIn1
			if bii.tuplesIn1[mapKey1] != nil { //duplicated tuple(s) have appeared
				bii.tuplesIn1[mapKey1].count++
			} else {
				bii.tuplesIn1[mapKey1] = &bagIntersectionStruct{ //discover a new tuple
					tuple: tuple1,
					count: 1}
			}
		}
	}

	for { //loop until an over large tuple is discovered, or iterator2 is exhausted

		if !bii.iterator2.HasNext() { //iterator2 is exhausted
			break
		}

		tuple2, error2 := bii.iterator2.GetNext()

		if error2 != nil {
			return nil, error2
		}

		mapKey2, key2Err := tuple2.TupleGetMapKey()
		if key2Err != nil { // this is an over large tuple,set it as nextTuple
			bii.nextTuple = tuple2
			break
		} else { //check if it is in tuplesIn1, if it is in set it as nextTuple and update tuplesIn1
			if bii.tuplesIn1[mapKey2] != nil {

				//update tuplesIn1
				if bii.tuplesIn1[mapKey2].count == 1 { //delete this bagUnionStruct from iterator1Tuples
					delete(bii.tuplesIn1, mapKey2)
				} else { //decrease bagUnionStruct count value
					bii.tuplesIn1[mapKey2].count--
				}

				//set nextTuple
				bii.nextTuple = tuple2
				break
			}
		}
	}

	return returnTuple, nil
}

func (bii *BagIntersectionIterator) HasNext() bool {
	return bii.hasNext
}

func (bii *BagIntersectionIterator) Close() {
	bii.nextTuple = nil
	bii.tuplesIn1 = make(map[string]*bagIntersectionStruct)
	bii.iterator1 = nil
	bii.iterator2 = nil
	bii.hasNext = true
}

func (bii *BagIntersectionIterator) GetSchema() *container.Schema {
	return bii.iterator1.GetSchema()
}

type bagIntersectionStruct struct {
	tuple *container.Tuple
	count int
}
