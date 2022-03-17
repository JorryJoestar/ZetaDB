package container

type ConstraintEnum uint8

const (
	CONSTRAINT_UNIQUE      ConstraintEnum = 1
	CONSTRAINT_PRIMARY_KEY ConstraintEnum = 2
	CONSTRAINT_FOREIGN_KEY ConstraintEnum = 3
	CONSTRAINT_NOT_NULL    ConstraintEnum = 4
	CONSTRAINT_DEFAULT     ConstraintEnum = 5
	CONSTRAINT_CHECK       ConstraintEnum = 6
)

type Constraint struct {
}
