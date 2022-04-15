package execution

import (
	"ZetaDB/container"
	"ZetaDB/parser"
	subOperator "ZetaDB/physicalPlan"
	"ZetaDB/storage"
	"ZetaDB/utility"
	"strconv"
	"sync"
)

type ExecutionEngine struct {
	se       *storage.StorageEngine
	parser   *parser.Parser
	rewriter *Rewriter
	ktm      *KeytableManager
	tm       *TableManipulator
}

//use GetExecutionEngine to get the unique ExecutionEngine
var eeInstance *ExecutionEngine
var eeOnce sync.Once

func GetExecutionEngine() *ExecutionEngine {

	eeOnce.Do(func() {
		eeInstance = &ExecutionEngine{
			se:       storage.GetStorageEngine(),
			parser:   parser.GetParser(),
			rewriter: &Rewriter{}}
	})
	eeInstance.tm = GetTableManipulator()
	eeInstance.ktm = GetKeytableManager()

	return eeInstance
}

func (ee *ExecutionEngine) Execute(executionPlan *container.ExecutionPlan) string {
	switch executionPlan.PlanType {
	case container.EP_INITIALIZE_SYSTEM:
	case container.EP_INSERT:
		tableName := executionPlan.Parameter[0]
		fieldStrings := executionPlan.Parameter[1:]
		ee.InsertOperator(tableName, fieldStrings)
	case container.EP_DELETE: //DELETE
		rw := GetRewriter()
		ktm := GetExecutionEngine().ktm
		tableName := executionPlan.Parameter[0]
		tableId, _, _ := ktm.Query_k_tableId_schema_FromTableName(tableName)
		physicalPlan := rw.LogicalPLanToPhysicalPlan(executionPlan.LogicalPlanRoot)
		var tuplesToDelete []*container.Tuple
		for physicalPlan.HasNext() {
			fetchedTuple, _ := physicalPlan.GetNext()
			tuplesToDelete = append(tuplesToDelete, fetchedTuple)
		}

		tm := GetTableManipulator()
		for _, tupleToDelete := range tuplesToDelete {
			tm.DeleteTupleFromTable(tableId, tupleToDelete.TupleGetTupleId())
		}
	case container.EP_UPDATE:
		//tableName := executionPlan.Parameter[0]

		rw := GetRewriter()
		physicalPlan := rw.LogicalPLanToPhysicalPlan(executionPlan.LogicalPlanRoot)
		var tuplesToDelete []*container.Tuple
		for physicalPlan.HasNext() {
			fetchedTuple, _ := physicalPlan.GetNext()
			tuplesToDelete = append(tuplesToDelete, fetchedTuple)
		}

		/* 		for _, oldTuple := range tuplesToDelete {
			oldTuple.get
		} */

		//fieldStrings := executionPlan.Parameter[1:]
		//ee.InsertOperator(tableName, fieldStrings)

	case container.EP_QUERY:
	case container.EP_CREATE_TABLE:
		userIdINT, _ := strconv.Atoi(executionPlan.Parameter[0])
		userId := int32(userIdINT)
		schemaString := executionPlan.Parameter[1]
		ee.CreateTableOperator(userId, schemaString)
	case container.EP_DROP_TABLE:
		tableName := executionPlan.Parameter[0]
		ee.DropTableOperator(tableName)
	case container.EP_ALTER_TABLE_ADD:
	case container.EP_ALTER_TABLE_DROP:
	case container.EP_CREATE_ASSERT:
	case container.EP_DROP_ASSERT:
	case container.EP_CREATE_VIEW:
	case container.EP_DROP_VIEW:
	case container.EP_CREATE_INDEX:
	case container.EP_DROP_INDEX:
	case container.EP_CREATE_TRIGGER:
	case container.EP_DROP_TRIGGER:
	case container.EP_CREATE_PSM:
	case container.EP_DROP_PSM:
	case container.EP_SHOW_TABLES:
	case container.EP_SHOW_ASSERTIONS:
	case container.EP_SHOW_VIEWS:
	case container.EP_SHOW_INDEXS:
	case container.EP_SHOW_TRIGGERS:
	case container.EP_SHOW_FUNCTIONS:
	case container.EP_SHOW_PROCEDURES:
	case container.EP_CREATE_USER:
	case container.EP_LOG_USER:
	case container.EP_PSM_CALL:
	}

	return ""
}

//initialze the whole system, create key tables and insert necessary tuples into these tables
func (ee *ExecutionEngine) InitializeSystem() {
	ee.ktm.InitializeSystem()
}

