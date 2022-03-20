package execution

import "ZetaDB/container"

type SetDifferenceIterator struct {
}

func (sdi *SetDifferenceIterator) Open(iterator1 *Iterator, iterator2 *Iterator) error {
	return nil
}

func (sdi *SetDifferenceIterator) GetNext() (*container.Tuple, error) {
	return nil, nil
}

func (sdi *SetDifferenceIterator) HasNext() bool {
	return false
}

func (sdi *SetDifferenceIterator) Close() error {
	return nil
}
