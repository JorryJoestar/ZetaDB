package execution

import "ZetaDB/container"

type SetIntersectionIterator struct {
}

func (sii *SetIntersectionIterator) Open(iterator1 *Iterator, iterator2 *Iterator) error {
	return nil
}

func (sii *SetIntersectionIterator) GetNext() (*container.Tuple, error) {
	return nil, nil
}

func (sii *SetIntersectionIterator) HasNext() bool {
	return false
}

func (sii *SetIntersectionIterator) Close() error {
	return nil
}
