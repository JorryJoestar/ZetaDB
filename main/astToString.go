package main

import (
	"ZetaDB/parser"
	"strconv"
)

func ASTToString(ast *parser.ASTNode) string {
	if ast == nil {
		return "ast nil"
	}

	s := "ASTNode\n"
	tabs := ""

	switch ast.Type {
	case parser.AST_DDL:
		s += "Type: AST_DDL\n"
		s += DDLToString(ast.Ddl, tabs+"\t")
	case parser.AST_DML:
		s += "Type: AST_DML\n"
		s += DMLToString(ast.Dml, tabs+"\t")
	case parser.AST_DCL:
		s += "Type: AST_DCL\n"
		s += DCLToString(ast.Dcl, tabs+"\t")
	case parser.AST_DQL:
		s += "Type: AST_DQL\n"
		s += DQLToString(ast.Dql, tabs+"\t")
	}
	return s
}

func DDLToString(ddl *parser.DDLNode, tabs string) string {
	s := ""

	s += tabs
	s += "DDLNode\n"

	switch ddl.Type {
	case parser.DDL_TABLE_CREATE:
		s += tabs
		s += "Type: DDL_TABLE_CREATE\n"
		s += TableCreateToString(ddl.Table, tabs+"\t")
	case parser.DDL_TABLE_DROP:
		s += tabs
		s += "Type: DDL_TABLE_DROP\n"
		s += TableDropToString(ddl.Table, tabs+"\t")
	case parser.DDL_TABLE_ALTER_ADD:
		s += tabs
		s += "Type: DDL_TABLE_ALTER_ADD\n"
		s += TableAlterAddToString(ddl.Table, tabs+"\t")
	case parser.DDL_TABLE_ALTER_DROP:
		s += tabs
		s += "Type: DDL_TABLE_ALTER_DROP\n"
		s += TableAlterDropToString(ddl.Table, tabs+"\t")
	case parser.DDL_ASSERT_CREATE:
		s += tabs
		s += "Type: DDL_ASSERT_CREATE\n"
		s += AssertCreateToString(ddl.Assert, tabs+"\t")
	case parser.DDL_ASSERT_DROP:
		s += tabs
		s += "Type: DDL_ASSERT_DROP\n"
		s += AssertDropToString(ddl.Assert, tabs+"\t")
	case parser.DDL_VIEW_CREATE:
		s += tabs
		s += "Type: DDL_VIEW_CREATE\n"
		s += ViewCreateToString(ddl.View, tabs+"\t")
	case parser.DDL_VIEW_DROP:
		s += tabs
		s += "Type: DDL_VIEW_DROP\n"
		s += ViewDropToString(ddl.View, tabs+"\t")
	case parser.DDL_INDEX_CREATE:
		s += tabs
		s += "Type: DDL_INDEX_CREATE\n"
		s += IndexCreateToString(ddl.Index, tabs+"\t")
	case parser.DDL_INDEX_DROP:
		s += tabs
		s += "Type: DDL_INDEX_DROP\n"
		s += IndexDropToString(ddl.Index, tabs+"\t")
	case parser.DDL_TRIGGER_CREATE:
		s += tabs
		s += "Type: DDL_TRIGGER_CREATE\n"
		s += TriggerCreateToString(ddl.Trigger, tabs+"\t")
	case parser.DDL_TRIGGER_DROP:
		s += tabs
		s += "Type: DDL_TRIGGER_DROP\n"
		s += TriggerDropToString(ddl.Trigger, tabs+"\t")
	case parser.DDL_PSM_CREATE:
		s += tabs
		s += "Type: DDL_PSM_CREATE\n"
		s += PsmCreateToString(ddl.Psm, tabs+"\t")
	case parser.DDL_PSM_DROP:
		s += tabs
		s += "Type: DDL_PSM_DROP\n"
		s += PsmDropToString(ddl.Psm, tabs+"\t")
	}

	return s
}

func DMLToString(dml *parser.DMLNode, tabs string) string {
	//TODO
	s := tabs + "DMLToString TODO \n"
	return s
}

