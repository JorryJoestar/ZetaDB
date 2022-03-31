storage/indexPage.go
    mode1 page: Modification on elements and pointerPageId would not influent pointerNum, always remember to update pointerNum if necessary.
    
    mode2 page: In order to insert a new indexRecord, use IndexPageInsertNewIndexRecord(), this will increase the length of records automatically. when calling IndexPageSetIndexRecordAt(), it is invalid to modify a position that have not contained an indexRecord before.

    mode3 page is similar to mode2

storage/storageEngine.go
    data buffer: for page whose pageId is under 17, they are stored in keyTableHeadPageBuffer and it is invalid to delete them.

execution/querySubOperator
    after tuples going through these operators, their tableId & tupleId is invalid

execution/initializationManager.go
    always remember before initialization, old files should be removed

container/logicalPlan.go
    all nodes are bag operation defaultly

storage
    valid method:
        storageEngine
            GetStorageEngine() *StorageEngine

            GetDataPage(pageId uint32, schema *Schema) (*DataPage, error)
            GetIndexPage(pageId uint32) (*IndexPage, error)

            InsertDataPage(page *DataPage) error
            InsertIndexPage(page *IndexPage) error

            EraseDataFile() error
            EraseIndexFile() error
            EraseLogFile() error

        transaction
            GetTransaction() *Transaction

            InsertDataPage(dataPage *DataPage)
            InsertIndexPage(indexPage *IndexPage)
            Recovery()
            PushTransactionIntoDisk()
        
        note:
            1. call InsertDataPage() & InsertIndexPage() of transaction, if these pages are to be pushed into disk
            2. always call PushTransactionIntoDisk() after each transaction, no matter whether InsertDataPage() or InsertIndexPage() is called before
            3.everytime the system boosters, call Recovery()
            4.whenever a page is newly created, insert it into buffers

execution/querySubOperator
    remember open these iterator
