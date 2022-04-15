package execution

import (
	"ZetaDB/container"
	"ZetaDB/parser"
	its "ZetaDB/physicalPlan"
	"ZetaDB/storage"
	. "ZetaDB/utility"
	"errors"
	"sync"
)

type KeytableManager struct {
}

//for singleton pattern
var ktmInstance *KeytableManager
var ktmOnce sync.Once

//to get KeytableManager, call this function
func GetKeytableManager() *KeytableManager {
	ktmOnce.Do(func() {
		ktmInstance = &KeytableManager{}
	})

	return ktmInstance
}

//from utility instead of from disk
//ignore if tableId >16
func (ktm *KeytableManager) GetKeyTableSchema(tableId uint32) *container.Schema {
	rewriter := GetRewriter()
	parser := parser.GetParser()

	//throw error if tableId >16
	if tableId > 16 {
		return nil
	}

	ast, _ := parser.ParseSql(DEFAULT_KEYTABLES_SCHEMA[tableId])
	schema, _ := rewriter.ASTNodeToSchema(ast)

	return schema
}

//get tableId & schema from dataFile according to tableName, k_tableId_schema table 8
//throw error if no such table
func (ktm *KeytableManager) Query_k_tableId_schema_FromTableName(tableName string) (uint32, *container.Schema, error) {
	astOfCreateTable8, _ := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[8])
	schemaOfCreateTable8, err := GetRewriter().ASTNodeToSchema(astOfCreateTable8)
	if err != nil {
		return 0, nil, err
	}

	seqIt := its.NewSequentialFileReaderIterator(8, schemaOfCreateTable8)
	seqIt.Open(nil, nil)
	for seqIt.HasNext() {
		table8Tuple, err := seqIt.GetNext()
		if err != nil {
			return 0, nil, err
		}

		//TODO mantain: if schema of table 8 is changed, this index could be asked to change
		tableIdBytes, err := table8Tuple.TupleGetFieldValue(0)
		if err != nil {
			return 0, nil, err
		}
		tableId, err := BytesToINT(tableIdBytes)
		if err != nil {
			return 0, nil, err
		}
		schemaStringBytes, err := table8Tuple.TupleGetFieldValue(1)
		if err != nil {
			return 0, nil, err
		}
		schemaString, err := BytesToVARCHAR(schemaStringBytes)
		if err != nil {
			return 0, nil, err
		}

		//parse this string to get schema
		ast, _ := parser.GetParser().ParseSql(schemaString)
		currentSchema, err := GetRewriter().ASTNodeToSchema(ast)
		if err != nil {
			return 0, nil, err
		}

		//found correct table schema
		if currentSchema.GetSchemaTableName() == tableName {
			return uint32(tableId), currentSchema, nil
		}

	}
	return 0, nil, errors.New("execution/keyTableManager.go    GetTableInfo() no such table")
}

//get table schema from its tableId
func (ktm *KeytableManager) Query_k_tableId_schema_FromTableId(tableId uint32) (*container.Schema, error) {

	astOfCreateTable8, _ := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[8])
	schemaOfCreateTable8, err := GetRewriter().ASTNodeToSchema(astOfCreateTable8)
	if err != nil {
		return nil, err
	}

	seqIt := its.NewSequentialFileReaderIterator(8, schemaOfCreateTable8)
	seqIt.Open(nil, nil)
	for seqIt.HasNext() {
		table8Tuple, err := seqIt.GetNext()
		if err != nil {
			return nil, err
		}

		//TODO mantain: if schema of table 8 is changed, this index could be asked to change
		tableIdBytes, err := table8Tuple.TupleGetFieldValue(0)
		if err != nil {
			return nil, err
		}
		tupleTableId, err := BytesToINT(tableIdBytes)
		if err != nil {
			return nil, err
		}
		schemaStringBytes, err := table8Tuple.TupleGetFieldValue(1)
		if err != nil {
			return nil, err
		}
		schemaString, err := BytesToVARCHAR(schemaStringBytes)
		if err != nil {
			return nil, err
		}

		//parse this string to get schema
		ast, _ := parser.GetParser().ParseSql(schemaString)
		currentSchema, err := GetRewriter().ASTNodeToSchema(ast)
		if err != nil {
			return nil, err
		}

		//found correct tuple
		if tupleTableId == int32(tableId) {
			return currentSchema, nil
		}

	}
	return nil, errors.New("execution/keyTableManager.go    GetTableSchema() no such table")
}

//if corresponding table is not in k_tableId_schema, a new tuple would be inserted
//if tableId <= 16, ignore
func (ktm *KeytableManager) Update_k_tableId_schema(tableId uint32, newSchemaString string) {
	ktm.Delete_k_tableId_schema(tableId)
	ktm.Insert_k_tableId_schema(tableId, newSchemaString)
}