func DCLToString(dcl *parser.DCLNode, tabs string) string {
	//TODO
	s := ""
	return s
}

func DQLToString(dql *parser.DQLNode, tabs string) string {
	//TODO
	s := ""
	return s
}

func TableCreateToString(table *parser.TableNode, tabs string) string {
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

func TableDropToString(table *parser.TableNode, tabs string) string {
	s := tabs + "TableNode\n"
	s += tabs + "TableName: " + table.TableName + "\n"
	return s
}

func TableAlterAddToString(table *parser.TableNode, tabs string) string {
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

func TableAlterDropToString(table *parser.TableNode, tabs string) string {
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

func AssertCreateToString(assert *parser.AssertNode, tabs string) string {
	s := tabs + "AssertNode\n"
	s += tabs + "AssertName: " + assert.AssertName + "\n"
	s += tabs + "Condition:\n"
	s += ConditionToString(assert.Condition, tabs+"\t")
	return s
}

func AssertDropToString(assert *parser.AssertNode, tabs string) string {
	s := tabs + "AssertNode\n"
	s += tabs + "AssertName: " + assert.AssertName + "\n"
	return s
}

func ViewCreateToString(view *parser.ViewNode, tabs string) string {
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

func ViewDropToString(view *parser.ViewNode, tabs string) string {
	s := tabs + "ViewNode\n"
	s += tabs + "ViewName: " + view.ViewName + "\n"
	return s
}

func IndexCreateToString(index *parser.IndexNode, tabs string) string {
	s := tabs + "IndexNode\n"
	s += tabs + "IndexName: " + index.IndexName + "\n"
	s += tabs + "TableName: " + index.TableName + "\n"
	s += tabs + "AttributeNameList:\n"
	for _, v := range index.AttributeNameList {
		s += tabs + "\t" + v + "\n"
	}
	return s
}

func IndexDropToString(index *parser.IndexNode, tabs string) string {
	s := tabs + "IndexNode\n"
	s += tabs + "IndexName: " + index.IndexName + "\n"
	return s
}

func TriggerCreateToString(trigger *parser.TriggerNode, tabs string) string {
	s := tabs + "TriggerNode\n"
	s += tabs + "BeforeAfterType: "
	switch trigger.BeforeAfterType {
	case parser.BEFORE_UPDATE_OF:
		s += "BEFORE_UPDATE_OF\n"
		s += tabs + "BeforeAfterAttriName: " + trigger.BeforeAfterAttriName + "\n"
	case parser.BEFORE_UPDATE:
		s += "BEFORE_UPDATE\n"
	case parser.AFTER_UPDATE_OF:
		s += "AFTER_UPDATE_OF\n"
		s += tabs + "BeforeAfterAttriName: " + trigger.BeforeAfterAttriName + "\n"
	case parser.AFTER_UPDATE:
		s += "AFTER_UPDATE\n"
	case parser.INSTEAD_UPDATE_OF:
		s += "INSTEAD_UPDATE_OF\n"
		s += tabs + "BeforeAfterAttriName: " + trigger.BeforeAfterAttriName + "\n"
	case parser.INSTEAD_UPDATE:
		s += "INSTEAD_UPDATE\n"
	case parser.BEFORE_INSERT:
		s += "BEFORE_INSERT\n"
	case parser.AFTER_INSERT:
		s += "AFTER_INSERT\n"
	case parser.INSTEAD_INSERT:
		s += "INSTEAD_INSERT\n"
	case parser.BEFORE_DELETE:
		s += "BEFORE_DELETE\n"
	case parser.AFTER_DELETE:
		s += "AFTER_DELETE\n"
	case parser.INSTEAD_DELETE:
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
	case parser.FOR_EACH_ROW:
		s += tabs + "ForEachType: FOR_EACH_ROW\n"
	case parser.FOR_EACH_STATEMENT:
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

func TriggerDropToString(trigger *parser.TriggerNode, tabs string) string {
	s := tabs + "TriggerNode\n"
	s += tabs + "TriggerName: " + trigger.TriggerName + "\n"
	return s
}

func PsmCreateToString(psm *parser.PsmNode, tabs string) string {
	s := tabs + "PsmNode\n"

	switch psm.Type {
	case parser.PSM_FUNCTION:
		s += tabs + "Type: PSM_FUNCTION\n"
	case parser.PSM_PROCEDURE:
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

func PsmParameterEntryToString(entry *parser.PsmParameterEntryNode, tabs string) string {
	s := tabs + "PsmParameterEntryNode\n"

	switch entry.Type {
	case parser.PSM_PARAMETER_IN:
		s += tabs + "Type: PSM_PARAMETER_IN\n"
	case parser.PSM_PARAMETER_OUT:
		s += tabs + "Type: PSM_PARAMETER_OUT\n"
	case parser.PSM_PARAMETER_INOUT:
		s += tabs + "Type: PSM_PARAMETER_INOUT\n"
	}

	s += tabs + "Name: " + entry.Name + "\n"

	s += tabs + "Domain:\n"
	s += DomainToString(entry.Domain, tabs+"\t")

	return s
}

func PsmLocalDeclarationEntryToString(entry *parser.PsmLocalDeclarationEntryNode, tabs string) string {
	s := tabs + "PsmLocalDeclarationEntryNode\n"
	s += tabs + "Name: " + entry.Name + "\n"
	s += tabs + "Domain:\n"
	s += DomainToString(entry.Domain, tabs+"\t")

	return s
}

func PsmExecEntryToString(entry *parser.PsmExecEntryNode, tabs string) string {
	s := tabs + "PsmExecEntryNode\n"
	switch entry.Type {
	case parser.PSM_EXEC_RETURN:
		s += tabs + "Type: PSM_EXEC_RETURN\n"
		s += tabs + "PsmValue:\n"
		s += PsmValueToString(entry.PsmValue, tabs+"\t")
	case parser.PSM_EXEC_SET:
		s += tabs + "Type: PSM_EXEC_SET\n"
		s += tabs + "VariableName: " + entry.VariableName + "\n"
		s += tabs + "PsmValue:\n"
		s += PsmValueToString(entry.PsmValue, tabs+"\t")
	case parser.PSM_EXEC_FOR_LOOP:
		s += tabs + "Type: PSM_EXEC_FOR_LOOP\n"
		s += tabs + "PsmForLoop:\n"
		s += PsmForLoopToString(entry.PsmForLoop, tabs+"\t")
	case parser.PSM_EXEC_BRANCH:
		s += tabs + "Type: PSM_EXEC_BRANCH\n"
		s += tabs + "PsmBranch:\n"
		s += PsmBranchToString(entry.PsmBranch, tabs+"\n")
	case parser.PSM_EXEC_DML:
		s += tabs + "Type: PSM_EXEC_DML\n"
		s += DMLToString(entry.Dml, tabs+"\t")
	}
	return s
}

func PsmBranchToString(branch *parser.PsmBranchNode, tabs string) string {
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

func PsmElseifEntryToString(entry *parser.PsmElseifEntryNode, tabs string) string {
	s := tabs + "PsmElseifEntryNode\n"

	s += tabs + "Condition:\n"
	s += ConditionToString(entry.Condition, tabs+"\t")

	s += tabs + "PsmExecList:\n"
	for _, v := range entry.PsmExecList {
		s += PsmExecEntryToString(v, tabs+"\t")
	}
	return s
}

func PsmForLoopToString(ForLoop *parser.PsmForLoopNode, tabs string) string {
	s := tabs + "PsmForLoopNode\n"
	s += tabs + "LoopName: " + ForLoop.LoopName + "\n"
	s += tabs + "CursorName" + ForLoop.CursorName + "\n"
	s += tabs + "Subquery:\n"
	s += QueryToString(ForLoop.Subquery, tabs+"\t")
	s += tabs + "PsmExecList:\n"
	for _, v := range ForLoop.PsmExecList {
		s += PsmExecEntryToString(v, tabs+"\t")
	}
	return s
}

func PsmValueToString(value *parser.PsmValueNode, tabs string) string {
	s := tabs + "PsmValueNode\n"
	switch value.Type {
	case parser.PSMVALUE_ELEMENTARY_VALUE:
		s += tabs + "Type: PSMVALUE_ELEMENTARY_VALUE\n"
		s += ElementaryValueToString(value.ElementaryValue, tabs+"\t")
	case parser.PSMVALUE_CALL:
		s += tabs + "Type: PSMVALUE_CALL\n"
		s += PsmCallStmtToString(value.PsmCall, tabs+"\t")
	case parser.PSMVALUE_EXPRESSION:
		s += tabs + "Type: PSMVALUE_EXPRESSION\n"
		s += ExpressionToString(value.Expression, tabs+"\t")
	case parser.PSMVALUE_ID:
		s += tabs + "Type: PSMVALUE_ID\n"
		s += tabs + "Id: " + value.Id + "\n"
	}

	return s
}

func ExpressionToString(expression *parser.ExpressionNode, tabs string) string {
	s := tabs + "ExpressionNode\n"

	switch expression.Type {
	case parser.EXPRESSION_OPERATOR_PLUS:
		s += tabs + "Type: EXPRESSION_OPERATOR_PLUS\n"
	case parser.EXPRESSION_OPERATOR_MINUS:
		s += tabs + "Type: EXPRESSION_OPERATOR_MINUS\n"
	case parser.EXPRESSION_OPERATOR_DIVISION:
		s += tabs + "Type: EXPRESSION_OPERATOR_DIVISION\n"
	case parser.EXPRESSION_OPERATOR_MULTIPLY:
		s += tabs + "Type: EXPRESSION_OPERATOR_MULTIPLY\n"
	case parser.EXPRESSION_OPERATOR_CONCATENATION:
		s += tabs + "Type: EXPRESSION_OPERATOR_CONCATENATION\n"
	}

	s += tabs + "ExpressionEntryL:\n"
	s += ExpressionEntryToString(expression.ExpressionEntryL, tabs+"\t")

	s += tabs + "ExpressionEntryR:\n"
	s += ExpressionEntryToString(expression.ExpressionEntryR, tabs+"\t")

	return s
}

func ExpressionEntryToString(entry *parser.ExpressionEntryNode, tabs string) string {
	s := tabs + "ExpressionEntryNode\n"
	switch entry.Type {
	case parser.EXPRESSION_ENTRY_ELEMENTARY_VALUE:
		s += tabs + "Type: EXPRESSION_ENTRY_ELEMENTARY_VALUE\n"
		s += ElementaryValueToString(entry.ElementaryValue, tabs+"\t")
	case parser.EXPRESSION_ENTRY_ATTRIBUTE_NAME:
		s += tabs + "Type: EXPRESSION_ENTRY_ATTRIBUTE_NAME\n"
		s += AttriNameOptionTableNameToString(entry.AttriNameOptionTableName, tabs+"\t")
	case parser.EXPRESSION_ENTRY_AGGREGATION:
		s += tabs + "Type: EXPRESSION_ENTRY_AGGREGATION\n"
		s += AggregationToString(entry.Aggregation, tabs+"\t")
	case parser.EXPRESSION_ENTRY_EXPRESSION:
		s += tabs + "Type: EXPRESSION_ENTRY_EXPRESSION\n"
		s += ExpressionToString(entry.Expression, tabs+"\t")
	}
	return s
}

func AggregationToString(aggregation *parser.AggregationNode, tabs string) string {
	s := tabs + "AggregationNode\n"

	switch aggregation.Type {
	case parser.AGGREGATION_SUM:
		s += tabs + "Type: AGGREGATION_SUM\n"
	case parser.AGGREGATION_AVG:
		s += tabs + "Type: AGGREGATION_AVG\n"
	case parser.AGGREGATION_MIN:
		s += tabs + "Type: AGGREGATION_MIN\n"
	case parser.AGGREGATION_MAX:
		s += tabs + "Type: AGGREGATION_MAX\n"
	case parser.AGGREGATION_COUNT:
		s += tabs + "Type: AGGREGATION_COUNT\n"
	case parser.AGGREGATION_COUNT_ALL:
		s += tabs + "Type: AGGREGATION_COUNT_ALL\n"
	}

	if aggregation.DistinctValid {
		s += tabs + "DistinctValid\n"
	}

	if aggregation.Type != parser.AGGREGATION_COUNT_ALL {
		s += AttriNameOptionTableNameToString(aggregation.AttriNameOptionTableName, tabs+"\t")
	}

	return s

}

func PsmCallStmtToString(psmCall *parser.PsmNode, tabs string) string {
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

func PsmDropToString(psm *parser.PsmNode, tabs string) string {
	s := tabs + "PsmNode\n"

	switch psm.Type {
	case parser.PSM_PROCEDURE:
		s += tabs + "Type: PSM_PROCEDURE\n"
	case parser.PSM_FUNCTION:
		s += tabs + "Type: PSM_FUNCTION\n"
	}

	s += tabs + "PsmName: " + psm.PsmName + "\n"

	return s
}

func DomainToString(domain *parser.DomainNode, tabs string) string {
	s := ""

	s += tabs
	s += "DomainNode\n"

	s += tabs
	s += "Type: "

	switch domain.Type {
	case parser.DOMAIN_CHAR:
		s += "DOMAIN_CHAR\n"
	case parser.DOMAIN_VARCHAR:
		s += "DOMAIN_VARCHAR\n"

		s += tabs
		s += "n: " + strconv.Itoa(domain.N) + "\n"
	case parser.DOMAIN_BIT:
		s += "DOMAIN_BIT\n"

		s += tabs
		s += "n: " + strconv.Itoa(domain.N) + "\n"
	case parser.DOMAIN_BITVARYING:
		s += "DOMAIN_BITVARYING\n"

		s += tabs
		s += "n: " + strconv.Itoa(domain.N) + "\n"
	case parser.DOMAIN_BOOLEAN:
		s += "DOMAIN_BOOLEAN\n"
	case parser.DOMAIN_INT:
		s += "DOMAIN_INT\n"
	case parser.DOMAIN_INTEGER:
		s += "DOMAIN_INTEGER\n"
	case parser.DOMAIN_SHORTINT:
		s += "DOMAIN_SHORTINT\n"
	case parser.DOMAIN_FLOAT:
		s += "DOMAIN_FLOAT\n"
	case parser.DOMAIN_REAL:
		s += "DOMAIN_REAL\n"
	case parser.DOMAIN_DOUBLEPRECISION:
		s += "DOMAIN_DOUBLEPRECISION\n"
	case parser.DOMAIN_DECIMAL:
		s += "DOMAIN_DECIMAL\n"

		s += tabs
		s += "n: " + strconv.Itoa(domain.N) + "\n"

		s += tabs
		s += "d: " + strconv.Itoa(domain.D) + "\n"
	case parser.DOMAIN_NUMERIC:
		s += "DOMAIN_NUMERIC\n"

		s += tabs
		s += "n: " + strconv.Itoa(domain.N) + "\n"

		s += tabs
		s += "d: " + strconv.Itoa(domain.D) + "\n"
	case parser.DOMAIN_DATE:
		s += "DOMAIN_DATE\n"
	case parser.DOMAIN_TIME:
		s += "DOMAIN_TIME\n"
	}

	return s
}

func ConstraintToString(constraint *parser.ConstraintNode, tabs string) string {
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
	case parser.CONSTRAINT_UNIQUE:
		s += "CONSTRAINT_UNIQUE\n"
		s += tabs
		s += "AttriNameList:\n"
		for _, v := range constraint.AttriNameList {
			s += tabs + "\t"
			s += v + "\n"
		}
	case parser.CONSTRAINT_PRIMARY_KEY:
		s += "CONSTRAINT_PRIMARY_KEY\n"
		s += tabs
		s += "AttriNameList:\n"
		for _, v := range constraint.AttriNameList {
			s += tabs + "\t"
			s += v + "\n"
		}
	case parser.CONSTRAINT_FOREIGN_KEY:
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
			case parser.CONSTRAINT_NOT_DEFERRABLE:
				s += "CONSTRAINT_NOT_DEFERRABLE\n"
			case parser.CONSTRAINT_INITIALLY_DEFERRED:
				s += "CONSTRAINT_INITIALLY_DEFERRED\n"
			case parser.CONSTRAINT_INITIALLY_IMMEDIATE:
				s += "CONSTRAINT_INITIALLY_IMMEDIATE\n"
			}
		}
		if constraint.UpdateSetValid {
			s += tabs
			s += "UpdateSet: "
			switch constraint.UpdateSet {
			case parser.CONSTRAINT_UPDATE_SET_NULL:
				s += "CONSTRAINT_UPDATE_SET_NULL\n"
			case parser.CONSTRAINT_UPDATE_SET_CASCADE:
				s += "CONSTRAINT_UPDATE_SET_CASCADE\n"
			}
		}
		if constraint.DeleteSetValid {
			s += tabs
			s += "DeleteSet: "
			switch constraint.DeleteSet {
			case parser.CONSTRAINT_DELETE_SET_NULL:
				s += "CONSTRAINT_DELETE_SET_NULL\n"
			case parser.CONSTRAINT_DELETE_SET_CASCADE:
				s += "CONSTRAINT_DELETE_SET_CASCADE\n"
			}
		}
	case parser.CONSTRAINT_NOT_NULL:
		s += "CONSTRAINT_NOT_NULL\n"
		s += tabs
		s += "AttributeNameLocal: " + constraint.AttributeNameLocal + "\n"
	case parser.CONSTRAINT_DEFAULT:
		s += "CONSTRAINT_DEFAULT\n"
		s += tabs
		s += "AttributeNameLocal: " + constraint.AttributeNameLocal + "\n"
		s += tabs
		s += "ElementaryValue:\n"
		s += ElementaryValueToString(constraint.ElementaryValue, tabs+"\t")
	case parser.CONSTRAINT_CHECK:
		s += "CONSTRAINT_CHECK\n"
		s += ConditionToString(constraint.Condition, tabs+"\t")
	}

	return s
}

func ElementaryValueToString(elementaryValue *parser.ElementaryValueNode, tabs string) string {
	s := ""

	s += tabs
	s += "ElementaryValueNode\n"
	s += tabs
	s += "Type: "

	switch elementaryValue.Type {
	case parser.ELEMENTARY_VALUE_INT:
		s += "ELEMENTARY_VALUE_INT\n"
		s += tabs
		s += "IntValue: "
		s += strconv.Itoa(elementaryValue.IntValue)
		s += "\n"
	case parser.ELEMENTARY_VALUE_FLOAT:
		s += "ELEMENTARY_VALUE_FLOAT\n"
		s += tabs
		s += "FloatValue: "
		s += strconv.FormatFloat(elementaryValue.FloatValue, 'g', -1, 64)
		s += "\n"
	case parser.ELEMENTARY_VALUE_STRING:
		s += "ELEMENTARY_VALUE_STRING\n"
		s += tabs
		s += "StringValue: "
		s += elementaryValue.StringValue
		s += "\n"
	case parser.ELEMENTARY_VALUE_BOOLEAN:
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

func ConditionToString(condition *parser.ConditionNode, tabs string) string {
	s := tabs
	s += "ConditionNode\n"

	s += tabs
	s += "Type: "
	switch condition.Type {
	case parser.CONDITION_PREDICATE:
		s += "CONDITION_PREDICATE\n"
		s += tabs
		s += "Predicate:\n"
		s += PredicateToString(condition.Predicate, tabs+"\t")
	case parser.CONDITION_AND:
		s += "CONDITION_AND\n"
		s += tabs
		s += "ConditionL:\n"
		s += ConditionToString(condition.ConditionL, tabs+"\t")
		s += tabs
		s += "ConditionR:\n"
		s += ConditionToString(condition.ConditionR, tabs+"\t")
	case parser.CONDITION_OR:
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

func PredicateToString(predicate *parser.PredicateNode, tabs string) string {
	s := ""
	s += tabs
	s += "PredicateNode\n"

	s += tabs
	s += "Type: "
	switch predicate.Type {
	case parser.PREDICATE_COMPARE_ELEMENTARY_VALUE:
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
	case parser.PREDICATE_LIKE_STRING_VALUE:
		s += "PREDICATE_LIKE_STRING_VALUE\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "ElementaryValue:\n"
		s += ElementaryValueToString(predicate.ElementaryValue, tabs+"\t")
	case parser.PREDICATE_IN_SUBQUERY:
		s += "PREDICATE_IN_SUBQUERY\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "Subquery:\n"
		s += QueryToString(predicate.Subquery, tabs+"\t")
	case parser.PREDICATE_NOT_IN_SUBQUERY:
		s += "PREDICATE_NOT_IN_SUBQUERY\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "Subquery:\n"
		s += QueryToString(predicate.Subquery, tabs+"\t")
	case parser.PREDICATE_IN_TABLE:
		s += "PREDICATE_IN_TABLE\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "TableName: "
		s += predicate.TableName + "\n"
	case parser.PREDICATE_NOT_IN_TABLE:
		s += "PREDICATE_NOT_IN_TABLE\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "TableName: "
		s += predicate.TableName + "\n"
	case parser.PREDICATE_COMPARE_ALL_SUBQUERY:
		s += "PREDICATE_COMPARE_ALL_SUBQUERY\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "Subquery:\n"
		s += QueryToString(predicate.Subquery, tabs+"\t")
		s += tabs
		s += "CompareMark: "
		s += CompareMarkToString(predicate.CompareMark) + "\n"
	case parser.PREDICATE_COMPARE_NOT_ALL_SUBQUERY:
		s += "PREDICATE_COMPARE_NOT_ALL_SUBQUERY\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "Subquery:\n"
		s += QueryToString(predicate.Subquery, tabs+"\t")
		s += tabs
		s += "CompareMark: "
		s += CompareMarkToString(predicate.CompareMark) + "\n"
	case parser.PREDICATE_COMPARE_ANY_SUBQUERY:
		s += "PREDICATE_COMPARE_ANY_SUBQUERY\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "Subquery:\n"
		s += QueryToString(predicate.Subquery, tabs+"\t")
		s += tabs
		s += "CompareMark: "
		s += CompareMarkToString(predicate.CompareMark) + "\n"
	case parser.PREDICATE_COMPARE_NOT_ANY_SUBQUERY:
		s += "PREDICATE_COMPARE_NOT_ANY_SUBQUERY\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
		s += tabs
		s += "Subquery:\n"
		s += QueryToString(predicate.Subquery, tabs+"\t")
		s += tabs
		s += "CompareMark: "
		s += CompareMarkToString(predicate.CompareMark) + "\n"
	case parser.PREDICATE_COMPARE_ALL_TABLE:
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
	case parser.PREDICATE_COMPARE_NOT_ALL_TABLE:
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
	case parser.PREDICATE_COMPARE_ANY_TABLE:
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
	case parser.PREDICATE_COMPARE_NOT_ANY_TABLE:
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
	case parser.PREDICATE_IS_NULL:
		s += "PREDICATE_IS_NULL\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
	case parser.PREDICATE_IS_NOT_NULL:
		s += "PREDICATE_IS_NOT_NULL\n"
		s += tabs
		s += "AttriNameWithTableNameL:\n"
		s += AttriNameOptionTableNameToString(predicate.AttriNameWithTableNameL, tabs+"\t")
	case parser.PREDICATE_TUPLE_IN_SUBQUERY:
		s += "PREDICATE_TUPLE_IN_SUBQUERY\n"
		s += tabs
		s += "AttriNameOptionTableNameList:\n"
		for _, v := range predicate.AttriNameOptionTableNameList {
			s += AttriNameOptionTableNameToString(v, tabs+"\t")
		}
		s += tabs
		s += "Subquery:\n"
		s += QueryToString(predicate.Subquery, tabs+"\t")
	case parser.PREDICATE_TUPLE_NOT_IN_SUBQUERY:
		s += "PREDICATE_TUPLE_NOT_IN_SUBQUERY\n"
		s += tabs
		s += "AttriNameOptionTableNameList:\n"
		for _, v := range predicate.AttriNameOptionTableNameList {
			s += AttriNameOptionTableNameToString(v, tabs+"\t")
		}
		s += tabs
		s += "Subquery:\n"
		s += QueryToString(predicate.Subquery, tabs+"\t")
	case parser.PREDICATE_TUPLE_IN_TABLE:
		s += "PREDICATE_TUPLE_IN_TABLE\n"
		s += tabs
		s += "AttriNameOptionTableNameList:\n"
		for _, v := range predicate.AttriNameOptionTableNameList {
			s += AttriNameOptionTableNameToString(v, tabs+"\t")
		}
		s += tabs
		s += "TableName: "
		s += predicate.TableName + "\n"
	case parser.PREDICATE_TUPLE_NOT_IN_TABLE:
		s += "PREDICATE_TUPLE_NOT_IN_TABLE\n"
		s += tabs
		s += "AttriNameOptionTableNameList:\n"
		for _, v := range predicate.AttriNameOptionTableNameList {
			s += AttriNameOptionTableNameToString(v, tabs+"\t")
		}
		s += tabs
		s += "TableName: "
		s += predicate.TableName + "\n"
	case parser.PREDICATE_SUBQUERY_EXISTS:
		s += "PREDICATE_SUBQUERY_EXISTS\n"
		s += tabs
		s += "Subquery:\n"
		s += QueryToString(predicate.Subquery, tabs+"\t")
	case parser.PREDICATE_SUBQUERY_NOT_EXISTS:
		s += "PREDICATE_SUBQUERY_NOT_EXISTS\n"
		s += tabs
		s += "Subquery:\n"
		s += QueryToString(predicate.Subquery, tabs+"\t")
	}

	return s
}

func AttriNameOptionTableNameToString(attriNameOpsTableName *parser.AttriNameOptionTableNameNode, tabs string) string {
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

func CompareMarkToString(compareMark parser.CompareMarkEnum) string {
	switch compareMark {
	case parser.COMPAREMARK_EQUAL:
		return "COMPAREMARK_EQUAL"
	case parser.COMPAREMARK_NOTEQUAL:
		return "COMPAREMARK_NOTEQUAL"
	case parser.COMPAREMARK_LESS:
		return "COMPAREMARK_LESS"
	case parser.COMPAREMARK_GREATER:
		return "COMPAREMARK_GREATER"
	case parser.COMPAREMARK_LESSEQUAL:
		return "COMPAREMARK_LESSEQUAL"
	case parser.COMPAREMARK_GREATEREQUAL:
		return "COMPAREMARK_GREATEREQUAL"
	}
	return ""
}

func QueryToString(query *parser.QueryNode, tabs string) string {
	//TODO
	s := tabs
	s += "QueryToString TODO\n"
	return s
}

func TriggerOldNewEntryToString(entry *parser.TriggerOldNewEntryNode, tabs string) string {
	s := tabs + "TriggerOldNewEntryNode\n"
	switch entry.Type {
	case parser.OLD_ROW_AS:
		s += tabs + "Type: OLD_ROW_AS\n"
	case parser.NEW_ROW_AS:
		s += tabs + "Type: NEW_ROW_AS\n"
	case parser.OLD_TABLE_AS:
		s += tabs + "Type: OLD_TABLE_AS\n"
	case parser.NEW_TABLE_AS:
		s += tabs + "Type: NEW_TABLE_AS\n"
	}
	s += tabs + "Name: " + entry.Name + "\n"
	return s
}
