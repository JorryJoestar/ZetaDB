package parser

import (
	"strconv"
)

func ASTToString(ast *ASTNode) string {
	if ast == nil {
		return "ast nil"
	}

	s := "ASTNode\n"
	tabs := ""

	switch ast.Type {
	case AST_DDL:
		s += "Type: AST_DDL\n"
		s += DDLToString(ast.Ddl, tabs+"\t")
	case AST_DML:
		s += "Type: AST_DML\n"
		s += DMLToString(ast.Dml, tabs+"\t")
	case AST_DCL:
		s += "Type: AST_DCL\n"
		s += DCLToString(ast.Dcl, tabs+"\t")
	case AST_DQL:
		s += "Type: AST_DQL\n"
		s += DQLToString(ast.Dql, tabs+"\t")
	}
	return s
}

func DDLToString(ddl *DDLNode, tabs string) string {
	s := ""

	s += tabs
	s += "DDLNode\n"

	switch ddl.Type {
	case DDL_TABLE_CREATE:
		s += tabs
		s += "Type: DDL_TABLE_CREATE\n"
		s += TableCreateToString(ddl.Table, tabs+"\t")
	case DDL_TABLE_DROP:
		s += tabs
		s += "Type: DDL_TABLE_DROP\n"
		s += TableDropToString(ddl.Table, tabs+"\t")
	case DDL_TABLE_ALTER_ADD:
		s += tabs
		s += "Type: DDL_TABLE_ALTER_ADD\n"
		s += TableAlterAddToString(ddl.Table, tabs+"\t")
	case DDL_TABLE_ALTER_DROP:
		s += tabs
		s += "Type: DDL_TABLE_ALTER_DROP\n"
		s += TableAlterDropToString(ddl.Table, tabs+"\t")
	case DDL_ASSERT_CREATE:
		s += tabs
		s += "Type: DDL_ASSERT_CREATE\n"
		s += AssertCreateToString(ddl.Assert, tabs+"\t")
	case DDL_ASSERT_DROP:
		s += tabs
		s += "Type: DDL_ASSERT_DROP\n"
		s += AssertDropToString(ddl.Assert, tabs+"\t")
	case DDL_VIEW_CREATE:
		s += tabs
		s += "Type: DDL_VIEW_CREATE\n"
		s += ViewCreateToString(ddl.View, tabs+"\t")
	case DDL_VIEW_DROP:
		s += tabs
		s += "Type: DDL_VIEW_DROP\n"
		s += ViewDropToString(ddl.View, tabs+"\t")
	case DDL_INDEX_CREATE:
		s += tabs
		s += "Type: DDL_INDEX_CREATE\n"
		s += IndexCreateToString(ddl.Index, tabs+"\t")
	case DDL_INDEX_DROP:
		s += tabs
		s += "Type: DDL_INDEX_DROP\n"
		s += IndexDropToString(ddl.Index, tabs+"\t")
	case DDL_TRIGGER_CREATE:
		s += tabs
		s += "Type: DDL_TRIGGER_CREATE\n"
		s += TriggerCreateToString(ddl.Trigger, tabs+"\t")
	case DDL_TRIGGER_DROP:
		s += tabs
		s += "Type: DDL_TRIGGER_DROP\n"
		s += TriggerDropToString(ddl.Trigger, tabs+"\t")
	case DDL_PSM_CREATE:
		s += tabs
		s += "Type: DDL_PSM_CREATE\n"
		s += PsmCreateToString(ddl.Psm, tabs+"\t")
	case DDL_PSM_DROP:
		s += tabs
		s += "Type: DDL_PSM_DROP\n"
		s += PsmDropToString(ddl.Psm, tabs+"\t")
	}

	return s
}

func DMLToString(dml *DMLNode, tabs string) string {
	s := tabs + "DMLNode\n"
	switch dml.Type {
	case DML_INSERT:
		s += InsertToString(dml.Insert, tabs+"\t")
	case DML_UPDATE:
		s += UpdateToString(dml.Update, tabs+"\t")
	case DML_DELETE:
		s += DeleteToString(dml.Delete, tabs+"\t")

	}
	return s
}

func InsertToString(insert *InsertNode, tabs string) string {
	s := tabs + "InsertNode\n"
	s += tabs + "TableName: " + insert.TableName + "\n"

	if insert.AttriNameListValid {
		s += tabs + "AttriNameList:\n"
		for _, v := range insert.AttriNameList {
			s += tabs + "\t" + v + "\n"
		}
	}

	switch insert.Type {
	case INSERT_FROM_SUBQUERY:
		s += tabs + "Type: INSERT_FROM_SUBQUERY\n"
		s += QueryToString(insert.Query, tabs+"\t")
	case INSERT_FROM_VALUELIST:
		s += tabs + "Type: INSERT_FROM_VALUELIST\n"
		s += tabs + "ElementaryValueList:\n"
		for _, v := range insert.ElementaryValueList {
			s += ElementaryValueToString(v, tabs+"\t")
		}
	}
	return s
}

func UpdateToString(update *UpdateNode, tabs string) string {
	s := tabs + "UpdateNode\n"

	s += tabs + "TableName: " + update.TableName + "\n"

	s += tabs + "Condition:\n"
	s += ConditionToString(update.Condition, tabs+"\t")

	s += tabs + "UpdateList:\n"
	for _, v := range update.UpdateList {
		s += UpdateListEntryToString(v, tabs+"\t")
	}
	return s
}

func UpdateListEntryToString(entry *UpdateListEntryNode, tabs string) string {
	s := tabs + "UpdateListEntryNode\n"
	s += tabs + "AttributeName" + entry.AttributeName + "\n"

	switch entry.Type {
	case UPDATE_LIST_ENTRY_EXPRESSION:
		s += tabs + "Type: UPDATE_LIST_ENTRY_EXPRESSION\n"
		s += tabs + "Expression:\n"
		s += ExpressionToString(entry.Expression, tabs+"\t")
	case UPDATE_LIST_ENTRY_ELEMENTARY_VALUE:
		s += tabs + "Type: UPDATE_LIST_ENTRY_ELEMENTARY_VALUE\n"
		s += tabs + "ElementaryValue:\n"
		s += ElementaryValueToString(entry.ElementaryValue, tabs+"\t")
	}
	return s
}