//if corresponding table is not in k_tableId_schema, ignore
//if tableId <=16, ignore
func (ktm *KeytableManager) Delete_k_tableId_schema(tableId uint32) {
	se := storage.GetStorageEngine()
	transaction := storage.GetTransaction()

	//if tableId <= 16, ignore
	if tableId <= 16 {
		return
	}

	//get schema8
	schema8 := ktm.GetKeyTableSchema(8)

	//get headPage8
	headPage8, _ := se.GetDataPage(8, schema8)

	var targetPage *storage.DataPage
	var targetTuple *container.Tuple

	//loop until find targetTuple or reach tailPage
	currentPage := headPage8
	for {
		for i := 0; i < int(currentPage.DpGetTupleNum()); i++ {
			currentTuple, _ := currentPage.GetTupleAt(i)

			currentTableIdBytes, _ := currentTuple.TupleGetFieldValue(0)
			currentTableIdInt32, _ := BytesToINT(currentTableIdBytes)
			currentTableId := uint32(currentTableIdInt32)

			if currentTableId == tableId {
				targetPage = currentPage
				targetTuple = currentTuple
				break
			}
		}

		if targetPage != nil && targetTuple != nil {
			break
		}

		nextPageId, _ := currentPage.DpGetNextPageId()
		if nextPageId == currentPage.DpGetPageId() { //reach tailPage, end loop
			return
		}

		currentPage, _ = se.GetDataPage(nextPageId, schema8)
	}

	//get tailPageId, lastTupleId, tupleNum by tableId
	_, tailPageId, lastTupleId, tupleNum, _ := ktm.Query_k_table(8)

	//delete targetTuple from targetPage
	targetPage.DpDeleteTuple(targetTuple.TupleGetTupleId())
	transaction.InsertDataPage(targetPage)

	if targetPage.DpGetTupleNum() == 0 { //empty page, should be deleted

		priorPageId, _ := targetPage.DpGetPriorPageId()
		nextPageId, _ := targetPage.DpGetNextPageId()
		priorPage, _ := se.GetDataPage(priorPageId, schema8)
		nextPage, _ := se.GetDataPage(nextPageId, schema8)

		if targetPage.DpGetPageId() == tailPageId { //delete tail page

			targetPage.DpSetPriorPageId(targetPage.DpGetTableId())
			priorPage.DpSetNextPageId(priorPage.DpGetPageId())
			transaction.InsertDataPage(priorPage)

			tailPageId = priorPageId

		} else { //delete non tail page

			targetPage.DpSetPriorPageId(targetPage.DpGetPageId())
			targetPage.DpSetNextPageId(targetPage.DpGetPageId())

			priorPage.DpSetNextPageId(nextPageId)
			nextPage.DpSetPriorPageId(priorPageId)
			transaction.InsertDataPage(priorPage)
			transaction.InsertDataPage(nextPage)

		}

	}

	//update k_table
	if targetTuple.TupleGetTupleId() == lastTupleId {
		lastTupleId--
	}
	ktm.Update_k_table(8, tailPageId, lastTupleId, tupleNum-1)
}

//be careful not to insert an appeared tableId, this would not be checked
func (ktm *KeytableManager) Insert_k_tableId_schema(tableId uint32, newSchemaString string) {
	se := storage.GetStorageEngine()
	transaction := storage.GetTransaction()

	//get schema8
	schema8 := ktm.GetKeyTableSchema(8)

	//query tailPageId, lastTupleId, tupleNum from key table 9
	_, tailPageId, lastTupleId, tupleNum, _ := ktm.Query_k_table(8)

	field0Bytes := INTToBytes(int32(tableId))
	field1Bytes, _ := VARCHARToBytes(newSchemaString)
	field0, _ := container.NewFieldFromBytes(field0Bytes)
	field1, _ := container.NewFieldFromBytes(field1Bytes)
	var fields []*container.Field
	fields = append(fields, field0)
	fields = append(fields, field1)
	newTuple, _ := container.NewTuple(8, lastTupleId+1, schema8, fields)

	//insert this newTuple into table 8
	//check if newTuple can fit into the current last page
	currentLastPage, _ := se.GetDataPage(tailPageId, schema8)
	if currentLastPage.DpVacantByteNum() > newTuple.TupleSizeInBytes() { //can hold newTuple
		currentLastPage.InsertTuple(newTuple)
		transaction.InsertDataPage(currentLastPage)
	} else { //create a new page to hold it
		newTailPageId := ktm.GetVacantDataPageId()
		newTailPage := storage.NewDataPageMode0(newTailPageId, 8, currentLastPage.DpGetPageId(), newTailPageId)
		newTailPage.InsertTuple(newTuple)
		se.InsertDataPage(newTailPage)
		transaction.InsertDataPage(newTailPage)

		//update oldLastPage
		currentLastPage.DpSetNextPageId(newTailPageId)
		transaction.InsertDataPage(currentLastPage)

		tailPageId = newTailPageId
	}

	//update k_table
	ktm.Update_k_table(8, tailPageId, lastTupleId+1, tupleNum+1)
}

