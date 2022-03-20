package execution

import "ZetaDB/container"

type ReadFileIterator struct {
}

func (rfi *ReadFileIterator) Open(iterator1 *Iterator, iterator2 *Iterator) error {
	return nil
}

func (rfi *ReadFileIterator) GetNext() (*container.Tuple, error) {
	return nil, nil
}

func (rfi *ReadFileIterator) HasNext() bool {
	return false
}

func (rfi *ReadFileIterator) Close() error {
	return nil
}