func DeleteToString(delete *DeleteNode, tabs string) string {
	s := tabs + "DeleteNode\n"
	s += tabs + "TableName: " + delete.TableName + "\n"
	s += ConditionToString(delete.Condition, tabs+"\t")
	return s
}

func DCLToString(dcl *DCLNode, tabs string) string {
	s := tabs + "DCLNode\n"
	switch dcl.Type {
	case DCL_TRANSACTION_BEGIN:
		s += tabs + "Type: DCL_TRANSACTION_BEGIN\n"
	case DCL_TRANSACTION_COMMIT:
		s += tabs + "Type: DCL_TRANSACTION_COMMIT\n"
	case DCL_TRANSACTION_ROLLBACK:
		s += tabs + "Type: DCL_TRANSACTION_ROLLBACK\n"
	case DCL_SHOW_TABLES:
		s += tabs + "Type: DCL_SHOW_TABLES\n"
	case DCL_SHOW_ASSERTIONS:
		s += tabs + "Type: DCL_SHOW_ASSERTIONS\n"
	case DCL_SHOW_VIEWS:
		s += tabs + "Type: DCL_SHOW_VIEWS\n"
	case DCL_SHOW_INDEXS:
		s += tabs + "Type: DCL_SHOW_INDEXS\n"
	case DCL_SHOW_TRIGGERS:
		s += tabs + "Type: DCL_SHOW_TRIGGERS\n"
	case DCL_SHOW_FUNCTIONS:
		s += tabs + "Type: DCL_SHOW_FUNCTIONS\n"
	case DCL_SHOW_PROCEDURES:
		s += tabs + "Type: DCL_SHOW_PROCEDURES\n"
	case DCL_CREATE_USER:
		s += tabs + "Type: DCL_CREATE_USER\n"
	case DCL_LOG_USER:
		s += tabs + "Type: DCL_LOG_USER\n"
	case DCL_PSMCALL:
		s += tabs + "Type: DCL_PSMCALL\n"
	}
	if dcl.Type == DCL_CREATE_USER || dcl.Type == DCL_LOG_USER {
		s += tabs + "UserName: " + dcl.UserName + "\n"
		s += tabs + "Password: " + dcl.Password + "\n"
	}
	if dcl.Type == DCL_PSMCALL {
		s += PsmCallStmtToString(dcl.PsmCall, tabs+"\t")
	}
	return s
}

func DQLToString(dql *DQLNode, tabs string) string {
	s := tabs + "DQLNode\n"
	switch dql.Type {
	case DQL_DIFFERENCE:
		s += tabs + "Type: DQL_DIFFERENCE\n"
	case DQL_INTERSECTION:
		s += tabs + "Type: DQL_INTERSECTION\n"
	case DQL_UNION:
		s += tabs + "Type: DQL_UNION\n"
	case DQL_SINGLE_QUERY:
		s += tabs + "Type: DQL_SINGLE_QUERY\n"
	}
	if dql.Type == DQL_SINGLE_QUERY {
		s += tabs + "Query:\n"
		s += QueryToString(dql.Query, tabs+"\t")
	} else {
		s += tabs + "DqlL:\n"
		s += DQLToString(dql.DqlL, tabs+"\t")
		s += tabs + "DqlR:\n"
		s += DQLToString(dql.DqlR, tabs+"\t")
	}
	s += "\n"
	return s
}

func TableCreateToString(table *TableNode, tabs string) string {
	s := ""
	s += tabs
	s += "TableNode\n"

	s += tabs
	s += "TableName: " + table.TableName + "\n"

	s += tabs
	s += "AttributeNameList:\n"
	for _, v := range table.AttributeNameList {
		s += tabs + "\t"
		s += v + "\n"
	}

	s += tabs
	s += "DomainList:\n"
	for _, v := range table.DomainList {
		s += DomainToString(v, tabs+"\t")
	}

	if table.ConstraintListValid {
		s += tabs
		s += "ConstraintList:\n"
		for _, v := range table.ConstraintList {
			s += ConstraintToString(v, tabs+"\t")
		}
	}

	return s
}

func TableDropToString(table *TableNode, tabs string) string {
	s := tabs + "TableNode\n"
	s += tabs + "TableName: " + table.TableName + "\n"
	return s
}

func TableAlterAddToString(table *TableNode, tabs string) string {
	s := tabs + "TableNode\n"
	s += tabs + "TableName: " + table.TableName + "\n"
	if len(table.AttributeNameList) != 0 {
		s += tabs + "AttributeNameList:\n"
		for _, v := range table.AttributeNameList {
			s += tabs + "\t" + v + "\n"
		}
	}
	if len(table.DomainList) != 0 {
		s += tabs + "DomainList:\n"
		for _, v := range table.DomainList {
			s += DomainToString(v, tabs+"\t")
		}
	}
	if table.ConstraintListValid {
		s += tabs + "ConstraintList:\n"
		for _, v := range table.ConstraintList {
			s += ConstraintToString(v, tabs+"\t")
		}
	}
	return s
}

func TableAlterDropToString(table *TableNode, tabs string) string {
	s := tabs + "TableNode\n"
	s += tabs + "TableName: " + table.TableName + "\n"
	if len(table.AttributeNameList) != 0 {
		s += tabs + "AttributeNameList:\n"
		for _, v := range table.AttributeNameList {
			s += tabs + "\t" + v + "\n"
		}
	} else {
		s += tabs + "ConstraintName: " + table.ConstraintName + "\n"
	}
	return s
}

func AssertCreateToString(assert *AssertNode, tabs string) string {
	s := tabs + "AssertNode\n"
	s += tabs + "AssertName: " + assert.AssertName + "\n"
	s += tabs + "Condition:\n"
	s += ConditionToString(assert.Condition, tabs+"\t")
	return s
}

func AssertDropToString(assert *AssertNode, tabs string) string {
	s := tabs + "AssertNode\n"
	s += tabs + "AssertName: " + assert.AssertName + "\n"
	return s
}

func ViewCreateToString(view *ViewNode, tabs string) string {
	s := tabs + "ViewNode\n"
	s += tabs + "ViewName: " + view.ViewName + "\n"
	if view.AttributeNameListValid {
		s += tabs + "AttributeNameList:\n"
		for _, v := range view.AttributeNameList {
			s += tabs + "\t" + v + "\n"
		}
	}
	s += tabs + "Query:\n"
	s += QueryToString(view.Query, tabs+"\t")
	return s
}