//get headPageId, tailPageId, lastTupleId, tupleNum by tableId
func (ktm *KeytableManager) Query_k_table(tableId uint32) (uint32, uint32, uint32, int32, error) {
	astOfCreateTable9, _ := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[9])
	schemaOfCreateTable9, err := GetRewriter().ASTNodeToSchema(astOfCreateTable9)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	seqIt := its.NewSequentialFileReaderIterator(9, schemaOfCreateTable9)
	seqIt.Open(nil, nil)
	for seqIt.HasNext() {
		table9Tuple, err := seqIt.GetNext()
		if err != nil {
			return 0, 0, 0, 0, err
		}

		//TODO mantain: if schema of table 9 is changed, this index could be asked to change
		tupleTableIdBytes, _ := table9Tuple.TupleGetFieldValue(0)
		tupleTableId32, _ := BytesToINT(tupleTableIdBytes)
		tupleTableId := uint32(tupleTableId32)

		headPageIdBytes, _ := table9Tuple.TupleGetFieldValue(1)
		headPageId32, _ := BytesToINT(headPageIdBytes)
		headPageId := uint32(headPageId32)

		tailPageIdBytes, _ := table9Tuple.TupleGetFieldValue(2)
		tailPageId32, _ := BytesToINT(tailPageIdBytes)
		tailPageId := uint32(tailPageId32)

		lastTupleIdBytes, _ := table9Tuple.TupleGetFieldValue(3)
		lastTupleId32, _ := BytesToINT(lastTupleIdBytes)
		lastTupleId := uint32(lastTupleId32)

		tupleNumBytes, _ := table9Tuple.TupleGetFieldValue(4)
		tupleNum, _ := BytesToINT(tupleNumBytes)

		//found correct tuple
		if tupleTableId == tableId {
			return headPageId, tailPageId, lastTupleId, tupleNum, nil
		}

	}
	return 0, 0, 0, 0, errors.New("execution/keyTableManager.go    Query_k_table_FromTableId() no such table")
}

//if no corresponding tuple, ignore
//if tableId <= 16, ignore
func (ktm *KeytableManager) Delete_k_table(tableId uint32) {
	//if tableId <= 16, ignore
	if tableId <= 16 {
		return
	}

	se := storage.GetStorageEngine()
	transaction := storage.GetTransaction()

	//get schema9
	schema9 := ktm.GetKeyTableSchema(9)

	//get headPage9
	headPage9, _ := se.GetDataPage(9, schema9)

	currentPage := headPage9
	var targetTuple *container.Tuple
	var targetPage *storage.DataPage
	var largestTupleId uint32 = 0
	var secondLargestTupleId uint32 = 0

	//loop until targetTuple is found or reach tailPage
	for {
		for i := 0; i < int(currentPage.DpGetTupleNum()); i++ {
			currentTuple, _ := currentPage.GetTupleAt(i)

			//update largestTupleId, secondLargestTupleId
			if currentTuple.TupleGetTupleId() > largestTupleId {
				secondLargestTupleId = largestTupleId
				largestTupleId = currentTuple.TupleGetTupleId()
			} else if currentTuple.TupleGetTupleId() > secondLargestTupleId {
				secondLargestTupleId = currentTuple.TupleGetTupleId()
			}

			//check if found targetTuple
			currentTableIdBytes, _ := currentTuple.TupleGetFieldValue(0)
			currentTableIdInt32, _ := BytesToINT(currentTableIdBytes)
			currentTableId := uint32(currentTableIdInt32)

			if currentTableId == tableId { //find targetTuple
				targetPage = currentPage
				targetTuple = currentTuple
			}
		}

		nextPageId, _ := currentPage.DpGetNextPageId()
		if nextPageId == currentPage.DpGetPageId() { //reach tailPage
			break
		}

		currentPage, _ = se.GetDataPage(nextPageId, schema9)
	}

	if targetTuple == nil || targetPage == nil { //targetTuple is not found
		return
	}

	targetPage.DpDeleteTuple(targetTuple.TupleGetTupleId())
	transaction.InsertDataPage(targetPage)

	//check if lastTupleId in k_table is tupleId of deleted tuple
	//if it is, set lastTupleId to secondLargestTupleId
	_, tailPageId, lastTupleId, tupleNum, _ := ktm.Query_k_table(9)
	if largestTupleId == targetTuple.TupleGetTupleId() {
		lastTupleId = secondLargestTupleId
	}

	//check if targetPage is empty
	if targetPage.DpGetTupleNum() == 0 {

		priorPageId, _ := targetPage.DpGetPriorPageId()
		nextPageId, _ := targetPage.DpGetNextPageId()
		priorPage, _ := se.GetDataPage(priorPageId, schema9)
		nextPage, _ := se.GetDataPage(nextPageId, schema9)

		//return pageId
		ktm.InsertVacantDataPageId(targetPage.DpGetPageId())

		if nextPageId == targetPage.DpGetPageId() { //targetPage is tailPage
			priorPage.DpSetNextPageId(priorPageId)
			transaction.InsertDataPage(priorPage)

			targetPage.DpSetPriorPageId(targetPage.DpGetPageId())

			tailPageId = priorPageId
		} else {
			priorPage.DpSetNextPageId(nextPageId)
			nextPage.DpSetPriorPageId(priorPageId)
			transaction.InsertDataPage(priorPage)
			transaction.InsertDataPage(nextPage)

			targetPage.DpSetPriorPageId(targetPage.DpGetPageId())
			targetPage.DpSetNextPageId(targetPage.DpGetPageId())
		}
	}

	//update k_table
	ktm.Update_k_table(9, tailPageId, lastTupleId, tupleNum-1)
}

