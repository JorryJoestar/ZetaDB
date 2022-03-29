package execution

import (
	"ZetaDB/container"
	its "ZetaDB/execution/querySubOperator"
	"ZetaDB/optimizer"
	"ZetaDB/parser"
	"ZetaDB/storage"
	. "ZetaDB/utility"
	"errors"
	"sync"
)

type KeytableManager struct {
}

//for singleton pattern
var instance *KeytableManager
var once sync.Once

//to get KeytableManager, call this function
func GetKeytableManager() *KeytableManager {
	once.Do(func() {
		instance = &KeytableManager{}
	})

	return instance
}

//get table info (tableId & schema) from dataFile according to tableName, k_tableId_schema table 8
//throw error if no such table
func (ktm *KeytableManager) Query_k_tableId_schema_FromTableName(tableName string) (uint32, *container.Schema, error) {
	astOfCreateTable8 := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[8])
	schemaOfCreateTable8, err := optimizer.GetRewriter().ASTNodeToSchema(astOfCreateTable8)
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
		ast := parser.GetParser().ParseSql(schemaString)
		currentSchema, err := optimizer.GetRewriter().ASTNodeToSchema(ast)
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

	astOfCreateTable8 := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[8])
	schemaOfCreateTable8, err := optimizer.GetRewriter().ASTNodeToSchema(astOfCreateTable8)
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
		ast := parser.GetParser().ParseSql(schemaString)
		currentSchema, err := optimizer.GetRewriter().ASTNodeToSchema(ast)
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

//get headPageId, lastTupleId, tupleNum by tableId
func (ktm *KeytableManager) Query_k_table_FromTableId(tableId uint32) (uint32, uint32, int32, error) {
	astOfCreateTable9 := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[9])
	schemaOfCreateTable9, err := optimizer.GetRewriter().ASTNodeToSchema(astOfCreateTable9)
	if err != nil {
		return 0, 0, 0, err
	}

	seqIt := its.NewSequentialFileReaderIterator(9, schemaOfCreateTable9)
	seqIt.Open(nil, nil)
	for seqIt.HasNext() {
		table9Tuple, err := seqIt.GetNext()
		if err != nil {
			return 0, 0, 0, err
		}

		//TODO mantain: if schema of table 9 is changed, this index could be asked to change
		tupleTableIdBytes, _ := table9Tuple.TupleGetFieldValue(0)
		tupleTableId32, _ := BytesToINT(tupleTableIdBytes)
		tupleTableId := uint32(tupleTableId32)

		headPageIdBytes, _ := table9Tuple.TupleGetFieldValue(1)
		headPageId32, _ := BytesToINT(headPageIdBytes)
		headPageId := uint32(headPageId32)

		lastTupleIdBytes, _ := table9Tuple.TupleGetFieldValue(2)
		lastTupleId32, _ := BytesToINT(lastTupleIdBytes)
		lastTupleId := uint32(lastTupleId32)

		tupleNumBytes, _ := table9Tuple.TupleGetFieldValue(3)
		tupleNum, _ := BytesToINT(tupleNumBytes)

		//found correct tuple
		if tupleTableId == tableId {
			return headPageId, lastTupleId, tupleNum, nil
		}

	}
	return 0, 0, 0, errors.New("execution/keyTableManager.go    Query_k_table_FromTableId() no such table")
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
	p8 := storage.NewDataPageMode0(8, 8, 8, 8)
	p9 := storage.NewDataPageMode0(9, 9, 9, 9)
	p15 := storage.NewDataPageMode0(15, 15, 15, 15)
	p16 := storage.NewDataPageMode0(16, 16, 16, 16)

	//insert tuple into page0
	//DEFAULT_ADMINISTRATOR_USER_ID, DEFAULT_ADMINISTRATOR_NAME
	table0CreateAst := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[0])
	table0Schema, _ := optimizer.GetRewriter().ASTNodeToSchema(table0CreateAst)

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
	table1CreateAst := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[1])
	table1Schema, _ := optimizer.GetRewriter().ASTNodeToSchema(table1CreateAst)

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
	table2CreateAst := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[2])
	table2Schema, _ := optimizer.GetRewriter().ASTNodeToSchema(table2CreateAst)

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
	table8CreateAst := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[8])
	table8Schema, _ := optimizer.GetRewriter().ASTNodeToSchema(table8CreateAst)

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
	table9CreateAst := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[9])
	table9Schema, _ := optimizer.GetRewriter().ASTNodeToSchema(table9CreateAst)

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
	table15CreateAst := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[15])
	table15Schema, _ := optimizer.GetRewriter().ASTNodeToSchema(table15CreateAst)

	field1500, _ := container.NewFieldFromBytes(INTToBytes(17))
	var fields15 []*container.Field
	fields15 = append(fields15, field1500)
	tuple150, _ := container.NewTuple(15, 0, table15Schema, fields15)
	p15.InsertTuple(tuple150)

	//insert tuple into page16
	//tuple0 is important: it keeps a pageId, for all indexPages whose pageId >= this value, they are not allocated
	//0
	table16CreateAst := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[16])
	table16Schema, _ := optimizer.GetRewriter().ASTNodeToSchema(table16CreateAst)

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
	se.InsertDataPage(p8)
	se.InsertDataPage(p9)
	se.InsertDataPage(p15)
	se.InsertDataPage(p16)

	//swap these pages into disk
	se.SwapDataPage(0)
	se.SwapDataPage(1)
	se.SwapDataPage(2)
	se.SwapDataPage(8)
	se.SwapDataPage(9)
	se.SwapDataPage(15)
	se.SwapDataPage(16)
}

//get a vacant dataPageId
//in k_emptyDataPageSlot (key table 15), tuple 0 is the id after&include which all dataPages are vacant
//tuple 0 should never be deleted
//if there are more than one tuple in k_emptyDataPageSlot, find the end one, return and delete it
func (ktm *KeytableManager) GetVacantDataPageId() uint32 {
	//get schema of table 15
	createTable15Ast := parser.GetParser().ParseSql(DEFAULT_KEYTABLES_SCHEMA[15])
	table15Schema, _ := optimizer.GetRewriter().ASTNodeToSchema(createTable15Ast)

	//prepare SequentialFileReaderIterator
	seqIt := its.NewSequentialFileReaderIterator(15, table15Schema)
	seqIt.Open(nil, nil)

	//get tuple0 & tupleEnd
	tuple0, _ := seqIt.GetNext()
	var tupleEnd *container.Tuple
	for seqIt.HasNext() {
		tupleEnd, _ = seqIt.GetNext()
	}

	if tupleEnd == nil { //only one tuple in table 15, return tuple0 and increase it
		//TODO mantain: if schema of k_emptyDataPageSlot changed, the index could need to be changed
		tableIdBytes, _ := tuple0.TupleGetFieldValue(0)
		tableIdInt32, _ := BytesToINT(tableIdBytes)
		tableId := uint32(tableIdInt32)

		//delete tuple0, insert a new one whose value is tuple0 + 1
		//TODO

		return tableId
	} else { //return and delete the end tuple

		//TODO
		return 0
	}
}

//get a vacant indexPageId
//in k_emptyIndexPageSlot (key table 16), tuple 0 is the id after&include which all indexPages are vacant
//tuple 0 should never be deleted
//if there are more than one tuple in k_emptyIndexPageSlot, find the end one, return and delete it
func (ktm *KeytableManager) GetVacantIndexPageId() uint32 {
	return 0
}