func ViewDropToString(view *ViewNode, tabs string) string {
	s := tabs + "ViewNode\n"
	s += tabs + "ViewName: " + view.ViewName + "\n"
	return s
}

func IndexCreateToString(index *IndexNode, tabs string) string {
	s := tabs + "IndexNode\n"
	s += tabs + "IndexName: " + index.IndexName + "\n"
	s += tabs + "TableName: " + index.TableName + "\n"
	s += tabs + "AttributeNameList:\n"
	for _, v := range index.AttributeNameList {
		s += tabs + "\t" + v + "\n"
	}
	return s
}

func IndexDropToString(index *IndexNode, tabs string) string {
	s := tabs + "IndexNode\n"
	s += tabs + "IndexName: " + index.IndexName + "\n"
	return s
}

func TriggerCreateToString(trigger *TriggerNode, tabs string) string {
	s := tabs + "TriggerNode\n"
	s += tabs + "BeforeAfterType: "
	switch trigger.BeforeAfterType {
	case BEFORE_UPDATE_OF:
		s += "BEFORE_UPDATE_OF\n"
		s += tabs + "BeforeAfterAttriName: " + trigger.BeforeAfterAttriName + "\n"
	case BEFORE_UPDATE:
		s += "BEFORE_UPDATE\n"
	case AFTER_UPDATE_OF:
		s += "AFTER_UPDATE_OF\n"
		s += tabs + "BeforeAfterAttriName: " + trigger.BeforeAfterAttriName + "\n"
	case AFTER_UPDATE:
		s += "AFTER_UPDATE\n"
	case INSTEAD_UPDATE_OF:
		s += "INSTEAD_UPDATE_OF\n"
		s += tabs + "BeforeAfterAttriName: " + trigger.BeforeAfterAttriName + "\n"
	case INSTEAD_UPDATE:
		s += "INSTEAD_UPDATE\n"
	case BEFORE_INSERT:
		s += "BEFORE_INSERT\n"
	case AFTER_INSERT:
		s += "AFTER_INSERT\n"
	case INSTEAD_INSERT:
		s += "INSTEAD_INSERT\n"
	case BEFORE_DELETE:
		s += "BEFORE_DELETE\n"
	case AFTER_DELETE:
		s += "AFTER_DELETE\n"
	case INSTEAD_DELETE:
		s += "INSTEAD_DELETE\n"
	}
	s += tabs + "BeforeAfterTableName: " + trigger.BeforeAfterTableName + "\n"

	if trigger.ReferencingValid {
		s += tabs + "OldNewList:\n"
		for _, v := range trigger.OldNewList {
			s += TriggerOldNewEntryToString(v, tabs+"\t")
		}
	}

	switch trigger.ForEachType {
	case FOR_EACH_ROW:
		s += tabs + "ForEachType: FOR_EACH_ROW\n"
	case FOR_EACH_STATEMENT:
		s += tabs + "ForEachType: FOR_EACH_STATEMENT\n"
	}

	if trigger.WhenValid {
		s += tabs + "Condition:\n"
		s += ConditionToString(trigger.Condition, tabs+"\t")
	}

	s += tabs + "DmlList:\n"
	for _, v := range trigger.DmlList {
		s += DMLToString(v, tabs+"\t")
	}

	return s
}

func TriggerDropToString(trigger *TriggerNode, tabs string) string {
	s := tabs + "TriggerNode\n"
	s += tabs + "TriggerName: " + trigger.TriggerName + "\n"
	return s
}

func PsmCreateToString(psm *PsmNode, tabs string) string {
	s := tabs + "PsmNode\n"

	switch psm.Type {
	case PSM_FUNCTION:
		s += tabs + "Type: PSM_FUNCTION\n"
	case PSM_PROCEDURE:
		s += tabs + "Type: PSM_PROCEDURE\n"
	}

	s += tabs + "PsmName: " + psm.PsmName + "\n"

	if psm.PsmParameterListValid {
		s += tabs + "PsmParameterList:\n"
		for _, v := range psm.PsmParameterList {
			s += PsmParameterEntryToString(v, tabs+"\t")
		}
	}

	if psm.PsmLocalDeclarationListValid {
		s += tabs + "PsmLocalDeclarationList:\n"
		for _, v := range psm.PsmLocalDeclarationList {
			s += PsmLocalDeclarationEntryToString(v, tabs+"\t")
		}
	}

	s += tabs + "PsmBody:\n"
	for _, v := range psm.PsmBody {
		s += PsmExecEntryToString(v, tabs+"\t")
	}

	return s
}

func PsmParameterEntryToString(entry *PsmParameterEntryNode, tabs string) string {
	s := tabs + "PsmParameterEntryNode\n"

	switch entry.Type {
	case PSM_PARAMETER_IN:
		s += tabs + "Type: PSM_PARAMETER_IN\n"
	case PSM_PARAMETER_OUT:
		s += tabs + "Type: PSM_PARAMETER_OUT\n"
	case PSM_PARAMETER_INOUT:
		s += tabs + "Type: PSM_PARAMETER_INOUT\n"
	}

	s += tabs + "Name: " + entry.Name + "\n"

	s += tabs + "Domain:\n"
	s += DomainToString(entry.Domain, tabs+"\t")

	return s
}

func PsmLocalDeclarationEntryToString(entry *PsmLocalDeclarationEntryNode, tabs string) string {
	s := tabs + "PsmLocalDeclarationEntryNode\n"
	s += tabs + "Name: " + entry.Name + "\n"
	s += tabs + "Domain:\n"
	s += DomainToString(entry.Domain, tabs+"\t")

	return s
}

