package storage

type domainType uint8

const (
	CHAR            domainType = 1
	VARCHAR         domainType = 2 //VARCHAR(n)
	BIT             domainType = 3 //BIT(n)
	BITVARYING      domainType = 4 //BITVARYING(n)
	BOOLEAN         domainType = 5
	INT             domainType = 6
	INTEGER         domainType = 7
	SHORTINT        domainType = 8
	FLOAT           domainType = 9
	REAL            domainType = 10
	DOUBLEPRECISION domainType = 11
	DECIMAL         domainType = 12 //DECIMAL(n,d)
	NUMERIC         domainType = 13 //NUMERIC(n,d)
	DATE            domainType = 14
	TIME            domainType = 15
)

type domain struct {
	domainName string
	domainType domainType
	n          uint8
	d          uint8
}

//return the size of this domain in bytes
func (domain *domain) DomainSizeInBytes() int {
	return 0
}
