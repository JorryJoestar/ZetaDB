package container

import "sync"

//Transaction is used to keep changed dataPages & indexPages that
//is about to be pushed into disk as a whole
//Transaction is unique within the entire system
type Transaction struct {
	dataPages  []uint32
	indexPages []uint32
}

//for singleton pattern
var instance *Transaction
var once sync.Once

//to get kernel, call this function
func GetTransaction() *Transaction {
	once.Do(func() {
		instance = &Transaction{
			dataPages:  make([]uint32, 0),
			indexPages: make([]uint32, 0)}
	})
	return instance
}

//if a dataPage is modified and is about to be swapped, insert its pageId into transaction
func (transaction *Transaction) InsertDataPageId(dataPageId uint32) {
	transaction.dataPages = append(transaction.dataPages, dataPageId)
}

//if an indexPage is modified and is about to be swapped, insert its pageId into transaction
func (transaction *Transaction) InsertIndexPageId(indexPageId uint32) {
	transaction.indexPages = append(transaction.indexPages, indexPageId)
}

//after an transaction is completed, initialize it
func (transaction *Transaction) InitializeTransaction() {
	transaction.dataPages = make([]uint32, 0)
	transaction.indexPages = make([]uint32, 0)
}
