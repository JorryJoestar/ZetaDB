package execution

import "ZetaDB/container"

type BagUnionIterator struct {
}

func (bui *BagUnionIterator) Open(iterator1 *Iterator, iterator2 *Iterator) error {
	return nil
}

func (bui *BagUnionIterator) GetNext() (*container.Tuple, error) {
	return nil, nil
}

func (bui *BagUnionIterator) HasNext() bool {
	return false
}

func (bui *BagUnionIterator) Close() error {
	return nil
}