func PsmExecEntryToString(entry *PsmExecEntryNode, tabs string) string {
	s := tabs + "PsmExecEntryNode\n"
	switch entry.Type {
	case PSM_EXEC_RETURN:
		s += tabs + "Type: PSM_EXEC_RETURN\n"
		s += tabs + "PsmValue:\n"
		s += PsmValueToString(entry.PsmValue, tabs+"\t")
	case PSM_EXEC_SET:
		s += tabs + "Type: PSM_EXEC_SET\n"
		s += tabs + "VariableName: " + entry.VariableName + "\n"
		s += tabs + "PsmValue:\n"
		s += PsmValueToString(entry.PsmValue, tabs+"\t")
	case PSM_EXEC_FOR_LOOP:
		s += tabs + "Type: PSM_EXEC_FOR_LOOP\n"
		s += tabs + "PsmForLoop:\n"
		s += PsmForLoopToString(entry.PsmForLoop, tabs+"\t")
	case PSM_EXEC_BRANCH:
		s += tabs + "Type: PSM_EXEC_BRANCH\n"
		s += tabs + "PsmBranch:\n"
		s += PsmBranchToString(entry.PsmBranch, tabs+"\n")
	case PSM_EXEC_DML:
		s += tabs + "Type: PSM_EXEC_DML\n"
		s += DMLToString(entry.Dml, tabs+"\t")
	}
	return s
}

func PsmBranchToString(branch *PsmBranchNode, tabs string) string {
	s := tabs + "PsmBranchNode\n"

	s += tabs + "Condition:\n"
	s += ConditionToString(branch.Condition, tabs+"\t")

	s += tabs + "IfPsmExecList:\n"
	for _, v := range branch.IfPsmExecList {
		s += PsmExecEntryToString(v, tabs+"\t")
	}

	if branch.PsmElseifListValid {
		s += tabs + "PsmElseifList:\n"
		for _, v := range branch.PsmElseifList {
			s += PsmElseifEntryToString(v, tabs+"\t")
		}
	}

	if branch.ElseValid {
		s += tabs + "ElsePsmExecList:\n"
		for _, v := range branch.ElsePsmExecList {
			s += PsmExecEntryToString(v, tabs+"\t")
		}
	}

	return s
}

func PsmElseifEntryToString(entry *PsmElseifEntryNode, tabs string) string {
	s := tabs + "PsmElseifEntryNode\n"

	s += tabs + "Condition:\n"
	s += ConditionToString(entry.Condition, tabs+"\t")

	s += tabs + "PsmExecList:\n"
	for _, v := range entry.PsmExecList {
		s += PsmExecEntryToString(v, tabs+"\t")
	}
	return s
}

func PsmForLoopToString(ForLoop *PsmForLoopNode, tabs string) string {
	s := tabs + "PsmForLoopNode\n"
	s += tabs + "LoopName: " + ForLoop.LoopName + "\n"
	s += tabs + "CursorName" + ForLoop.CursorName + "\n"
	s += tabs + "Query:\n"
	s += QueryToString(ForLoop.Query, tabs+"\t")
	s += tabs + "PsmExecList:\n"
	for _, v := range ForLoop.PsmExecList {
		s += PsmExecEntryToString(v, tabs+"\t")
	}
	return s
}

func PsmValueToString(value *PsmValueNode, tabs string) string {
	s := tabs + "PsmValueNode\n"
	switch value.Type {
	case PSMVALUE_ELEMENTARY_VALUE:
		s += tabs + "Type: PSMVALUE_ELEMENTARY_VALUE\n"
		s += ElementaryValueToString(value.ElementaryValue, tabs+"\t")
	case PSMVALUE_CALL:
		s += tabs + "Type: PSMVALUE_CALL\n"
		s += PsmCallStmtToString(value.PsmCall, tabs+"\t")
	case PSMVALUE_EXPRESSION:
		s += tabs + "Type: PSMVALUE_EXPRESSION\n"
		s += ExpressionToString(value.Expression, tabs+"\t")
	case PSMVALUE_ID:
		s += tabs + "Type: PSMVALUE_ID\n"
		s += tabs + "Id: " + value.Id + "\n"
	}

	return s
}

func ExpressionToString(expression *ExpressionNode, tabs string) string {
	s := tabs + "ExpressionNode\n"

	switch expression.Type {
	case EXPRESSION_OPERATOR_PLUS:
		s += tabs + "Type: EXPRESSION_OPERATOR_PLUS\n"
	case EXPRESSION_OPERATOR_MINUS:
		s += tabs + "Type: EXPRESSION_OPERATOR_MINUS\n"
	case EXPRESSION_OPERATOR_DIVISION:
		s += tabs + "Type: EXPRESSION_OPERATOR_DIVISION\n"
	case EXPRESSION_OPERATOR_MULTIPLY:
		s += tabs + "Type: EXPRESSION_OPERATOR_MULTIPLY\n"
	case EXPRESSION_OPERATOR_CONCATENATION:
		s += tabs + "Type: EXPRESSION_OPERATOR_CONCATENATION\n"
	}

	s += tabs + "ExpressionEntryL:\n"
	s += ExpressionEntryToString(expression.ExpressionEntryL, tabs+"\t")

	s += tabs + "ExpressionEntryR:\n"
	s += ExpressionEntryToString(expression.ExpressionEntryR, tabs+"\t")

	return s
}

func ExpressionEntryToString(entry *ExpressionEntryNode, tabs string) string {
	s := tabs + "ExpressionEntryNode\n"
	switch entry.Type {
	case EXPRESSION_ENTRY_ELEMENTARY_VALUE:
		s += tabs + "Type: EXPRESSION_ENTRY_ELEMENTARY_VALUE\n"
		s += ElementaryValueToString(entry.ElementaryValue, tabs+"\t")
	case EXPRESSION_ENTRY_ATTRIBUTE_NAME:
		s += tabs + "Type: EXPRESSION_ENTRY_ATTRIBUTE_NAME\n"
		s += AttriNameOptionTableNameToString(entry.AttriNameOptionTableName, tabs+"\t")
	case EXPRESSION_ENTRY_AGGREGATION:
		s += tabs + "Type: EXPRESSION_ENTRY_AGGREGATION\n"
		s += AggregationToString(entry.Aggregation, tabs+"\t")
	case EXPRESSION_ENTRY_EXPRESSION:
		s += tabs + "Type: EXPRESSION_ENTRY_EXPRESSION\n"
		s += ExpressionToString(entry.Expression, tabs+"\t")
	}
	return s
}

