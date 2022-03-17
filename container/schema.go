package container

import "errors"

type Schema struct {
	tableName   string
	domains     []*Domain
	constraints []*Constraint
}

//create a new schema
//throw error if length of domainList is 0, or name is empty string
func NewSchema(name string, domainList []*Domain, constraintList []*Constraint) (*Schema, error) {

	//if length of domainList is 0, throw error
	if len(domainList) == 0 {
		return nil, errors.New("domainList length 0")
	}

	//if name is "" (empty string), throw error
	if name == "" {
		return nil, errors.New("empty string as name")
	}

	newSchema := &Schema{
		tableName:   name,
		domains:     domainList,
		constraints: constraintList}

	return newSchema, nil
}

//number of domains whose size is not fixed
func (s *Schema) UnfixedDomainNum() int {
	num := 0

	for _, d := range s.domains {
		if d.DomainSizeUnfixed() {
			num++
		}
	}

	return num
}

//table name getter
func (s *Schema) GetSchemaTableName() string {
	return s.tableName
}

//domains getter
func (s *Schema) GetSchemaDomains() []*Domain {
	return s.domains
}

//single domain getter according to index
//throw error if index is invalid
func (s *Schema) GetSchemaDomain(index int) (*Domain, error) {

	//throw error if index is invalid
	if index >= s.GetSchemaDomainNum() {
		return nil, errors.New("index invalid")
	}

	return s.domains[index], nil
}

//constraints getter
func (s *Schema) GetSchemaConstraints() []*Constraint {
	return s.constraints
}

//return number of domains
func (s *Schema) GetSchemaDomainNum() int {
	return len(s.domains)
}

//return number of constraints
func (s *Schema) GetSchemaConstraintNum() int {
	return len(s.constraints)
}
