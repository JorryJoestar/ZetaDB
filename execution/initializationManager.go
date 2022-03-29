package execution

import (
	"ZetaDB/container"
	"ZetaDB/optimizer"
	"ZetaDB/parser"
	"ZetaDB/storage"
	"ZetaDB/utility"
)

type InitializationManager struct {
	se       *storage.StorageEngine
	parser   *parser.Parser
	rewriter *optimizer.Rewriter
}

func NewInitializationManager(se *storage.StorageEngine, parser *parser.Parser, rewriter *optimizer.Rewriter) *InitializationManager {
	im := &InitializationManager{
		se:       se,
		parser:   parser,
		rewriter: rewriter}

	return im
}

func (im *InitializationManager) InitializeSystem() {
	//erase dataPage, indexPage and logPage
	im.se.EraseDataFile()
	im.se.EraseIndexFile()
	im.se.EraseLogFile()

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
	table0CreateAst := im.parser.ParseSql(utility.DEFAULT_KEYTABLES_SCHEMA[0])
	table0Schema, _ := im.rewriter.ASTNodeToSchema(table0CreateAst)

	field000, _ := container.NewFieldFromBytes(utility.INTToBytes(int32(utility.DEFAULT_ADMINISTRATOR_USER_ID)))
	userNameBytes, _ := utility.VARCHARToBytes(utility.DEFAULT_ADMINISTRATOR_NAME)
	field001, _ := container.NewFieldFromBytes(userNameBytes)

	var fields00 []*container.Field
	fields00 = append(fields00, field000)
	fields00 = append(fields00, field001)

	tuple00, _ := container.NewTuple(0, 0, table0Schema, fields00)
	p0.InsertTuple(tuple00)

	//insert tuple into page1
	//DEFAULT_ADMINISTRATOR_USER_ID, DEFAULT_ADMINISTRATOR_PASSWORD
	table1CreateAst := im.parser.ParseSql(utility.DEFAULT_KEYTABLES_SCHEMA[1])
	table1Schema, _ := im.rewriter.ASTNodeToSchema(table1CreateAst)

	field100, _ := container.NewFieldFromBytes(utility.INTToBytes(int32(utility.DEFAULT_ADMINISTRATOR_USER_ID)))
	userPasswordBytes, _ := utility.VARCHARToBytes(utility.DEFAULT_ADMINISTRATOR_PASSWORD)
	field101, _ := container.NewFieldFromBytes(userPasswordBytes)

	var fields10 []*container.Field
	fields10 = append(fields10, field100)
	fields10 = append(fields10, field101)

	tuple10, _ := container.NewTuple(1, 0, table1Schema, fields10)
	p1.InsertTuple(tuple10)

	//insert tuple into page2
	//all key tabls belong to administrator
	//0 to 16, DEFAULT_ADMINISTRATOR_USER_ID
	table2CreateAst := im.parser.ParseSql(utility.DEFAULT_KEYTABLES_SCHEMA[2])
	table2Schema, _ := im.rewriter.ASTNodeToSchema(table2CreateAst)

	f2x1, _ := container.NewFieldFromBytes(utility.INTToBytes(int32(utility.DEFAULT_ADMINISTRATOR_USER_ID)))

	for i := 0; i <= 16; i++ {
		f2x0, _ := container.NewFieldFromBytes(utility.INTToBytes(int32(i)))
		var fields []*container.Field
		fields = append(fields, f2x0)
		fields = append(fields, f2x1)
		newTuple, _ := container.NewTuple(2, uint32(i), table2Schema, fields)
		p2.InsertTuple(newTuple)
	}

	//insert tuple into page8
	//0 to 16, DEFAULT_KEY_TABLE_0_SCHEMA to DEFAULT_KEY_TABLE_16_SCHEMA
	table8CreateAst := im.parser.ParseSql(utility.DEFAULT_KEYTABLES_SCHEMA[8])
	table8Schema, _ := im.rewriter.ASTNodeToSchema(table8CreateAst)

	for i := 0; i <= 16; i++ {
		f8x0, _ := container.NewFieldFromBytes(utility.INTToBytes(int32(i)))
		f8x1Bytes, _ := utility.VARCHARToBytes(utility.DEFAULT_KEYTABLES_SCHEMA[i])
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
	table9CreateAst := im.parser.ParseSql(utility.DEFAULT_KEYTABLES_SCHEMA[9])
	table9Schema, _ := im.rewriter.ASTNodeToSchema(table9CreateAst)

	for i := 0; i <= 16; i++ {

		f9x0, _ := container.NewFieldFromBytes(utility.INTToBytes(int32(i)))
		f9x1, _ := container.NewFieldFromBytes(utility.INTToBytes(int32(i)))
		var f9x2 *container.Field
		var f9x3 *container.Field
		var fields9 []*container.Field

		if i == 0 || i == 1 || i == 15 || i == 16 { //these tables have one tuple
			//k_userId_userName, k_userId_password, k_emptyDataPageSlot, k_emptyIndexPageSlot
			f9x2, _ = container.NewFieldFromBytes(utility.INTToBytes(int32(0)))
			f9x3, _ = container.NewFieldFromBytes(utility.INTToBytes(int32(1)))
		} else if i == 2 || i == 8 || i == 9 { //these tables have 17 tuples
			//k_tableId_userId, k_tableId_schema, k_table
			f9x2, _ = container.NewFieldFromBytes(utility.INTToBytes(int32(16)))
			f9x3, _ = container.NewFieldFromBytes(utility.INTToBytes(int32(17)))
		} else {
			//all other key tables
			f9x2, _ = container.NewFieldFromBytes(utility.INTToBytes(int32(0)))
			f9x3, _ = container.NewFieldFromBytes(utility.INTToBytes(int32(0)))
		}

		fields9 = append(fields9, f9x0)
		fields9 = append(fields9, f9x1)
		fields9 = append(fields9, f9x2)
		fields9 = append(fields9, f9x3)

		tuple9, _ := container.NewTuple(9, uint32(i), table9Schema, fields9)
		p9.InsertTuple(tuple9)
	}

	//insert tuple into page15
	//tuple0 is important: it keeps a pageId, for all dataPages whose pageId >= this value, they are not allocated
	//the first 17 pages have been occupied by key tables
	//17
	table15CreateAst := im.parser.ParseSql(utility.DEFAULT_KEYTABLES_SCHEMA[15])
	table15Schema, _ := im.rewriter.ASTNodeToSchema(table15CreateAst)

	field1500, _ := container.NewFieldFromBytes(utility.INTToBytes(17))
	var fields15 []*container.Field
	fields15 = append(fields15, field1500)
	tuple150, _ := container.NewTuple(15, 0, table15Schema, fields15)
	p15.InsertTuple(tuple150)

	//insert tuple into page16
	//tuple0 is important: it keeps a pageId, for all indexPages whose pageId >= this value, they are not allocated
	//0
	table16CreateAst := im.parser.ParseSql(utility.DEFAULT_KEYTABLES_SCHEMA[16])
	table16Schema, _ := im.rewriter.ASTNodeToSchema(table16CreateAst)

	field1600, _ := container.NewFieldFromBytes(utility.INTToBytes(0))
	var fields16 []*container.Field
	fields16 = append(fields16, field1600)
	tuple160, _ := container.NewTuple(16, 0, table16Schema, fields16)
	p16.InsertTuple(tuple160)

	//insert these pages into storageEngine
	im.se.InsertDataPage(p0)
	im.se.InsertDataPage(p1)
	im.se.InsertDataPage(p2)
	im.se.InsertDataPage(p8)
	im.se.InsertDataPage(p9)
	im.se.InsertDataPage(p15)
	im.se.InsertDataPage(p16)

	//swap these pages into disk
	im.se.SwapDataPage(0)
	im.se.SwapDataPage(1)
	im.se.SwapDataPage(2)
	im.se.SwapDataPage(8)
	im.se.SwapDataPage(9)
	im.se.SwapDataPage(15)
	im.se.SwapDataPage(16)
}
