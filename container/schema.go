package storage

type Schema struct {
	tableName   string
	domains     []domain
	constraints []constraint
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

//domains getter
func (s *Schema) GetDomains() []domain {
	return s.domains
}

//return number of domains
func (s *Schema) GetDomainNum() int {
	return len(s.domains)
}