//if tableId has appeared, ignore
func (ktm *KeytableManager) Insert_k_table(tableId uint32, headPageId uint32, tailPageId uint32, lastTupleId uint32, tupleNum int32) {
	se := storage.GetStorageEngine()
	transaction := storage.GetTransaction()

	//if tableId has appeared, ignore
	_, _, _, _, err := ktm.Query_k_table(tableId)
	if err == nil {
		return
	}

	//get schema9
	schema9 := ktm.GetKeyTableSchema(9)

	//get k_table tuple info
	_, tailPageId9, lastTupleId9, tupleNum9, _ := ktm.Query_k_table(9)
	oldTailPage9, _ := se.GetDataPage(tailPageId9, schema9)

	//create newTuple
	newTupleId := lastTupleId9 + 1

	//create fields
	field_tableId, _ := container.NewFieldFromBytes(INTToBytes(int32(tableId)))
	field_headPageId, _ := container.NewFieldFromBytes(INTToBytes(int32(headPageId)))
	field_tailPageId, _ := container.NewFieldFromBytes(INTToBytes(int32(tailPageId)))
	field_lastTupleId, _ := container.NewFieldFromBytes(INTToBytes(int32(lastTupleId)))
	field_tupleNum, _ := container.NewFieldFromBytes(INTToBytes(tupleNum))
	var fields []*container.Field
	fields = append(fields, field_tableId)
	fields = append(fields, field_headPageId)
	fields = append(fields, field_tailPageId)
	fields = append(fields, field_lastTupleId)
	fields = append(fields, field_tupleNum)

	//create newTuple
	newTuple, _ := container.NewTuple(9, newTupleId, schema9, fields)

	//check if oldTailPage9 is full
	if oldTailPage9.DpVacantByteNum() > newTuple.TupleSizeInBytes() { //oldTailPage9 can hold this new tuple
		//update k_table
		ktm.Update_k_table(9, tailPageId9, newTupleId, tupleNum9+1)

		oldTailPage9.InsertTuple(newTuple)
		transaction.InsertDataPage(oldTailPage9)
	} else { //create a new tail page to hold the tuple
		//get a vacant pageId
		vacantPageId := ktm.GetVacantDataPageId()
		newTailPage9 := storage.NewDataPageMode0(vacantPageId, 9, oldTailPage9.DpGetPageId(), vacantPageId)

		//update k_table
		ktm.Update_k_table(9, vacantPageId, newTupleId, tupleNum9+1)

		oldTailPage9.DpSetNextPageId(vacantPageId)
		transaction.InsertDataPage(oldTailPage9)

		newTailPage9.InsertTuple(newTuple)
		se.InsertDataPage(newTailPage9)
		transaction.InsertDataPage(newTailPage9)
	}
}

//update a tuple in k_table
//if no such old tuple, just ignore it
//TODO update: use index to accelerate
func (ktm *KeytableManager) Update_k_table(tableId uint32, tailPageId uint32, lastTupleId uint32, tupleNum int32) {
	se := storage.GetStorageEngine()
	transaction := storage.GetTransaction()

	//get headPage of k_table
	schema9 := ktm.GetKeyTableSchema(9)
	headPage9, _ := se.GetDataPage(9, schema9)

	currentPage := headPage9
	var targetTuple *container.Tuple
	var targetPage *storage.DataPage

	//loop until targetTuple is found or reach tailPage
	for {
		for i := 0; i < int(currentPage.DpGetTupleNum()); i++ {
			currentTuple, _ := currentPage.GetTupleAt(i)

			//check if found targetTuple
			currentTableIdBytes, _ := currentTuple.TupleGetFieldValue(0)
			currentTableIdInt32, _ := BytesToINT(currentTableIdBytes)
			currentTableId := uint32(currentTableIdInt32)

			if currentTableId == tableId { //find targetTuple
				targetPage = currentPage
				targetTuple = currentTuple
				break
			}
		}

		nextPageId, _ := currentPage.DpGetNextPageId()
		if nextPageId == currentPage.DpGetPageId() { //reach tailPage
			break
		}

		currentPage, _ = se.GetDataPage(nextPageId, schema9)
	}

	if targetTuple == nil || targetPage == nil { //targetTuple is not found
		return
	}

	//if targetTuple is not nil (related tuple already in k_table), delete it
	//if no such target tuple, just ignore
	if targetTuple != nil {
		targetPage.DpDeleteTuple(targetTuple.TupleGetTupleId())
	} else {
		return
	}

	//get old headPageId
	oldHeadPageIdBytes, _ := targetTuple.TupleGetFieldValue(1)
	oldHeadPageIdInt32, _ := BytesToINT(oldHeadPageIdBytes)

	//insert tuple
	//tableId INT, headPageId INT, tailPageId INT, lastTupleId INT, tupleNum INT
	field_tableId, _ := container.NewFieldFromBytes(INTToBytes(int32(tableId)))
	field_headPageId, _ := container.NewFieldFromBytes(INTToBytes(oldHeadPageIdInt32))
	field_tailPageId, _ := container.NewFieldFromBytes(INTToBytes(int32(tailPageId)))
	field_lastTupleId, _ := container.NewFieldFromBytes(INTToBytes(int32(lastTupleId)))
	field_tupleNum, _ := container.NewFieldFromBytes(INTToBytes(tupleNum))
	var fields []*container.Field
	fields = append(fields, field_tableId)
	fields = append(fields, field_headPageId)
	fields = append(fields, field_tailPageId)
	fields = append(fields, field_lastTupleId)
	fields = append(fields, field_tupleNum)
	newTuple, _ := container.NewTuple(9, tableId, schema9, fields)
	targetPage.InsertTuple(newTuple)

	//insert targetPage into transaction
	transaction.InsertDataPage(targetPage)
}

