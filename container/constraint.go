package container

type Constraint interface {
}

//if value is not assigned for an attribute, the default value would be assigned to
type DefaultConstraint struct {
	
}

type UniqueConstraint struct {
}

type PrimarykeyConstraint struct {
}

type NotNullConstraint struct {
}

type ForeignkeyConstraint struct {
}

type CheckConstraint struct {
}