//create a new table
//insert a tuple into key table 9: k_table
//insert a tuple into key table 2: k_tableId_userId
//insert a tuple into key table 8: k_tableId_schema
//assign an empty headPage for this table
func (ee *ExecutionEngine) CreateTableOperator(userId int32, schemaString string) {
	transaction := storage.GetTransaction()

	//insert a tuple into key table 9: k_table
	//new tableId is lastTupleId+1 in k_table
	_, _, lastTupleId9, _, _ := ee.ktm.Query_k_table(9)
	newTableId := lastTupleId9 + 1
	newHeadPageId := ee.ktm.GetVacantDataPageId()
	ee.ktm.Insert_k_table(newTableId, newHeadPageId, newHeadPageId, 0, 0)

	//update newHeadPage
	newHeadPage := storage.NewDataPageMode0(newHeadPageId, newTableId, newHeadPageId, newHeadPageId)
	transaction.InsertDataPage(newHeadPage)
	ee.se.InsertDataPage(newHeadPage)

	//insert a tuple into key table 2: k_tableId_userId
	_, _, lastTupleId2, _, _ := ee.ktm.Query_k_table(2)
	schema2 := ee.ktm.GetKeyTableSchema(2)
	fields20, _ := container.NewFieldFromBytes(utility.INTToBytes(int32(newTableId)))
	fields21, _ := container.NewFieldFromBytes(utility.INTToBytes(userId))
	var fields2 []*container.Field
	fields2 = append(fields2, fields20)
	fields2 = append(fields2, fields21)
	tuple2, _ := container.NewTuple(2, lastTupleId2+1, schema2, fields2)
	ee.tm.InsertTupleIntoTable(2, tuple2)

	//insert a tuple into key table 8: k_tableId_schema
	ee.ktm.Insert_k_tableId_schema(newTableId, schemaString)
}

//drop a table by its name
//delete a tuple in key table 9: k_table
//delete a tuple in key table 2: k_tableId_userId
//delete a tuple in key table 8: k_tableId_schema
//delete all pages belong to this table
//TODO update: use index to accelerate
func (ee *ExecutionEngine) DropTableOperator(tableName string) {
	var tableId int32
	var tableSchema *container.Schema

	//delete a tuple in key table 8: k_tableId_schema
	//get tableId and tableSchema
	schema8 := ee.ktm.GetKeyTableSchema(8)
	seqIt8 := subOperator.NewSequentialFileReaderIterator(8, schema8)
	seqIt8.Open(nil, nil)
	for seqIt8.HasNext() {
		tuple8, _ := seqIt8.GetNext()
		bytes0, _ := tuple8.TupleGetFieldValue(0)
		bytes1, _ := tuple8.TupleGetFieldValue(1)
		tableId8, _ := utility.BytesToInteger(bytes0)
		tableSchemaString8, _ := utility.BytesToVARCHAR(bytes1)
		ast, _ := ee.parser.ParseSql(tableSchemaString8)
		tableSchema8, _ := ee.rewriter.ASTNodeToSchema(ast)
		if tableSchema8.GetSchemaTableName() == tableName {
			tableId = tableId8
			tableSchema = tableSchema8
			ee.ktm.Delete_k_tableId_schema(uint32(tableId))
			break
		}
	}

	//if tableSchema is nil, no such table in database, return immediately
	if tableSchema == nil {
		return
	}

	//get headPageId of this table
	headPageId, _, _, _, _ := ee.ktm.Query_k_table(uint32(tableId))

	//delete a tuple in key table 9: k_table
	ee.ktm.Delete_k_table(uint32(tableId))

	//delete a tuple in key table 2: k_tableId_userId
	schema2 := ee.ktm.GetKeyTableSchema(2)
	seqIt2 := subOperator.NewSequentialFileReaderIterator(2, schema2)
	seqIt2.Open(nil, nil)
	for seqIt2.HasNext() {
		tuple2, _ := seqIt2.GetNext()
		bytes0, _ := tuple2.TupleGetFieldValue(0)
		tableId2, _ := utility.BytesToINT(bytes0)
		if tableId2 == tableId {
			tupleId2 := tuple2.TupleGetTupleId()
			ee.tm.DeleteTupleFromTable(2, tupleId2)
			break
		}
	}

	//delete all pages belong to this table
	headPage, _ := ee.se.GetDataPage(headPageId, tableSchema)
	nextPageId, _ := headPage.DpGetNextPageId()
	ee.ktm.InsertVacantDataPageId(headPage.DpGetPageId())

	var currentPage *storage.DataPage
	if headPage.DpGetPageId() == nextPageId {
		return
	} else {
		currentPage, _ = ee.se.GetDataPage(nextPageId, tableSchema)
	}

	for {
		ee.ktm.InsertVacantDataPageId(currentPage.DpGetPageId())

		if currentPage.DataPageMode() == 1 {
			nextLinkPageId, _ := currentPage.DpGetLinkNextPageId()
			linkPage, _ := ee.se.GetDataPage(nextLinkPageId, tableSchema)
			for {
				ee.ktm.InsertVacantDataPageId(linkPage.DpGetPageId())
				nextLinkPageId, _ = linkPage.DpGetLinkNextPageId()
				if nextLinkPageId == linkPage.DpGetPageId() {
					break
				} else {
					linkPage, _ = ee.se.GetDataPage(nextLinkPageId, tableSchema)
				}
			}
		}

		nextPageId, _ := currentPage.DpGetNextPageId()
		if nextPageId == currentPage.DpGetPageId() {
			return
		} else {
			currentPage, _ = ee.se.GetDataPage(nextPageId, tableSchema)
		}
	}
}

