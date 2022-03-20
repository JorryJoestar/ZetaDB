package execution

import "ZetaDB/container"

type SetUnionIterator struct {
}

func (sui *SetUnionIterator) Open(iterator1 *Iterator, iterator2 *Iterator) error {
	return nil
}

func (sui *SetUnionIterator) GetNext() (*container.Tuple, error) {
	return nil, nil
}

func (sui *SetUnionIterator) HasNext() bool {
	return false
}

func (sui *SetUnionIterator) Close() error {
	return nil
}
