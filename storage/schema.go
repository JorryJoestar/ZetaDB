package storage

type schema struct {
	tableName   string
	domains     []domain
	constraints []constraint
}

//number of domains whose size is not fixed
func (s *schema) UnfixedDomainNum() int {
	num := 0

	for _, v := range s.domains {
		if v.DomainSizeUnfixed() {
			num++
		}
	}

	return num
}

func (s *schema) GetDomains() []domain {
	return s.domains
}
