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
	DATE            domainType = 14 //YYYY-MM-DD
	TIME            domainType = 15 //hh:mm:ss.nn
)

type domain struct {
	domainName string
	domainType domainType
	n          uint8
	d          uint8
}

//size of VARCHAR and BITVARYING is not fixed
func (d *domain) DomainSizeUnfixed() bool {
	if d.domainType == VARCHAR || d.domainType == BITVARYING {
		return true
	}
	return false
}

//return size in byte of this domain, if not fixed return -1
func (d *domain) DomainSizeInBytes() int {
	switch d.domainType {
	case CHAR:
		return 1
	case VARCHAR:
		return -1
	case BIT:
		if int(d.n)%8 == 0 {
			return int(d.n) / 8
		} else {
			return int(d.n)/8 + 1
		}
	case BITVARYING:
		return -1
	case BOOLEAN:
		return 1
	case INT:
		return 4
	case INTEGER:
		return 4
	case SHORTINT:
		return 2
	case FLOAT:
		return 4
	case REAL:
		return 4
	case DOUBLEPRECISION:
		return 8
	case DECIMAL:
		if int(d.n)%2 == 0 {
			return int(d.n) / 2
		} else {
			return int(d.n)/2 + 1
		}
	case NUMERIC:
		if int(d.n)%2 == 0 {
			return int(d.n) / 2
		} else {
			return int(d.n)/2 + 1
		}
	case DATE:
		return 8
	case TIME:
		return 8
	}
	return -1
}
