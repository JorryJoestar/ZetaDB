package execution

import "ZetaDB/container"

type ProductIterator struct {
}

func (pi *ProductIterator) Open(iterator1 *Iterator, iterator2 *Iterator) error {
	return nil
}

func (pi *ProductIterator) GetNext() (*container.Tuple, error) {
	return nil, nil
}

func (pi *ProductIterator) HasNext() bool {
	return false
}

func (pi *ProductIterator) Close() error {
	return nil
}
