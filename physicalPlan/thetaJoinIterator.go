package execution

import "ZetaDB/container"

type ThetaJoinIterator struct {
}

func (tji *ThetaJoinIterator) Open(iterator1 *Iterator, iterator2 *Iterator) error {
	return nil
}

func (tji *ThetaJoinIterator) GetNext() (*container.Tuple, error) {
	return nil, nil
}

func (tji *ThetaJoinIterator) HasNext() bool {
	return false
}

func (tji *ThetaJoinIterator) Close() error {
	return nil
}
