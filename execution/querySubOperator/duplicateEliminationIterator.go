package execution

import "ZetaDB/container"

//this iterator is used to delete all duplicate tuples
type DuplicateEliminationIterator struct {
	hasNext        bool
	appearedTuples map[string]*container.Tuple
}

//DuplicateEliminationIterator constructor
func NewDuplicateEliminationIterator() {

}

//iterator2 should always be nil
//throw error if iterator2 is not nil
func (dei *DuplicateEliminationIterator) Open(iterator1 *Iterator, iterator2 *Iterator) error {
	return nil
}

func (dei *DuplicateEliminationIterator) GetNext() (*container.Tuple, error) {
	return nil, nil
}

func (dei *DuplicateEliminationIterator) HasNext() bool {
	return false
}

func (dei *DuplicateEliminationIterator) Close() error {
	return nil
}
