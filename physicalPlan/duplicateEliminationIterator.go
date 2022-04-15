package execution

import (
	"ZetaDB/container"
	"errors"
)

//this iterator is used to delete all duplicate tuples
type DuplicateEliminationIterator struct {
	hasNext        bool
	appearedTuples map[string]*container.Tuple
	currentTuple   *container.Tuple
	inputIterator  Iterator
}

//DuplicateEliminationIterator constructor
func NewDuplicateEliminationIterator() *DuplicateEliminationIterator {
	return &DuplicateEliminationIterator{
		hasNext:        true,
		appearedTuples: make(map[string]*container.Tuple)}
}

//iterator2 should always be nil
//throw error if iterator2 is not nil or iterator1 is nil
func (dei *DuplicateEliminationIterator) Open(iterator1 Iterator, iterator2 Iterator) error {

	//throw error if iterator2 is not nil or iterator1 is nil
	if iterator1 == nil || iterator2 != nil {
		return errors.New("duplicateEliminationIterator.go    Open() input iterators invalid")
	}

	dei.inputIterator = iterator1
	if !dei.inputIterator.HasNext() { //input iterator is empty
		dei.hasNext = false
	} else { //set currentTuple and push it into appearedTuples
		dei.currentTuple, _ = dei.inputIterator.GetNext()
		currentMapKey, keyErr := dei.currentTuple.TupleGetMapKey()
		if keyErr == nil { //currentMapKey generated
			dei.appearedTuples[currentMapKey] = dei.currentTuple
		}
	}

	return nil
}

//throw error if HasNext() is false
func (dei *DuplicateEliminationIterator) GetNext() (*container.Tuple, error) {
	//throw error if HasNext() is false
	if !dei.HasNext() {
		return nil, errors.New("duplicateEliminationIterator.go    GetNext() hasNext false")
	}

	//save currentTuple, ready to return it
	currentTup := dei.currentTuple

	//update hasNext, appearedTuples, currentTuple
	for { //loop until no more tuples from inputIterator or find a tuple that has not appeared yet
		if !dei.inputIterator.HasNext() { //no more tuples
			dei.hasNext = false
			break
		}

		inputTuple, inputErr := dei.inputIterator.GetNext()
		if inputErr != nil {
			return nil, inputErr
		}

		mapKey, mapKeyErr := inputTuple.TupleGetMapKey()
		if mapKeyErr == nil { //mapKey generated, check if it is already appeared in appearedTuples
			if dei.appearedTuples[mapKey] == nil { //this page has not appeared yet
				dei.appearedTuples[mapKey] = inputTuple
				dei.currentTuple = inputTuple
				break
			}
		} else { //inputTuple is too long
			dei.currentTuple = inputTuple
			break
		}
	}

	return currentTup, nil
}

func (dei *DuplicateEliminationIterator) HasNext() bool {
	return dei.hasNext
}

func (dei *DuplicateEliminationIterator) Close() {
	dei.hasNext = true
	dei.appearedTuples = make(map[string]*container.Tuple)
	dei.currentTuple = nil
}

func (dei *DuplicateEliminationIterator) GetSchema() *container.Schema {
	return dei.inputIterator.GetSchema()
}
