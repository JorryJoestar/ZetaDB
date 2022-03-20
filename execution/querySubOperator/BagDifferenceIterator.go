package execution

import "ZetaDB/container"

type BagDifferenceIterator struct {
}

func (bdi *BagDifferenceIterator) Open(iterator1 *Iterator, iterator2 *Iterator) error {
	return nil
}

func (bdi *BagDifferenceIterator) GetNext() (*container.Tuple, error) {
	return nil, nil
}

func (bdi *BagDifferenceIterator) HasNext() bool {
	return false
}

func (bdi *BagDifferenceIterator) Close() error {
	return nil
}