//get a vacant dataPageId and delete it from k_emptyDataPageSlot
//in k_emptyDataPageSlot (key table 15), tuple 0 is the id after&include which all dataPages are vacant
//tuple 0 should never be deleted
//if there are more than one tuple in k_emptyDataPageSlot, find the end one, return and delete it
func (ktm *KeytableManager) GetVacantDataPageId() uint32 {
	se := storage.GetStorageEngine()
	transaction := storage.GetTransaction()

	//get schema of key table 15
	schema15 := ktm.GetKeyTableSchema(15)

	//get headPageId, tailPageId, lastTupleId, tupleNum
	headPageId, tailPageId, lastTupleId, tupleNum, _ := ktm.Query_k_table(15)

	//get headPage & tailPage
	headPage, _ := se.GetDataPage(headPageId, schema15)
	tailPage, _ := se.GetDataPage(tailPageId, schema15)

	var returnValue uint32

	//check if there is only tuple0
	if tupleNum == 1 {
		//get tuple0 and returnValue
		tuple0, _ := headPage.GetTuple(0)
		bytes, _ := tuple0.TupleGetFieldValue(0)
		returnValueInt32, _ := BytesToINT(bytes)
		returnValue = uint32(returnValueInt32)

		//delete tuple0
		headPage.DpDeleteTuple(0)

		//create a new tuple0 and insert it into headPage
		newField, _ := container.NewFieldFromBytes(INTToBytes(returnValueInt32 + 1))
		var fields []*container.Field
		fields = append(fields, newField)
		newTuple0, _ := container.NewTuple(15, 0, schema15, fields)
		headPage.InsertTuple(newTuple0)
		transaction.InsertDataPage(headPage)

	} else { //delete the last tuple, if tail page is empty, delete the page as well

		//get lastTuple & returnValue
		lastTuple, _ := tailPage.GetTuple(lastTupleId)
		bytes, _ := lastTuple.TupleGetFieldValue(0)
		returnValueInt32, _ := BytesToINT(bytes)
		returnValue = uint32(returnValueInt32)

		//delete lastTuple
		tailPage.DpDeleteTuple(lastTupleId)
		transaction.InsertDataPage(tailPage)

		//check if tailPage is empty
		if tailPage.DpGetTupleNum() == 0 { //this is already an empty page, delete it
			//get newTailPageId
			newTailPageId, _ := tailPage.DpGetPriorPageId()

			//set old tailPage to isolate
			tailPage.DpSetPriorPageId(tailPageId)

			//recycle pageId of old tailPage
			ktm.InsertVacantDataPageId(tailPageId)

			//update nexPageId of newTailPage
			newTailPage, _ := se.GetDataPage(newTailPageId, schema15)
			newTailPage.DpSetNextPageId(newTailPageId)
			transaction.InsertDataPage(newTailPage)

			//update k_table
			ktm.Update_k_table(15, newTailPageId, lastTupleId-1, tupleNum-1)
		} else {
			//update k_table
			ktm.Update_k_table(15, tailPageId, lastTupleId-1, tupleNum-1)
		}
	}
	return returnValue
}

//get a vacant indexPageId and delete it from k_emptyIndexPageSlot
//in k_emptyIndexPageSlot (key table 16), tuple 0 is the id after&include which all indexPages are vacant
//tuple 0 should never be deleted
//if there are more than one tuple in k_emptyIndexPageSlot, find the end one, return and delete it
func (ktm *KeytableManager) GetVacantIndexPageId() uint32 {
	se := storage.GetStorageEngine()
	transaction := storage.GetTransaction()

	//get schema of key table 16
	schema16 := ktm.GetKeyTableSchema(16)

	//get headPageId, tailPageId, lastTupleId, tupleNum
	headPageId, tailPageId, lastTupleId, tupleNum, _ := ktm.Query_k_table(16)

	//get headPage & tailPage
	headPage, _ := se.GetDataPage(headPageId, schema16)
	tailPage, _ := se.GetDataPage(tailPageId, schema16)

	var returnValue uint32

	//check if there is only tuple0
	if tupleNum == 1 {
		//get tuple0 and returnValue
		tuple0, _ := headPage.GetTuple(0)
		bytes, _ := tuple0.TupleGetFieldValue(0)
		returnValueInt32, _ := BytesToINT(bytes)
		returnValue = uint32(returnValueInt32)

		//delete tuple0
		headPage.DpDeleteTuple(0)

		//create a new tuple0 and insert it into headPage
		newField, _ := container.NewFieldFromBytes(INTToBytes(returnValueInt32 + 1))
		var fields []*container.Field
		fields = append(fields, newField)
		newTuple0, _ := container.NewTuple(16, 0, schema16, fields)
		headPage.InsertTuple(newTuple0)
		transaction.InsertDataPage(headPage)

	} else { //delete the last tuple, if tail page is empty, delete the page as well

		//get lastTuple & returnValue
		lastTuple, _ := tailPage.GetTuple(lastTupleId)
		bytes, _ := lastTuple.TupleGetFieldValue(0)
		returnValueInt32, _ := BytesToINT(bytes)
		returnValue = uint32(returnValueInt32)

		//delete lastTuple
		tailPage.DpDeleteTuple(lastTupleId)
		transaction.InsertDataPage(tailPage)

		//check if tailPage is empty
		if tailPage.DpGetTupleNum() == 0 { //this is already an empty page, delete it
			//get newTailPageId
			newTailPageId, _ := tailPage.DpGetPriorPageId()

			//set old tailPage to isolate
			tailPage.DpSetPriorPageId(tailPageId)

			//recycle pageId of old tailPage
			ktm.InsertVacantDataPageId(tailPageId)

			//update nexPageId of newTailPage
			newTailPage, _ := se.GetDataPage(newTailPageId, schema16)
			newTailPage.DpSetNextPageId(newTailPageId)
			transaction.InsertDataPage(newTailPage)

			//update k_table
			ktm.Update_k_table(16, newTailPageId, lastTupleId-1, tupleNum-1)
		} else {
			//update k_table
			ktm.Update_k_table(16, tailPageId, lastTupleId-1, tupleNum-1)
		}
	}
	return returnValue
}

