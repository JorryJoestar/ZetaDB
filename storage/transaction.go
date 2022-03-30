package storage

import (
	"ZetaDB/container"
	"ZetaDB/utility"
	"sync"
)

//Transaction is used to keep changed dataPages & indexPages that
//is about to be pushed into disk as a whole
//Transaction is unique within the entire system
type Transaction struct {
	dataPages  map[uint32]*DataPage
	indexPages map[uint32]*IndexPage
}

//for singleton pattern
var instance *Transaction
var once sync.Once

//to get kernel, call this function
func GetTransaction() *Transaction {
	once.Do(func() {
		instance = &Transaction{
			dataPages:  make(map[uint32]*DataPage),
			indexPages: make(map[uint32]*IndexPage)}
	})
	return instance
}

//if a dataPage is modified and is about to be swapped, insert its pageId into transaction
func (transaction *Transaction) InsertDataPage(dataPage *DataPage) {
	pageId := dataPage.DpGetPageId()
	transaction.dataPages[pageId] = dataPage
}

//if an indexPage is modified and is about to be swapped, insert its pageId into transaction
func (transaction *Transaction) InsertIndexPage(indexPage *IndexPage) {
	pageId := indexPage.IndexPageGetPageId()
	transaction.indexPages[pageId] = indexPage
}

//check if last transaction succeed
//if flag is true, there is stored and unfinished transaction, try to do it again, if suceed, set flag to false
//if flag is false, do nothing
func (transaction *Transaction) Recovery() {
	se := GetStorageEngine()

	//get flag, if no such flag, just return
	flag, err := se.getLogFlag()
	if err != nil {
		return
	}

	//if flag is false, do nothing
	if !flag {
		return
	}

	//fetch all dataPages & indexPages, push them into dataFile & indexFile
	var logPageBytes []byte
	var logPage *logPage
	var logPageId uint32 = 1 // first log page is at 1

	for {
		logPageBytes, err = se.getPageBytesFromLogFile(logPageId)
		if err != nil {
			return
		}
		logPage, err = NewLogPageFromBytes(logPageBytes)
		if err != nil {
			return
		}

		for _, log := range logPage.LogPageGetLogs() {
			mode := log.LogGetFileMode()
			logPageId := log.LogGetlogPageId()
			objectPageId := log.LogGetobjectPageId()

			bytes, err := se.iom.BytesFromLogFile(logPageId*uint32(utility.DEFAULT_PAGE_SIZE), utility.DEFAULT_PAGE_SIZE)
			if err != nil {
				return
			}

			if mode == 1 { //dataPage
				se.dataPageBytesToDataFile(bytes, objectPageId)
			} else if mode == 2 { //indexPage
				se.indexPageBytesToIndexFile(bytes, objectPageId)
			}
		}

		if logPage.LogPageGetLogNextPageId() == logPage.LogPageGetLogPageId() { //reach the end
			break
		} else { //change logPageId to next page
			logPageId = logPage.LogPageGetLogNextPageId()
		}
	}

	//set flag to false
	se.setLogFlag(false)
}

//push current pages belong to transaction into logfile
//set flag to true (first byte of first logPage is 0b11111111)
//set all page to unmodified
//call Recovery()
//initialize transaction
func (transaction *Transaction) PushTransactionIntoDisk() {
	se := GetStorageEngine()

	//push current pages belong to transaction into logfile
	var currentLogPageId uint32 = 1
	var currentSavePageId uint32 = 2
	currentLogPage := NewLogPage(currentLogPageId, currentLogPageId, currentLogPageId)

	for _, dataPage := range transaction.dataPages {
		if currentLogPage.LogPageIsFull() { //currentLogPage is full, push it into disk, create a new one

			//change currentLogPageId to currentSavePageId, increase currentSavePageId
			currentLogPageId = currentSavePageId
			currentSavePageId++

			//create a new logPage
			nextLogPage := NewLogPage(currentLogPageId, currentLogPage.LogPageGetLogPageId(), currentLogPageId)

			//change nextPageId of logPage which is ready to be swapped
			currentLogPage.LogPageSetLogNextPageId(currentLogPageId)

			//swap currentLogPage
			se.swapPageBytesIntoLogFile(currentLogPage.LogPageToBytes(), currentLogPage.LogPageGetLogPageId())

			//set nextLogPage as currentLogPage
			currentLogPage = nextLogPage
		}

		//set dataPage to unmodified
		dataPage.UnmodifyDataPage()

		//create a log for page that is ready to be pushed
		newLog, _ := container.NewLog(1, currentSavePageId, dataPage.DpGetPageId())

		//insert this log into logPage
		currentLogPage.LogPageInsertLog(newLog)

		//swap this dataPage into disk
		dataPageBytes, _ := dataPage.DataPageToBytes()
		se.swapPageBytesIntoLogFile(dataPageBytes, currentSavePageId)

		//update currentSavePageId
		currentSavePageId++
	}

	for _, indexPage := range transaction.indexPages {
		if currentLogPage.LogPageIsFull() { //currentLogPage is full, push it into disk, create a new one

			//change currentLogPageId to currentSavePageId, increase currentSavePageId
			currentLogPageId = currentSavePageId
			currentSavePageId++

			//create a new logPage
			nextLogPage := NewLogPage(currentLogPageId, currentLogPage.LogPageGetLogPageId(), currentLogPageId)

			//change nextPageId of logPage which is ready to be swapped
			currentLogPage.LogPageSetLogNextPageId(currentLogPageId)

			//swap currentLogPage
			se.swapPageBytesIntoLogFile(currentLogPage.LogPageToBytes(), currentLogPage.LogPageGetLogPageId())

			//set nextLogPage as currentLogPage
			currentLogPage = nextLogPage
		}

		//set dataPage to unmodified
		indexPage.IndexPageUnModify()

		//create a log for page that is ready to be pushed
		newLog, _ := container.NewLog(2, currentSavePageId, indexPage.IndexPageGetPageId())

		//insert this log into logPage
		currentLogPage.LogPageInsertLog(newLog)

		//swap this indexPage into disk
		se.swapPageBytesIntoLogFile(indexPage.IndexPageToBytes(), currentSavePageId)

		//update currentSavePageId
		currentSavePageId++
	}
	//swap the last logPage
	se.swapPageBytesIntoLogFile(currentLogPage.LogPageToBytes(), currentLogPage.LogPageGetLogPageId())

	//set flag to true (first byte of first logPage is 0b11111111)
	se.setLogFlag(true)

	//call Recovery()
	transaction.Recovery()

	//initialize transaction
	transaction.dataPages = make(map[uint32]*DataPage)
	transaction.indexPages = make(map[uint32]*IndexPage)
}