func AggregationToString(aggregation *AggregationNode, tabs string) string {
	s := tabs + "AggregationNode\n"

	switch aggregation.Type {
	case AGGREGATION_SUM:
		s += tabs + "Type: AGGREGATION_SUM\n"
	case AGGREGATION_AVG:
		s += tabs + "Type: AGGREGATION_AVG\n"
	case AGGREGATION_MIN:
		s += tabs + "Type: AGGREGATION_MIN\n"
	case AGGREGATION_MAX:
		s += tabs + "Type: AGGREGATION_MAX\n"
	case AGGREGATION_COUNT:
		s += tabs + "Type: AGGREGATION_COUNT\n"
	case AGGREGATION_COUNT_ALL:
		s += tabs + "Type: AGGREGATION_COUNT_ALL\n"
	}

	if aggregation.DistinctValid {
		s += tabs + "DistinctValid\n"
	}

	if aggregation.Type != AGGREGATION_COUNT_ALL {
		s += AttriNameOptionTableNameToString(aggregation.AttriNameOptionTableName, tabs+"\t")
	}

	return s

}

func PsmCallStmtToString(psmCall *PsmNode, tabs string) string {
	s := tabs + "PsmNode\n"
	s += tabs + "PsmName: " + psmCall.PsmName + "\n"
	if psmCall.PsmValueListValid {
		s += tabs + "PsmValueList:\n"
		for _, v := range psmCall.PsmValueList {
			s += PsmValueToString(v, tabs+"\t")
		}
	}
	return s
}

func PsmDropToString(psm *PsmNode, tabs string) string {
	s := tabs + "PsmNode\n"

	switch psm.Type {
	case PSM_PROCEDURE:
		s += tabs + "Type: PSM_PROCEDURE\n"
	case PSM_FUNCTION:
		s += tabs + "Type: PSM_FUNCTION\n"
	}

	s += tabs + "PsmName: " + psm.PsmName + "\n"

	return s
}

func DomainToString(domain *DomainNode, tabs string) string {
	s := ""

	s += tabs
	s += "DomainNode\n"

	s += tabs
	s += "Type: "

	switch domain.Type {
	case DOMAIN_CHAR:
		s += "DOMAIN_CHAR\n"
	case DOMAIN_VARCHAR:
		s += "DOMAIN_VARCHAR\n"

		s += tabs
		s += "n: " + strconv.Itoa(domain.N) + "\n"
	case DOMAIN_BIT:
		s += "DOMAIN_BIT\n"

		s += tabs
		s += "n: " + strconv.Itoa(domain.N) + "\n"
	case DOMAIN_BITVARYING:
		s += "DOMAIN_BITVARYING\n"

		s += tabs
		s += "n: " + strconv.Itoa(domain.N) + "\n"
	case DOMAIN_BOOLEAN:
		s += "DOMAIN_BOOLEAN\n"
	case DOMAIN_INT:
		s += "DOMAIN_INT\n"
	case DOMAIN_INTEGER:
		s += "DOMAIN_INTEGER\n"
	case DOMAIN_SHORTINT:
		s += "DOMAIN_SHORTINT\n"
	case DOMAIN_FLOAT:
		s += "DOMAIN_FLOAT\n"
	case DOMAIN_REAL:
		s += "DOMAIN_REAL\n"
	case DOMAIN_DOUBLEPRECISION:
		s += "DOMAIN_DOUBLEPRECISION\n"
	case DOMAIN_DECIMAL:
		s += "DOMAIN_DECIMAL\n"

		s += tabs
		s += "n: " + strconv.Itoa(domain.N) + "\n"

		s += tabs
		s += "d: " + strconv.Itoa(domain.D) + "\n"
	case DOMAIN_NUMERIC:
		s += "DOMAIN_NUMERIC\n"

		s += tabs
		s += "n: " + strconv.Itoa(domain.N) + "\n"

		s += tabs
		s += "d: " + strconv.Itoa(domain.D) + "\n"
	case DOMAIN_DATE:
		s += "DOMAIN_DATE\n"
	case DOMAIN_TIME:
		s += "DOMAIN_TIME\n"
	}

	return s
}

func ConstraintToString(constraint *ConstraintNode, tabs string) string {
	s := ""

	s += tabs
	s += "ConstraintNode\n"

	if constraint.ConstraintNameValid {
		s += tabs
		s += "ConstraintName: "
		s += constraint.ConstraintName + "\n"
	}

	s += tabs
	s += "Type: "

	switch constraint.Type {
	case CONSTRAINT_UNIQUE:
		s += "CONSTRAINT_UNIQUE\n"
		s += tabs
		s += "AttriNameList:\n"
		for _, v := range constraint.AttriNameList {
			s += tabs + "\t"
			s += v + "\n"
		}
	case CONSTRAINT_PRIMARY_KEY:
		s += "CONSTRAINT_PRIMARY_KEY\n"
		s += tabs
		s += "AttriNameList:\n"
		for _, v := range constraint.AttriNameList {
			s += tabs + "\t"
			s += v + "\n"
		}
	case CONSTRAINT_FOREIGN_KEY:
		s += "CONSTRAINT_FOREIGN_KEY\n"
		s += tabs
		s += "AttributeNameLocal: " + constraint.AttributeNameLocal + "\n"
		s += tabs
		s += "AttributeNameForeign: " + constraint.AttributeNameForeign + "\n"
		s += tabs
		s += "ForeignTableName: " + constraint.ForeignTableName + "\n"
		if constraint.DeferrableValid {
			s += tabs
			s += "Deferrable: "
			switch constraint.Deferrable {
			case CONSTRAINT_NOT_DEFERRABLE:
				s += "CONSTRAINT_NOT_DEFERRABLE\n"
			case CONSTRAINT_INITIALLY_DEFERRED:
				s += "CONSTRAINT_INITIALLY_DEFERRED\n"
			case CONSTRAINT_INITIALLY_IMMEDIATE:
				s += "CONSTRAINT_INITIALLY_IMMEDIATE\n"
			}
		}
		if constraint.UpdateSetValid {
			s += tabs
			s += "UpdateSet: "
			switch constraint.UpdateSet {
			case CONSTRAINT_UPDATE_SET_NULL:
				s += "CONSTRAINT_UPDATE_SET_NULL\n"
			case CONSTRAINT_UPDATE_SET_CASCADE:
				s += "CONSTRAINT_UPDATE_SET_CASCADE\n"
			}
		}
		if constraint.DeleteSetValid {
			s += tabs
			s += "DeleteSet: "
			switch constraint.DeleteSet {
			case CONSTRAINT_DELETE_SET_NULL:
				s += "CONSTRAINT_DELETE_SET_NULL\n"
			case CONSTRAINT_DELETE_SET_CASCADE:
				s += "CONSTRAINT_DELETE_SET_CASCADE\n"
			}
		}
	case CONSTRAINT_NOT_NULL:
		s += "CONSTRAINT_NOT_NULL\n"
		s += tabs
		s += "AttributeNameLocal: " + constraint.AttributeNameLocal + "\n"
	case CONSTRAINT_DEFAULT:
		s += "CONSTRAINT_DEFAULT\n"
		s += tabs
		s += "AttributeNameLocal: " + constraint.AttributeNameLocal + "\n"
		s += tabs
		s += "ElementaryValue:\n"
		s += ElementaryValueToString(constraint.ElementaryValue, tabs+"\t")
	case CONSTRAINT_CHECK:
		s += "CONSTRAINT_CHECK\n"
		s += ConditionToString(constraint.Condition, tabs+"\t")
	}

	return s
}

