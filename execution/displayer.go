package execution

import (
	"ZetaDB/container"
	its "ZetaDB/physicalPlan"
	. "ZetaDB/utility"
	"fmt"
	"strconv"
	"sync"
)

type Displayer struct {
}

//for singleton pattern
var displayerInstance *Displayer
var displayerOnce sync.Once

//to get Predicator, call this function
func GetDisplayer() *Displayer {
	displayerOnce.Do(func() {
		displayerInstance = &Displayer{}
	})
	return displayerInstance
}

func (displayer *Displayer) TableToString(schema *container.Schema, tuples []*container.Tuple) string {
	tableString := ""

	fieldsNum := schema.GetSchemaDomainNum()
	tupleNum := len(tuples)

	var maxLengths []int = make([]int, fieldsNum)
	var domainTypes []container.DomainType = make([]container.DomainType, fieldsNum)
	var domainNames []string = make([]string, fieldsNum)
	var tupleStrings [][]string = make([][]string, tupleNum)
	var domains []*container.Domain = make([]*container.Domain, fieldsNum)

	for i := 0; i < tupleNum; i++ {
		tupleStrings[i] = make([]string, fieldsNum)
	}

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
					fieldString = BytesToHexString(fieldBytes)
				} else { //null
					fieldString = "NULL"
				}

			case container.BITVARYING:

				if nullErr == nil { //not null
					fieldString = BytesToHexString(fieldBytes)
				} else { //null
					fieldString = "NULL"
				}

			case container.BOOLEAN:

				if nullErr == nil { //not null
					boolValue := ByteToBOOLEAN(fieldBytes[0])
					if boolValue {
						fieldString = "TRUE"
					} else {
						fieldString = "FALSE"
					}
				} else { //null
					fieldString = "NULL"
				}

			case container.INT:

				if nullErr == nil { //not null
					intValue, _ := BytesToINT(fieldBytes)

					fieldString = strconv.Itoa(int(intValue))

				} else { //null
					fieldString = "NULL"
				}

			case container.INTEGER:

				if nullErr == nil { //not null
					intValue, _ := BytesToInteger(fieldBytes)
					fieldString = strconv.Itoa(int(intValue))
				} else { //null
					fieldString = "NULL"
				}

			case container.SHORTINT:

				if nullErr == nil { //not null
					int16Value, _ := BytesToSHORTINT(fieldBytes)
					fieldString = strconv.Itoa(int(int16Value))
				} else { //null
					fieldString = "NULL"
				}

			case container.FLOAT:

				if nullErr == nil { //not null
					floatValue, _ := BytesToFLOAT(fieldBytes)
					float64Value := float64(floatValue)
					fieldString = strconv.FormatFloat(float64Value, 'E', -1, 64)
				} else { //null
					fieldString = "NULL"
				}

			case container.REAL:

				if nullErr == nil { //not null
					floatValue, _ := BytesToREAL(fieldBytes)
					float64Value := float64(floatValue)
					fieldString = strconv.FormatFloat(float64Value, 'E', -1, 64)
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
	tableString += frameLine + "\n"
	tableString += schemaLine + "\n"
	tableString += frameLine + "\n"

	//loop and insert tupleLine
	for _, tuple := range tupleStrings {
		tupleString := "|"
		for i, field := range tuple {
			tupleString += " "
			tupleString += field
			remainSpaces := maxLengths[i] - len(field)
			for i := 0; i < remainSpaces+1; i++ {
				tupleString += " "
			}
			tupleString += "|"
		}

		tableString += tupleString + "\n"
	}

	//insert emptyLine if result is empty
	if tupleNum == 0 {
		emptyLine := "|"
		for i := 0; i < fieldsNum; i++ {
			for j := 0; j < maxLengths[i]+2; j++ {
				emptyLine += " "
			}
			emptyLine += "|"
		}

		tableString += emptyLine + "\n"
	}

	//add last frameLine
	tableString += frameLine
	return tableString
}

func (displayer *Displayer) PrintTable(tableId uint32) {

	ktm := GetKeytableManager()
	schema, _ := ktm.Query_k_tableId_schema_FromTableId(tableId)

	headPageId, _, _, _, _ := ktm.Query_k_table(tableId)

	seqIt := its.NewSequentialFileReaderIterator(headPageId, schema)
	seqIt.Open(nil, nil)

	var tuples []*container.Tuple

	for seqIt.HasNext() {
		tuple, _ := seqIt.GetNext()
		tuples = append(tuples, tuple)
	}

	result := displayer.TableToString(schema, tuples)
	fmt.Println(result)
}

func (displayer *Displayer) PrintTableByName(tableName string) {
	ktm := GetKeytableManager()
	tableId, _, _ := ktm.Query_k_tableId_schema_FromTableName(tableName)
	displayer.PrintTable(tableId)
}
