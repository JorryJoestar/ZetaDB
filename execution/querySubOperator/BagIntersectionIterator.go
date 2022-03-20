package execution

import "ZetaDB/container"

type BagIntersectionIterator struct {
}

func (bii *BagIntersectionIterator) Open(iterator1 *Iterator, iterator2 *Iterator) error {
	return nil
}

func (bii *BagIntersectionIterator) GetNext() (*container.Tuple, error) {
	return nil, nil
}

func (bii *BagIntersectionIterator) HasNext() bool {
	return false
}

func (bii *BagIntersectionIterator) Close() error {
	return nil
}
