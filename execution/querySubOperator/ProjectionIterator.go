package execution

import "ZetaDB/container"

type ProjectionIterator struct {
}

func (pi *ProjectionIterator) Open(iterator1 *Iterator, iterator2 *Iterator) error {
	return nil
}

func (pi *ProjectionIterator) GetNext() (*container.Tuple, error) {
	return nil, nil
}

func (pi *ProjectionIterator) HasNext() bool {
	return false
}

func (pi *ProjectionIterator) Close() error {
	return nil
}
