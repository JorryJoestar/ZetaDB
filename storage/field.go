package storage

//CHAR
type charField struct {
	data rune
}

//VARCHAR(n)
type varcharField struct {
	n    uint32
	data []rune
}

//BIT(n) & BITVARYING(n)
//little endian
type bitField struct {
	n    uint32
	data []byte
}

//BOOLEAN
type booleanField struct {
	data bool
}

//INT & INTEGER
type intField struct {
	data int32
}

//SHORTINT
type shortintField struct {
	data int16
}

//FLOAT
type floatField struct {
	data float64
}

//REAL & DOUBLEPRECISION
type doubleField struct {
	
}

//DECIMAL(n,d) & NUMERIC(n,d)
type decimalField struct {
}

//DATE
type dateField struct {
}

//TIME
type timeField struct {
}