//insert pageId at the end of k_emptyDataPageSlot
func (ktm *KeytableManager) InsertVacantDataPageId(pageId uint32) {
	se := storage.GetStorageEngine()
	transaction := storage.GetTransaction()

	//get schema of k_emptyDataPageSlot
	schema15 := ktm.GetKeyTableSchema(15)

	//get tailPageId,lastTupleId,tupleNum of k_emptyDataPageSlot
	_, tailPageId, lastTupleId, tupleNum, _ := ktm.Query_k_table(15)

	//create newTuple for pageId
	field, _ := container.NewFieldFromBytes(INTToBytes(int32(pageId)))
	var fields []*container.Field
	fields = append(fields, field)
	newTuple, _ := container.NewTuple(15, lastTupleId+1, schema15, fields)

	//get tailPage of k_emptyDataPageSlot
	tailPage15, _ := se.GetDataPage(tailPageId, schema15)

	//check if tailPage can not hold new tuple
	//TODO maintain: index could change is schema of k_emptyDataPageSlot changes
	if tailPage15.DpVacantByteNum() >= newTuple.TupleSizeInBytes() { // enough space

		//insert pageId into k_emptyDataPageSlot
		tailPage15.InsertTuple(newTuple)
		transaction.InsertDataPage(tailPage15)

		//update k_table, tuple for k_emptyDataPageSlot
		ktm.Update_k_table(15, tailPageId, lastTupleId+1, tupleNum+1)

	} else { //no enough space, should create a new page
		newTailPageId := ktm.GetVacantDataPageId()
		newTailPage := storage.NewDataPageMode0(newTailPageId, 15, tailPageId, newTailPageId)

		//insert newTailPage into dataBuffer
		se.InsertDataPage(newTailPage)

		//set old tailPage nextPageId
		tailPage15.DpSetNextPageId(newTailPageId)
		transaction.InsertDataPage(tailPage15)

		//insert pageId into k_emptyDataPageSlot
		newTailPage.InsertTuple(newTuple)
		transaction.InsertDataPage(newTailPage)

		//update k_table, tuple for k_emptyDataPageSlot
		ktm.Update_k_table(15, newTailPageId, lastTupleId+1, tupleNum+1)
	}
}

//insert pageId at the end of k_emptyIndexPageSlot
func (ktm *KeytableManager) InsertVacantIndexPageId(pageId uint32) {
	se := storage.GetStorageEngine()
	transaction := storage.GetTransaction()

	//get schema of k_emptyIndexPageSlot
	schema16 := ktm.GetKeyTableSchema(16)

	//get tailPageId,lastTupleId,tupleNum of k_emptyIndexPageSlot
	_, tailPageId, lastTupleId, tupleNum, _ := ktm.Query_k_table(16)

	//create newTuple for pageId
	field, _ := container.NewFieldFromBytes(INTToBytes(int32(pageId)))
	var fields []*container.Field
	fields = append(fields, field)
	newTuple, _ := container.NewTuple(16, lastTupleId+1, schema16, fields)

	//get tailPage of k_emptyIndexPageSlot
	tailPage16, _ := se.GetDataPage(tailPageId, schema16)

	//check if tailPage can not hold new tuple
	//TODO maintain: index could change is schema of k_emptyIndexPageSlot changes
	if tailPage16.DpVacantByteNum() >= newTuple.TupleSizeInBytes() { // enough space

		//insert pageId into k_emptyIndexPageSlot
		tailPage16.InsertTuple(newTuple)
		transaction.InsertDataPage(tailPage16)

		//update k_table, tuple for k_emptyIndexPageSlot
		ktm.Update_k_table(16, tailPageId, lastTupleId+1, tupleNum+1)

	} else { //no enough space, should create a new page
		newTailPageId := ktm.GetVacantDataPageId()
		newTailPage := storage.NewDataPageMode0(newTailPageId, 16, tailPageId, newTailPageId)

		//insert newTailPage into dataBuffer
		se.InsertDataPage(newTailPage)

		//set old tailPage nextPageId
		tailPage16.DpSetNextPageId(newTailPageId)
		transaction.InsertDataPage(tailPage16)

		//insert pageId into k_emptyIndexPageSlot
		newTailPage.InsertTuple(newTuple)
		transaction.InsertDataPage(newTailPage)

		//update k_table, tuple for k_emptyIndexPageSlot
		ktm.Update_k_table(16, newTailPageId, lastTupleId+1, tupleNum+1)
	}
}

