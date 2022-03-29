package execution

import (
	"ZetaDB/container"
	its "ZetaDB/execution/querySubOperator"
	"ZetaDB/optimizer"
	"ZetaDB/parser"
	"ZetaDB/storage"
	. "ZetaDB/utility"
	"errors"
)

type KeytableManager struct {
	parser   *parser.Parser
	rewriter *optimizer.Rewriter
	se       *storage.StorageEngine
}

func NewKeytableManager(parser *parser.Parser, rewriter *optimizer.Rewriter, se *storage.StorageEngine) *KeytableManager {
	ktm := &KeytableManager{
		parser:   parser,
		rewriter: rewriter,
		se:       se}
	return ktm
}

//get table info (tableId & schema) from dataFile according to tableName, k_tableId_schema table 8
//throw error if no such table
func (ktm *KeytableManager) GetTableInfo(tableName string) (uint32, *container.Schema, error) {
	astOfCreateTable8 := ktm.parser.ParseSql(DEFAULT_KEYTABLES_SCHEMA[8])
	schemaOfCreateTable8, err := ktm.rewriter.ASTNodeToSchema(astOfCreateTable8)
	if err != nil {
		return 0, nil, err
	}

	seqIt := its.NewSequentialFileReaderIterator(ktm.se, 8, schemaOfCreateTable8)
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
		ast := ktm.parser.ParseSql(schemaString)
		currentSchema, err := ktm.rewriter.ASTNodeToSchema(ast)
		if err != nil {
			return 0, nil, err
		}

		//found correct table schema
		if currentSchema.GetSchemaTableName() == tableName {
			return uint32(tableId), currentSchema, nil
		}

	}
	return 0, nil, errors.New("execution/executionEngine.go    GetSchemaFromFileByTableName() no such table")
}
