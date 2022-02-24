package storage

type schema struct {
	tableName   string
	domains     []domain
	constraints []constraint //TODO
}

//number of domains whose size is not fixed
func (s *schema) UnfixedDomainNum() int {
	num := 0

	for _, d := range s.domains {
		if d.DomainSizeUnfixed() {
			num++
		}
	}

	return num
}

//domains getter
func (s *schema) GetDomains() []domain {
	return s.domains
}

//return number of domains
func (s *schema) GetDomainNum() int {
	return len(s.domains)
}