//initialize the whole system
func (ktm *KeytableManager) InitializeSystem() {
	//erase dataPage, indexPage and logPage
	storage.GetStorageEngine().EraseDataFile()
	storage.GetStorageEngine().EraseIndexFile()
	storage.GetStorageEngine().EraseLogFile()

	//create p0, p1, p2, p8, p9, p15, p16
	p0 := storage.NewDataPageMode0(0, 0, 0, 0)
	p1 := storage.NewDataPageMode0(1, 1, 1, 1)
	p2 := storage.NewDataPageMode0(2, 2, 2, 2)
	p3 := storage.NewDataPageMode0(3, 3, 3, 3)
	p4 := storage.NewDataPageMode0(4, 4, 4, 4)
	p5 := storage.NewDataPageMode0(5, 5, 5, 5)
	p6 := storage.NewDataPageMode0(6, 6, 6, 6)
	p7 := storage.NewDataPageMode0(7, 7, 7, 7)
	p8 := storage.NewDataPageMode0(8, 8, 8, 8)
	p9 := storage.NewDataPageMode0(9, 9, 9, 9)
	p10 := storage.NewDataPageMode0(10, 10, 10, 10)
	p11 := storage.NewDataPageMode0(11, 11, 11, 11)
	p12 := storage.NewDataPageMode0(12, 12, 12, 12)
	p13 := storage.NewDataPageMode0(13, 13, 13, 13)
	p14 := storage.NewDataPageMode0(14, 14, 14, 14)
	p15 := storage.NewDataPageMode0(15, 15, 15, 15)
	p16 := storage.NewDataPageMode0(16, 16, 16, 16)

	//insert tuple into page0
	//DEFAULT_ADMINISTRATOR_USER_ID, DEFAULT_ADMINISTRATOR_NAME
	table0CreateAst, _ := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[0])
	table0Schema, _ := GetRewriter().ASTNodeToSchema(table0CreateAst)

	field000, _ := container.NewFieldFromBytes(INTToBytes(int32(DEFAULT_ADMINISTRATOR_USER_ID)))
	userNameBytes, _ := VARCHARToBytes(DEFAULT_ADMINISTRATOR_NAME)
	field001, _ := container.NewFieldFromBytes(userNameBytes)

	var fields00 []*container.Field
	fields00 = append(fields00, field000)
	fields00 = append(fields00, field001)

	tuple00, _ := container.NewTuple(0, 0, table0Schema, fields00)
	p0.InsertTuple(tuple00)

	//insert tuple into page1
	//DEFAULT_ADMINISTRATOR_USER_ID, DEFAULT_ADMINISTRATOR_PASSWORD
	table1CreateAst, _ := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[1])
	table1Schema, _ := GetRewriter().ASTNodeToSchema(table1CreateAst)

	field100, _ := container.NewFieldFromBytes(INTToBytes(int32(DEFAULT_ADMINISTRATOR_USER_ID)))
	userPasswordBytes, _ := VARCHARToBytes(DEFAULT_ADMINISTRATOR_PASSWORD)
	field101, _ := container.NewFieldFromBytes(userPasswordBytes)

	var fields10 []*container.Field
	fields10 = append(fields10, field100)
	fields10 = append(fields10, field101)

	tuple10, _ := container.NewTuple(1, 0, table1Schema, fields10)
	p1.InsertTuple(tuple10)

	//insert tuple into page2
	//all key tabls belong to administrator
	//0 to 16, DEFAULT_ADMINISTRATOR_USER_ID
	table2CreateAst, _ := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[2])
	table2Schema, _ := GetRewriter().ASTNodeToSchema(table2CreateAst)

	f2x1, _ := container.NewFieldFromBytes(INTToBytes(int32(DEFAULT_ADMINISTRATOR_USER_ID)))

	for i := 0; i <= 16; i++ {
		f2x0, _ := container.NewFieldFromBytes(INTToBytes(int32(i)))
		var fields []*container.Field
		fields = append(fields, f2x0)
		fields = append(fields, f2x1)
		newTuple, _ := container.NewTuple(2, uint32(i), table2Schema, fields)
		p2.InsertTuple(newTuple)
	}

	//insert tuple into page8
	//0 to 16, DEFAULT_KEY_TABLE_0_SCHEMA to DEFAULT_KEY_TABLE_16_SCHEMA
	table8CreateAst, _ := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[8])
	table8Schema, _ := GetRewriter().ASTNodeToSchema(table8CreateAst)

	for i := 0; i <= 16; i++ {
		f8x0, _ := container.NewFieldFromBytes(INTToBytes(int32(i)))
		f8x1Bytes, _ := VARCHARToBytes(DEFAULT_KEYTABLES_SCHEMA[i])
		f8x1, _ := container.NewFieldFromBytes(f8x1Bytes)

		var fields8 []*container.Field
		fields8 = append(fields8, f8x0)
		fields8 = append(fields8, f8x1)

		tuple8, _ := container.NewTuple(8, uint32(i), table8Schema, fields8)
		p8.InsertTuple(tuple8)
	}

	//insert tuple into page9
	//0 to 16, 0 to 16, depend on tableId, depend on tableId
	//lastTupleId is invalid when tupleNum = 0
	table9CreateAst, _ := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[9])
	table9Schema, _ := GetRewriter().ASTNodeToSchema(table9CreateAst)

	for i := 0; i <= 16; i++ {

		f9x0, _ := container.NewFieldFromBytes(INTToBytes(int32(i)))
		f9x1, _ := container.NewFieldFromBytes(INTToBytes(int32(i)))
		f9x2, _ := container.NewFieldFromBytes(INTToBytes(int32(i)))
		var f9x3 *container.Field
		var f9x4 *container.Field
		var fields9 []*container.Field

		if i == 0 || i == 1 || i == 15 || i == 16 { //these tables have one tuple
			//k_userId_userName, k_userId_password, k_emptyDataPageSlot, k_emptyIndexPageSlot
			f9x3, _ = container.NewFieldFromBytes(INTToBytes(int32(0)))
			f9x4, _ = container.NewFieldFromBytes(INTToBytes(int32(1)))
		} else if i == 2 || i == 8 || i == 9 { //these tables have 17 tuples
			//k_tableId_userId, k_tableId_schema, k_table
			f9x3, _ = container.NewFieldFromBytes(INTToBytes(int32(16)))
			f9x4, _ = container.NewFieldFromBytes(INTToBytes(int32(17)))
		} else {
			//all other key tables
			f9x3, _ = container.NewFieldFromBytes(INTToBytes(int32(0)))
			f9x4, _ = container.NewFieldFromBytes(INTToBytes(int32(0)))
		}

		fields9 = append(fields9, f9x0)
		fields9 = append(fields9, f9x1)
		fields9 = append(fields9, f9x2)
		fields9 = append(fields9, f9x3)
		fields9 = append(fields9, f9x4)

		tuple9, _ := container.NewTuple(9, uint32(i), table9Schema, fields9)
		p9.InsertTuple(tuple9)
	}

	//insert tuple into page15
	//tuple0 is important: it keeps a pageId, for all dataPages whose pageId >= this value, they are not allocated
	//the first 17 pages have been occupied by key tables
	//17
	table15CreateAst, _ := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[15])
	table15Schema, _ := GetRewriter().ASTNodeToSchema(table15CreateAst)

	field1500, _ := container.NewFieldFromBytes(INTToBytes(17))
	var fields15 []*container.Field
	fields15 = append(fields15, field1500)
	tuple150, _ := container.NewTuple(15, 0, table15Schema, fields15)
	p15.InsertTuple(tuple150)

	//insert tuple into page16
	//tuple0 is important: it keeps a pageId, for all indexPages whose pageId >= this value, they are not allocated
	//0
	table16CreateAst, _ := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[16])
	table16Schema, _ := GetRewriter().ASTNodeToSchema(table16CreateAst)

	field1600, _ := container.NewFieldFromBytes(INTToBytes(0))
	var fields16 []*container.Field
	fields16 = append(fields16, field1600)
	tuple160, _ := container.NewTuple(16, 0, table16Schema, fields16)
	p16.InsertTuple(tuple160)

	//insert these pages into storageEngine
	se := storage.GetStorageEngine()
	se.InsertDataPage(p0)
	se.InsertDataPage(p1)
	se.InsertDataPage(p2)
	se.InsertDataPage(p3)
	se.InsertDataPage(p4)
	se.InsertDataPage(p5)
	se.InsertDataPage(p6)
	se.InsertDataPage(p7)
	se.InsertDataPage(p8)
	se.InsertDataPage(p9)
	se.InsertDataPage(p10)
	se.InsertDataPage(p11)
	se.InsertDataPage(p12)
	se.InsertDataPage(p13)
	se.InsertDataPage(p14)
	se.InsertDataPage(p15)
	se.InsertDataPage(p16)

	//push pages into transaction, conduct transaction
	transaction := storage.GetTransaction()
	transaction.InsertDataPage(p0)
	transaction.InsertDataPage(p1)
	transaction.InsertDataPage(p2)
	transaction.InsertDataPage(p3)
	transaction.InsertDataPage(p4)
	transaction.InsertDataPage(p5)
	transaction.InsertDataPage(p6)
	transaction.InsertDataPage(p7)
	transaction.InsertDataPage(p8)
	transaction.InsertDataPage(p9)
	transaction.InsertDataPage(p10)
	transaction.InsertDataPage(p11)
	transaction.InsertDataPage(p12)
	transaction.InsertDataPage(p13)
	transaction.InsertDataPage(p14)
	transaction.InsertDataPage(p15)
	transaction.InsertDataPage(p16)
	transaction.PushTransactionIntoDisk()
}
