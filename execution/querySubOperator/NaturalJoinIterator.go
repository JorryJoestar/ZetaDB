package execution

import "ZetaDB/container"

type NaturalJoinIterator struct {
}

func (nji *NaturalJoinIterator) Open(iterator1 *Iterator, iterator2 *Iterator) error {
	return nil
}

func (nji *NaturalJoinIterator) GetNext() (*container.Tuple, error) {
	return nil, nil
}

func (nji *NaturalJoinIterator) HasNext() bool {
	return false
}

func (nji *NaturalJoinIterator) Close() error {
	return nil
}
