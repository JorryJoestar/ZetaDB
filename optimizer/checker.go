package optimizer

import (
	"ZetaDB/parser"
)

type Checker struct {
}

//throw error if this table name already exists that belongs to current user
func (checker *Checker) Check_TABLE_CREATE() error {
	return nil
}

func (checker *Checker) Check_TABLE_DROP() error {
	return nil
}

func (checker *Checker) Check_TABLE_ALTER_ADD() error {
	return nil
}

func (checker *Checker) Check_TABLE_ALTER_DROP() error {
	return nil
}

func (checker *Checker) Check_ASSERT_CREATE() error {
	return nil
}

func (checker *Checker) Check_ASSERT_DROP() error {
	return nil
}

func (checker *Checker) Check_VIEW_CREATE() error {
	return nil
}

func (checker *Checker) Check_VIEW_DROP() error {
	return nil
}

func (checker *Checker) Check_INDEX_CREATE() error {
	return nil
}

func (checker *Checker) Check_INDEX_DROP() error {
	return nil
}

func (checker *Checker) Check_TRIGGER_CREATE() error {
	return nil
}

func (checker *Checker) Check_TRIGGER_DROP() error {
	return nil
}

func (checker *Checker) Check_PSM_CREATE() error {
	return nil
}

func (checker *Checker) Check_PSM_DROP() error {
	return nil
}

//throw error if tableName invalid, that means no such table belongs to current user
//throw error if some attributes in attriNameList are not belong to this table
//throw error if length of attriNameList is different from length of subQuery/elementaryValueList
//throw error if data type of attriNameList is different from data type of subQuery/elementaryValueList
//throw error if unappeared attribute is not allowed to be null and has no default value
func (checker *Checker) Check_INSERT() error {
	return nil
}

//throw error if tableName invalid, that means no such table belongs to current user
//throw error if attributeNames in condition is not belong to this table
//throw error if condition is invalid
//throw error if data type of changed attribute is different from the data type of given value
func (checker *Checker) Check_UPDATE() error {
	return nil
}

//throw error if tableName invalid, that means no such table belongs to current user
//throw error if attributeNames in condition is not belong to this table
//throw error if condition is invalid
func (checker *Checker) Check_DELETE(deleteNode *parser.DeleteNode) error {
	return nil
}

func (checker *Checker) Check_PSMCALL() error {
	return nil
}

//throw error if relations in fromStmt do not exist
//TODO
func (checker *Checker) Check_DQL() error {
	return nil
}

func (checker *Checker) Check_Predicate() error {
	return nil
}

func (checker *Checker) Check_Condition() error {
	return nil
}
