package main

import (
	"ZetaDB/container"
	. "ZetaDB/utility"
	"strconv"
)

type Displayer struct {
}

func (displayer *Displayer) TableToString(schema *container.Schema, tuples []*container.Tuple) string {
	tableString := ""

	var maxLengths []int
	var domainTypes []container.DomainType
	var domainNames []string
	var tupleStrings [][]string
	var domains []*container.Domain

	for i, d := range schema.GetSchemaDomains() {
		domainName := d.GetDomainName()
		domainType := d.GetDomainType()

		//update maxLengths
		maxLengths[i] = len(domainName)

		//update domainNames
		domainNames[i] = domainName

		//update dataTypes
		domainTypes[i] = domainType

		//domains
		domains[i] = d

	}

	for tupleIndex, tuple := range tuples {
		for fieldIndex, field := range tuple.TupleGetFields() {

			fieldBytes, nullErr := field.FieldToBytes()
			var fieldString string

			switch domainTypes[fieldIndex] {
			case container.CHAR:

				if nullErr == nil { //not null
					fieldString, _ = BytesToCHAR(fieldBytes)
				} else { //null
					fieldString = "NULL"
				}

			case container.VARCHAR:

				if nullErr == nil { //not null
					fieldString, _ = BytesToVARCHAR(fieldBytes)
				} else { //null
					fieldString = "NULL"
				}

			case container.BIT:

				if nullErr == nil { //not null
					fieldString = "bit_type_value"
				} else { //null
					fieldString = "NULL"
				}

			case container.BITVARYING:

				if nullErr == nil { //not null
					fieldString = "bitvarying_type_value"
				} else { //null
					fieldString = "NULL"
				}

			case container.BOOLEAN:

				if nullErr == nil { //not null
					boolValue := ByteToBOOLEAN(fieldBytes[0])
					if boolValue {
						fieldString = "true"
					} else {
						fieldString = "false"
					}
				} else { //null
					fieldString = "NULL"
				}

			case container.INT:

				if nullErr == nil { //not null
					intValue, _ := BytesToINT(fieldBytes)
					fieldString = string(intValue)
				} else { //null
					fieldString = "NULL"
				}

			case container.INTEGER:

				if nullErr == nil { //not null
					intValue, _ := BytesToInteger(fieldBytes)
					fieldString = string(intValue)
				} else { //null
					fieldString = "NULL"
				}

			case container.SHORTINT:

				if nullErr == nil { //not null
					int16Value, _ := BytesToSHORTINT(fieldBytes)
					intValue := int(int16Value)
					fieldString = string(intValue)
				} else { //null
					fieldString = "NULL"
				}

			case container.FLOAT:

				if nullErr == nil { //not null
					floatValue, _ := BytesToFLOAT(fieldBytes)
					float64Value := float64(floatValue)
					fieldString = strconv.FormatFloat(float64Value, 'E', -1, 32)
				} else { //null
					fieldString = "NULL"
				}

			case container.REAL:

				if nullErr == nil { //not null
					floatValue, _ := BytesToREAL(fieldBytes)
					float64Value := float64(floatValue)
					fieldString = strconv.FormatFloat(float64Value, 'E', -1, 32)
				} else { //null
					fieldString = "NULL"
				}

			case container.DOUBLEPRECISION:

				if nullErr == nil { //not null
					floatValue, _ := BytesToDOUBLEPRECISION(fieldBytes)
					fieldString = strconv.FormatFloat(floatValue, 'E', -1, 64)
				} else { //null
					fieldString = "NULL"
				}

			case container.DECIMAL:

				if nullErr == nil { //not null
					d, _ := domains[fieldIndex].GetDomainD()
					n, _ := domains[fieldIndex].GetDomainN()
					floatValue, _ := BytesToDECIMAL(fieldBytes, int(n), int(d))
					fieldString = strconv.FormatFloat(floatValue, 'E', -1, 64)
				} else { //null
					fieldString = "NULL"
				}

			case container.NUMERIC:

				if nullErr == nil { //not null
					d, _ := domains[fieldIndex].GetDomainD()
					n, _ := domains[fieldIndex].GetDomainN()
					floatValue, _ := BytesToNUMERIC(fieldBytes, int(n), int(d))
					fieldString = strconv.FormatFloat(floatValue, 'E', -1, 64)
				} else { //null
					fieldString = "NULL"
				}

			case container.DATE:

				if nullErr == nil { //not null
					fieldString, _ = BytesToDATE(fieldBytes)
				} else { //null
					fieldString = "NULL"
				}

			case container.TIME:

				if nullErr == nil { //not null
					fieldString, _ = BytesToTIME(fieldBytes)
				} else { //null
					fieldString = "NULL"
				}

			}

			if maxLengths[fieldIndex] < len(fieldString) {
				maxLengths[fieldIndex] = len(fieldString)
			}
			tupleStrings[tupleIndex][fieldIndex] = fieldString

		}
	}

	//generate frameLine
	frameLine := "+"
	for _, maxLen := range maxLengths {
		for i := 0; i < maxLen+2; i++ {
			frameLine += "-"
		}
		frameLine += "+"
	}

	//generate schemaLine
	schemaLine := "|"
	for i, domainName := range domainNames {
		schemaLine += " "
		schemaLine += domainName
		remainSpaces := maxLengths[i] - len(domainName)
		for i := 0; i < remainSpaces+1; i++ {
			schemaLine += " "
		}
		schemaLine += "|"
	}

	//add schemaLine into tableString
	tableString += frameLine
	tableString += schemaLine
	tableString += frameLine

	//loop and insert tupleLine
	for _, tuple := range tupleStrings {
		tupleString := "| "
		for i, field := range tuple {
			tupleString += field
			remainSpaces := maxLengths[i] - len(tupleString)
			for i := 0; i < remainSpaces+1; i++ {
				tupleString += " "
			}
			tupleString += "|"
		}
		tableString += tupleString
	}

	//add last frameLine
	tableString += frameLine

	return tableString
}
