package container

type Constraint struct {
	/*
		CONSTRAINT_UNIQUE      ConstraintType = 1
		CONSTRAINT_PRIMARY_KEY ConstraintType = 2
		CONSTRAINT_FOREIGN_KEY ConstraintType = 3
		CONSTRAINT_NOT_NULL    ConstraintType = 4
		CONSTRAINT_DEFAULT     ConstraintType = 5
		CONSTRAINT_CHECK       ConstraintType = 6
	*/
	ConstraintType uint8

	ConstraintNameValid bool
	ConstraintName      string

	AttriNameList []string //CONSTRAINT_UNIQUE,CONSTRAINT_PRIMARY_KEY

	/*
		DefaultIntValue     DefaultValueType = 1
		DefaultFloatValue   DefaultValueType = 2
		DefaultStringValue  DefaultValueType = 3
		DefaultBooleanValue DefaultValueType = 4
	*/
	DefaultValueType    uint8
	DefaultIntValue     int
	DefaultFloatValue   float64
	DefaultStringValue  string
	DefaultBooleanValue bool

	Condition *Condition //CONSTRAINT_CHECK

	AttributeNameLocal string //CONSTRAINT_FOREIGN_KEY, CONSTRAINT_NOT_NULL,CONSTRAINT_DEFAULT

	AttributeNameForeign string //CONSTRAINT_FOREIGN_KEY
	ForeignTableName     string //CONSTRAINT_FOREIGN_KEY

	/*
		CONSTRAINT_NOT_DEFERRABLE      Deferrable = 1
		CONSTRAINT_INITIALLY_DEFERRED  Deferrable = 2
		CONSTRAINT_INITIALLY_IMMEDIATE Deferrable = 3
	*/
	DeferrableValid bool  //CONSTRAINT_FOREIGN_KEY
	Deferrable      uint8 //CONSTRAINT_FOREIGN_KEY

	/*
		CONSTRAINT_UPDATE_SET_NULL    UpdateSetValid = 1
		CONSTRAINT_UPDATE_SET_CASCADE UpdateSetValid = 2
	*/
	UpdateSetValid bool  //CONSTRAINT_FOREIGN_KEY
	UpdateSet      uint8 //CONSTRAINT_FOREIGN_KEY

	/*
		CONSTRAINT_DELETE_SET_NULL    DeleteSet = 1
		CONSTRAINT_DELETE_SET_CASCADE DeleteSet = 2
	*/
	DeleteSetValid bool  //CONSTRAINT_FOREIGN_KEY
	DeleteSet      uint8 //CONSTRAINT_FOREIGN_KEY
}