func ElementaryValueToString(elementaryValue *ElementaryValueNode, tabs string) string {
	s := ""

	s += tabs
	s += "ElementaryValueNode\n"
	s += tabs
	s += "Type: "

	switch elementaryValue.Type {
	case ELEMENTARY_VALUE_INT:
		s += "ELEMENTARY_VALUE_INT\n"
		s += tabs
		s += "IntValue: "
		s += strconv.Itoa(elementaryValue.IntValue)
		s += "\n"
	case ELEMENTARY_VALUE_FLOAT:
		s += "ELEMENTARY_VALUE_FLOAT\n"
		s += tabs
		s += "FloatValue: "
		s += strconv.FormatFloat(elementaryValue.FloatValue, 'g', -1, 64)
		s += "\n"
	case ELEMENTARY_VALUE_STRING:
		s += "ELEMENTARY_VALUE_STRING\n"
		s += tabs
		s += "StringValue: "
		s += elementaryValue.StringValue
		s += "\n"
	case ELEMENTARY_VALUE_BOOLEAN:
		s += "ELEMENTARY_VALUE_BOOLEAN\n"
		s += tabs
		s += "BooleanValue: "
		if elementaryValue.BooleanValue {
			s += "true\n"
		} else {
			s += "false\n"
		}
	}
	return s
}

func ConditionToString(condition *ConditionNode, tabs string) string {
	s := tabs
	s += "ConditionNode\n"

	s += tabs
	s += "Type: "
	switch condition.Type {
	case CONDITION_PREDICATE:
		s += "CONDITION_PREDICATE\n"
		s += tabs
		s += "Predicate:\n"
		s += PredicateToString(condition.Predicate, tabs+"\t")
	case CONDITION_AND:
		s += "CONDITION_AND\n"
		s += tabs
		s += "ConditionL:\n"
		s += ConditionToString(condition.ConditionL, tabs+"\t")
		s += tabs
		s += "ConditionR:\n"
		s += ConditionToString(condition.ConditionR, tabs+"\t")
	case CONDITION_OR:
		s += "CONDITION_OR\n"
		s += tabs
		s += "ConditionL:\n"
		s += ConditionToString(condition.ConditionL, tabs)
		s += tabs
		s += "ConditionR:\n"
		s += ConditionToString(condition.ConditionR, tabs)
	}

	return s
}

