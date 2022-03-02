package storage

type indexBuffer struct{}

//in order to fetch a indexBuffer, call this function
func GetIndexBuffer() *indexBuffer {
	ib := &indexBuffer{}
	return ib
}
