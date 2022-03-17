package container

import "errors"

type DomainType uint8

const (
	CHAR            DomainType = 1
	VARCHAR         DomainType = 2 //VARCHAR(n)
	BIT             DomainType = 3 //BIT(n)
	BITVARYING      DomainType = 4 //BITVARYING(n)
	BOOLEAN         DomainType = 5
	INT             DomainType = 6
	INTEGER         DomainType = 7
	SHORTINT        DomainType = 8
	FLOAT           DomainType = 9
	REAL            DomainType = 10
	DOUBLEPRECISION DomainType = 11
	DECIMAL         DomainType = 12 //DECIMAL(n,d)
	NUMERIC         DomainType = 13 //NUMERIC(n,d)
	DATE            DomainType = 14 //YYYY-MM-DD
	TIME            DomainType = 15 //hh:mm:ss.nn
)

type Domain struct {
	domainName string
	domainType DomainType
	n          int32
	d          int32
}

//if n, d is unnecessary, pass value 0 as parameter
//throw error if n, d are not 0 and are unnecessary
func NewDomain(name string, t DomainType, inN int32, inD int32) (*Domain, error) {

	//if domain type is not DECIMAL or NUMERIC, d is unnecessary
	if t != DECIMAL && t != NUMERIC && inD != 0 {
		return nil, errors.New("d unnecessary")
	}

	//if domain type is not DECIMAL or NUMERIC or VARCHAR or BIT or BITVARYING, n is unnecessary
	if t != DECIMAL && t != NUMERIC && t != VARCHAR && t != BIT && t != BITVARYING && inN != 0 {
		return nil, errors.New("n unnecessary")
	}

	newDomain := &Domain{
		domainName: name,
		domainType: t,
		n:          inN,
		d:          inD}

	return newDomain, nil
}

//domainName getter
func (d *Domain) GetDomainName() string {
	return d.domainName
}

//domainType getter
func (d *Domain) GetDomainType() DomainType {
	return d.domainType
}

//domain n getter
//throw error if n is unnecessary for the current domain
func (d *Domain) GetDomainN() (int32, error) {

	t := d.domainType
	if t != DECIMAL && t != NUMERIC && t != VARCHAR && t != BIT && t != BITVARYING {
		return 0, errors.New("n is unnecessary")
	}

	return d.n, nil
}

//domain d getter
//throw error if d is unnecessary for the current domain
func (d *Domain) GetDomainD() (int32, error) {

	t := d.domainType
	if t != DECIMAL && t != NUMERIC {
		return 0, errors.New("d is unnecessary")
	}

	return d.d, nil
}

//size of VARCHAR and BITVARYING is not fixed
func (d *Domain) DomainSizeUnfixed() bool {
	if d.domainType == VARCHAR || d.domainType == BITVARYING {
		return true
	}
	return false
}

//return size in byte of this domain
//throw error if domain size is unfixed
func (d *Domain) DomainSizeInBytes() (int, error) {

	switch d.domainType {
	case CHAR:
		return 1, nil
	case VARCHAR:
		return 0, errors.New("domain size unfixed")
	case BIT:
		if int(d.n)%8 == 0 {
			return int(d.n) / 8, nil
		} else {
			return int(d.n)/8 + 1, nil
		}
	case BITVARYING:
		return 0, errors.New("domain size unfixed")
	case BOOLEAN:
		return 1, nil
	case INT:
		return 4, nil
	case INTEGER:
		return 4, nil
	case SHORTINT:
		return 2, nil
	case FLOAT:
		return 4, nil
	case REAL:
		return 4, nil
	case DOUBLEPRECISION:
		return 8, nil
	case DECIMAL:
		if int(d.n)%2 == 0 { //add the sign(+/-) byte
			return int(d.n)/2 + 1, nil
		} else {
			return int(d.n)/2 + 2, nil
		}
	case NUMERIC:
		if int(d.n)%2 == 0 {
			return int(d.n)/2 + 1, nil
		} else {
			return int(d.n)/2 + 2, nil
		}
	case DATE:
		return 4, nil
	case TIME:
		return 4, nil
	}
	return 0, errors.New("domain type unknown")
}