func PredicateToString(predicate *PredicateNode, tabs string) string {
	s := ""
	s += tabs
	s += "PredicateNode\n"

	s += tabs
	s += "Type: "
	switch predicate.Type {
	case PREDICATE_COMPARE_ELEMENTARY_VALUE:
		s += "PREDICATE_COMPARE_ELEMENTARY_VALUE\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "ElementaryValue:\n"
		s += ElementaryValueToString(predicate.ElementaryValue, tabs+"\t")
		s += tabs
		s += "CompareMark: "
		s += CompareMarkToString(predicate.CompareMark) + "\n"
	case PREDICATE_LIKE_STRING_VALUE:
		s += "PREDICATE_LIKE_STRING_VALUE\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "ElementaryValue:\n"
		s += ElementaryValueToString(predicate.ElementaryValue, tabs+"\t")
	case PREDICATE_IN_SUBQUERY:
		s += "PREDICATE_IN_SUBQUERY\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "Query:\n"
		s += QueryToString(predicate.Query, tabs+"\t")
	case PREDICATE_NOT_IN_SUBQUERY:
		s += "PREDICATE_NOT_IN_SUBQUERY\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "Query:\n"
		s += QueryToString(predicate.Query, tabs+"\t")
	case PREDICATE_IN_TABLE:
		s += "PREDICATE_IN_TABLE\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "TableName: "
		s += predicate.TableName + "\n"
	case PREDICATE_NOT_IN_TABLE:
		s += "PREDICATE_NOT_IN_TABLE\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "TableName: "
		s += predicate.TableName + "\n"
	case PREDICATE_COMPARE_ALL_SUBQUERY:
		s += "PREDICATE_COMPARE_ALL_SUBQUERY\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "Query:\n"
		s += QueryToString(predicate.Query, tabs+"\t")
		s += tabs
		s += "CompareMark: "
		s += CompareMarkToString(predicate.CompareMark) + "\n"
	case PREDICATE_COMPARE_NOT_ALL_SUBQUERY:
		s += "PREDICATE_COMPARE_NOT_ALL_SUBQUERY\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "Query:\n"
		s += QueryToString(predicate.Query, tabs+"\t")
		s += tabs
		s += "CompareMark: "
		s += CompareMarkToString(predicate.CompareMark) + "\n"
	case PREDICATE_COMPARE_ANY_SUBQUERY:
		s += "PREDICATE_COMPARE_ANY_SUBQUERY\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "Query:\n"
		s += QueryToString(predicate.Query, tabs+"\t")
		s += tabs
		s += "CompareMark: "
		s += CompareMarkToString(predicate.CompareMark) + "\n"
	case PREDICATE_COMPARE_NOT_ANY_SUBQUERY:
		s += "PREDICATE_COMPARE_NOT_ANY_SUBQUERY\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "Query:\n"
		s += QueryToString(predicate.Query, tabs+"\t")
		s += tabs
		s += "CompareMark: "
		s += CompareMarkToString(predicate.CompareMark) + "\n"
	case PREDICATE_COMPARE_ALL_TABLE:
		s += "PREDICATE_COMPARE_ALL_TABLE\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "TableName: "
		s += predicate.TableName + "\n"
		s += tabs
		s += "CompareMark: "
		s += CompareMarkToString(predicate.CompareMark) + "\n"
	case PREDICATE_COMPARE_NOT_ALL_TABLE:
		s += "PREDICATE_COMPARE_NOT_ALL_TABLE\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "TableName: "
		s += predicate.TableName + "\n"
		s += tabs
		s += "CompareMark: "
		s += CompareMarkToString(predicate.CompareMark) + "\n"
	case PREDICATE_COMPARE_ANY_TABLE:
		s += "PREDICATE_COMPARE_ANY_TABLE\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "TableName: "
		s += predicate.TableName + "\n"
		s += tabs
		s += "CompareMark: "
		s += CompareMarkToString(predicate.CompareMark) + "\n"
	case PREDICATE_COMPARE_NOT_ANY_TABLE:
		s += "PREDICATE_COMPARE_NOT_ANY_TABLE\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "TableName: "
		s += predicate.TableName + "\n"
		s += tabs
		s += "CompareMark: "
		s += CompareMarkToString(predicate.CompareMark) + "\n"
	case PREDICATE_IS_NULL:
		s += "PREDICATE_IS_NULL\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
	case PREDICATE_IS_NOT_NULL:
		s += "PREDICATE_IS_NOT_NULL\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
	case PREDICATE_TUPLE_IN_SUBQUERY:
		s += "PREDICATE_TUPLE_IN_SUBQUERY\n"
		s += tabs
		s += "AttriNameOptionTableNameList:\n"
		for _, v := range predicate.AttriNameOptionTableNameList {
			s += AttriNameOptionTableNameToString(v, tabs+"\t")
		}
		s += tabs
		s += "Query:\n"
		s += QueryToString(predicate.Query, tabs+"\t")
	case PREDICATE_TUPLE_NOT_IN_SUBQUERY:
		s += "PREDICATE_TUPLE_NOT_IN_SUBQUERY\n"
		s += tabs
		s += "AttriNameOptionTableNameList:\n"
		for _, v := range predicate.AttriNameOptionTableNameList {
			s += AttriNameOptionTableNameToString(v, tabs+"\t")
		}
		s += tabs
		s += "Query:\n"
		s += QueryToString(predicate.Query, tabs+"\t")
	case PREDICATE_TUPLE_IN_TABLE:
		s += "PREDICATE_TUPLE_IN_TABLE\n"
		s += tabs
		s += "AttriNameOptionTableNameList:\n"
		for _, v := range predicate.AttriNameOptionTableNameList {
			s += AttriNameOptionTableNameToString(v, tabs+"\t")
		}
		s += tabs
		s += "TableName: "
		s += predicate.TableName + "\n"
	case PREDICATE_TUPLE_NOT_IN_TABLE:
		s += "PREDICATE_TUPLE_NOT_IN_TABLE\n"
		s += tabs
		s += "AttriNameOptionTableNameList:\n"
		for _, v := range predicate.AttriNameOptionTableNameList {
			s += AttriNameOptionTableNameToString(v, tabs+"\t")
		}
		s += tabs
		s += "TableName: "
		s += predicate.TableName + "\n"
	case PREDICATE_SUBQUERY_EXISTS:
		s += "PREDICATE_SUBQUERY_EXISTS\n"
		s += tabs
		s += "Query:\n"
		s += QueryToString(predicate.Query, tabs+"\t")
	case PREDICATE_SUBQUERY_NOT_EXISTS:
		s += "PREDICATE_SUBQUERY_NOT_EXISTS\n"
		s += tabs
		s += "Query:\n"
		s += QueryToString(predicate.Query, tabs+"\t")
	}

	return s
}

func AttriNameOptionTableNameToString(attriNameOpsTableName *AttriNameOptionTableNameNode, tabs string) string {
	s := tabs
	s += "AttriNameOptionTableNameNode\n"
	if attriNameOpsTableName.TableNameValid {
		s += tabs
		s += "TableName: "
		s += attriNameOpsTableName.TableName + "\n"
	}
	s += tabs
	s += "AttributeName: "
	s += attriNameOpsTableName.AttributeName + "\n"
	return s
}

func CompareMarkToString(compareMark CompareMarkEnum) string {
	switch compareMark {
	case COMPAREMARK_EQUAL:
		return "COMPAREMARK_EQUAL"
	case COMPAREMARK_NOTEQUAL:
		return "COMPAREMARK_NOTEQUAL"
	case COMPAREMARK_LESS:
		return "COMPAREMARK_LESS"
	case COMPAREMARK_GREATER:
		return "COMPAREMARK_GREATER"
	case COMPAREMARK_LESSEQUAL:
		return "COMPAREMARK_LESSEQUAL"
	case COMPAREMARK_GREATEREQUAL:
		return "COMPAREMARK_GREATEREQUAL"
	}
	return ""
}

func QueryToString(query *QueryNode, tabs string) string {
	s := tabs + "QueryNode\n"
	if query.StarValid {
		s += tabs + "StarValid\n"
	} else {
		if query.DistinctValid {
			s += tabs + "DistinctValid\n"
		}
		s += tabs + "SelectList:\n"
		for _, v := range query.SelectList {
			s += SelectListEntryToString(v, tabs+"\t")
		}
	}

	if query.FromListValid {
		s += tabs + "FromList:\n"
		for _, v := range query.FromList {
			s += FromListEntryToString(v, tabs+"\t")
		}
	} else {
		s += tabs + "Join:\n"
		s += JoinToString(query.Join, tabs+"\t")
	}

	if query.WhereValid {
		s += tabs + "WhereCondition:\n"
		s += ConditionToString(query.WhereCondition, tabs+"\t")
	}

	if query.GroupByValid {
		s += tabs + "GroupByList:\n"
		for _, v := range query.GroupByList {
			s += AttriNameOptionTableNameToString(v, tabs+"\t")
		}
	}

	if query.HavingValid {
		s += tabs + "HavingCondition:\n"
		s += ConditionToString(query.HavingCondition, tabs+"\t")
	}

	if query.OrderByValid {
		s += tabs + "OrderByList\n"
		for _, v := range query.OrderByList {
			s += OrderByListEntryToString(v, tabs+"\t")
		}
	}

	if query.LimitValid {
		s += tabs + "InitialPos: " + strconv.Itoa(query.InitialPos) + "\n"
		s += tabs + "OffsetPos: " + strconv.Itoa(query.OffsetPos) + "\n"
	}

	return s
}

