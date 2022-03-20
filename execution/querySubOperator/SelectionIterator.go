package execution

import "ZetaDB/container"

type SelectionIterator struct {
}

func (si *SelectionIterator) Open(iterator1 *Iterator, iterator2 *Iterator) error {
	return nil
}

func (si *SelectionIterator) GetNext() (*container.Tuple, error) {
	return nil, nil
}

func (si *SelectionIterator) HasNext() bool {
	return false
}

func (si *SelectionIterator) Close() error {
	return nil
}