//insert
func (ee *ExecutionEngine) InsertOperator(tableName string, fieldStrings []string) {
	tableId, tableSchema, _ := ee.ktm.Query_k_tableId_schema_FromTableName(tableName)

	var fields []*container.Field
	for i, fieldString := range fieldStrings {
		domain, _ := tableSchema.GetSchemaDomain(i)
		var dataBytes []byte
		switch domain.GetDomainType() {
		case container.CHAR:
			dataBytes, _ = utility.CHARToBytes(fieldString)
		case container.VARCHAR:
			dataBytes, _ = utility.VARCHARToBytes(fieldString)
		case container.BIT:
			dataBytes = []byte(fieldString)
		case container.BITVARYING:
			dataBytes = []byte(fieldString)
		case container.BOOLEAN:
			var boolValue bool
			if fieldString == "FALSE" {
				boolValue = false
			} else {
				boolValue = true
			}
			dataByte := utility.BOOLEANToByte(boolValue)
			dataBytes[0] = dataByte
		case container.INT:
			i, _ := strconv.ParseInt(fieldString, 10, 32)
			dataBytes = utility.INTToBytes(int32(i))
		case container.INTEGER:
			i, _ := strconv.ParseInt(fieldString, 10, 32)
			dataBytes = utility.IntegerToBytes(int32(i))
		case container.SHORTINT:
			i, _ := strconv.ParseInt(fieldString, 10, 16)
			dataBytes = utility.SHORTINTToBytes(int16(i))
		case container.FLOAT:
			f, _ := strconv.ParseFloat(fieldString, 32)
			dataBytes = utility.FLOATToBytes(float32(f))
		case container.REAL:
			f, _ := strconv.ParseFloat(fieldString, 32)
			dataBytes = utility.REALToBytes(float32(f))
		case container.DOUBLEPRECISION:
			f, _ := strconv.ParseFloat(fieldString, 64)
			dataBytes = utility.DOUBLEPRECISIONToBytes(f)
		case container.DECIMAL:
			f, _ := strconv.ParseFloat(fieldString, 64)
			n, _ := domain.GetDomainN()
			d, _ := domain.GetDomainD()
			dataBytes, _ = utility.DECIMALToBytes(f, int(n), int(d))
		case container.NUMERIC:
			f, _ := strconv.ParseFloat(fieldString, 64)
			n, _ := domain.GetDomainN()
			d, _ := domain.GetDomainD()
			dataBytes, _ = utility.NUMERICToBytes(f, int(n), int(d))
		case container.DATE:
			dataBytes, _ = utility.DATEToBytes(fieldString)
		case container.TIME:
			dataBytes, _ = utility.TIMEToBytes(fieldString)
		}

		newField, _ := container.NewFieldFromBytes(dataBytes)
		fields = append(fields, newField)
	}

	newTuple, _ := container.NewTuple(tableId, 0, tableSchema, fields)

	ee.tm.InsertTupleIntoTable(tableId, newTuple)
}

//generate a tree of iterators from physicalPlan and execute it
//return resultSchema and resultTuples
func (ee *ExecutionEngine) QueryOperator(executionPlan *container.ExecutionPlan) (*container.Schema, []*container.Tuple) {
	return nil, nil
}

func (ee *ExecutionEngine) DeleteOperator(tableName string, tupleId uint32) {
	tableId, _, _ := ee.ktm.Query_k_tableId_schema_FromTableName(tableName)
	ee.tm.DeleteTupleFromTable(tableId, tupleId)
}

func (ee *ExecutionEngine) UpdateOperator() {}

//TODO
func (ee *ExecutionEngine) AlterTableAddOperator()  {}
func (ee *ExecutionEngine) AlterTableDropOperator() {}
func (ee *ExecutionEngine) CreateAssertOperator()   {}
func (ee *ExecutionEngine) DropAssertOperator()     {}
func (ee *ExecutionEngine) CreateViewOperator()     {}
func (ee *ExecutionEngine) DropViewOperator()       {}
func (ee *ExecutionEngine) CreateIndexOperator()    {}
func (ee *ExecutionEngine) DropIndexOperator()      {}
func (ee *ExecutionEngine) CreateTriggerOperator()  {}
func (ee *ExecutionEngine) DropTriggerOperator()    {}
func (ee *ExecutionEngine) CreatePsmOperator()      {}
func (ee *ExecutionEngine) DropPsmOperator()        {}
func (ee *ExecutionEngine) ShowTablesOperator()     {}
func (ee *ExecutionEngine) ShowAssertionsOperator() {}
func (ee *ExecutionEngine) ShowViewsOperator()      {}
func (ee *ExecutionEngine) ShowIndexsOperator()     {}
func (ee *ExecutionEngine) ShowTriggersOperator()   {}
func (ee *ExecutionEngine) ShowFunctionsOperator()  {}
func (ee *ExecutionEngine) ShowProceduresOperator() {}
func (ee *ExecutionEngine) CreateUserOperator()     {}
func (ee *ExecutionEngine) LogUserOperator()        {}
func (ee *ExecutionEngine) PsmCallOperator()        {}