func OrderByListEntryToString(entry *OrderByListEntryNode, tabs string) string {
	s := tabs + "OrderByListEntryNode\n"
	switch entry.Type {
	case ORDER_BY_LIST_ENTRY_EXPRESSION:
		s += "Type: ORDER_BY_LIST_ENTRY_EXPRESSION\n"
		s += "Expression:\n"
		s += ExpressionToString(entry.Expression, tabs+"\t")
	case ORDER_BY_LIST_ENTRY_ATTRIBUTE:
		s += "Type: ORDER_BY_LIST_ENTRY_ATTRIBUTE\n"
		s += AttriNameOptionTableNameToString(entry.AttriNameOptionTableName, tabs+"\t")
	}

	switch entry.Trend {
	case ORDER_BY_LIST_ENTRY_ASC:
		s += tabs + "Trend: ORDER_BY_LIST_ENTRY_ASC\n"
	case ORDER_BY_LIST_ENTRY_DESC:
		s += tabs + "Trend: ORDER_BY_LIST_ENTRY_DESC\n"
	}

	return s
}

func JoinToString(join *JoinNode, tabs string) string {
	s := tabs + "JoinNode\n"
	s += tabs + "JoinTableNameL: " + join.JoinTableNameL + "\n"
	s += tabs + "JoinTableNameR: " + join.JoinTableNameR + "\n"

	switch join.Type {
	case CROSS_JOIN:
		s += tabs + "Type: CROSS_JOIN\n"
	case JOIN_ON:
		s += tabs + "Type: JOIN_ON\n"
		s += tabs + "OnList:\n"
		for _, v := range join.OnList {
			s += OnListEntryToString(v, tabs+"\t")
		}

	case NATURAL_JOIN:
		s += tabs + "Type: NATURAL_JOIN\n"
	case NATURAL_FULL_OUTER_JOIN:
		s += tabs + "Type: NATURAL_FULL_OUTER_JOIN\n"
	case NATURAL_LEFT_OUTER_JOIN:
		s += tabs + "Type: NATURAL_LEFT_OUTER_JOIN\n"
	case NATURAL_RIGHT_OUTER_JOIN:
		s += tabs + "Type: NATURAL_RIGHT_OUTER_JOIN\n"
	case FULL_OUTER_JOIN_ON:
		s += tabs + "Type: FULL_OUTER_JOIN_ON\n"
		s += tabs + "OnList:\n"
		for _, v := range join.OnList {
			s += OnListEntryToString(v, tabs+"\t")
		}

	case LEFT_OUTER_JOIN_ON:
		s += tabs + "Type: LEFT_OUTER_JOIN_ON\n"
		s += tabs + "OnList:\n"
		for _, v := range join.OnList {
			s += OnListEntryToString(v, tabs+"\t")
		}

	case RIGHT_OUTER_JOIN_ON:
		s += tabs + "Type: RIGHT_OUTER_JOIN_ON\n"
		s += tabs + "OnList:\n"
		for _, v := range join.OnList {
			s += OnListEntryToString(v, tabs+"\t")
		}

	}

	return s
}

func OnListEntryToString(entry *OnListEntryNode, tabs string) string {
	s := tabs + "OnListEntryNode\n"

	s += tabs + "AttriNameWithTableNameL:\n"
	s += AttriNameOptionTableNameToString(entry.AttriNameWithTableNameL, tabs+"\t")

	s += tabs + "AttriNameWithTableNameR:\n"
	s += AttriNameOptionTableNameToString(entry.AttriNameWithTableNameR, tabs+"\t")

	return s
}

func FromListEntryToString(entry *FromListEntryNode, tabs string) string {
	s := tabs + "FromListEntryNode\n"
	switch entry.Type {
	case FROM_LIST_ENTRY_SUBQUERY:
		s += tabs + "Type: FROM_LIST_ENTRY_SUBQUERY\n"
		s += tabs + "Query:\n"
		s += QueryToString(entry.Query, tabs+"\t")
	case FROM_LIST_ENTRY_TABLE:
		s += tabs + "Type: FROM_LIST_ENTRY_TABLE\n"
		s += tabs + "TableName: " + entry.TableName + "\t"
	}

	if entry.AliasValid {
		s += tabs + "Alias: " + entry.Alias
	}
	return s
}

func SelectListEntryToString(entry *SelectListEntryNode, tabs string) string {
	s := tabs + "SelectListEntryNode\n"
	switch entry.Type {
	case SELECT_LIST_ENTRY_ATTRIBUTE_NAME:
		s += tabs + "Type: SELECT_LIST_ENTRY_ATTRIBUTE_NAME\n"
		s += tabs + "AttriNameOptionTableName:\n"
		s += AttriNameOptionTableNameToString(entry.AttriNameOptionTableName, tabs+"\t")
	case SELECT_LIST_ENTRY_AGGREGATION:
		s += tabs + "Type: SELECT_LIST_ENTRY_AGGREGATION\n"
		s += tabs + "Aggregation:\n"
		s += AggregationToString(entry.Aggregation, tabs+"\t")
	case SELECT_LIST_ENTRY_EXPRESSION:
		s += tabs + "Type: SELECT_LIST_ENTRY_EXPRESSION\n"
		s += tabs + "Expression:\n"
		s += ExpressionToString(entry.Expression, tabs+"\t")
	}

	if entry.AliasValid {
		s += tabs + "Alias: " + entry.Alias + "\n"
	}

	return s
}

func TriggerOldNewEntryToString(entry *TriggerOldNewEntryNode, tabs string) string {
	s := tabs + "TriggerOldNewEntryNode\n"
	switch entry.Type {
	case OLD_ROW_AS:
		s += tabs + "Type: OLD_ROW_AS\n"
	case NEW_ROW_AS:
		s += tabs + "Type: NEW_ROW_AS\n"
	case OLD_TABLE_AS:
		s += tabs + "Type: OLD_TABLE_AS\n"
	case NEW_TABLE_AS:
		s += tabs + "Type: NEW_TABLE_AS\n"
	}
	s += tabs + "Name: " + entry.Name + "\n"
	return s
}
