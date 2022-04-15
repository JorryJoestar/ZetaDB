package assister

import (
	"ZetaDB/container"
	"ZetaDB/utility"
	"sync"
)

type Predicator struct {
}

//for singleton pattern
var redicatorInstance *Predicator
var redicatorOnce sync.Once

//to get Predicator, call this function
func GetPredicator() *Predicator {
	redicatorOnce.Do(func() {
		redicatorInstance = &Predicator{}
	})
	return redicatorInstance
}

//TODO unfinished
//see container/predicate.go for PredicateType, CompareMark
func (predicator *Predicator) CheckPredicate(predicate *container.Predicate, tuple *container.Tuple) bool {
	switch predicate.PredicateType {
	case 1: //CompareValue
		fieldBytes, _ := tuple.TupleGetFieldValue(predicate.LeftAttributeIndex)
		switch predicate.CompareValueType {
		case 1: // CompareIntValue
			tupleInt32, _ := utility.BytesToINT(fieldBytes)
			return predicator.CompareInt(int(tupleInt32), predicate.CompareIntValue, predicate.CompareMark)
		case 2: // CompareFloatValue
			tupleFloat32, _ := utility.BytesToFLOAT(fieldBytes)
			return predicator.CompareFloat(float64(tupleFloat32), predicate.CompareFloatValue, predicate.CompareMark)
		case 3: // CompareStringValue
			tupleString, _ := utility.BytesToVARCHAR(fieldBytes)
			return predicator.CompareString(tupleString, predicate.CompareStringValue, predicate.CompareMark)
		}
	case 2: //LikeString
	case 3: //InRelation
	case 4: //NotInRelation
	case 5: //CompareAllRelation
	case 6: //CompareNotAllRelation
	case 7: //CompareAnyRelation
	case 8: //CompareNotAnyRelation
	case 9: //IsNull
	case 10: //IsNotNull
	case 11: //AttriListInRelation
	case 12: //AttriListNotInRelation
	case 13: //Exists
	case 14: //NotExists
	}
	return false
}

//check if a tuple meets a condition
func (predicator *Predicator) CheckCondition(condition *container.Condition, tuple *container.Tuple) bool {
	switch condition.ConditionType {
	case container.CONDITION_PREDICATE:
		return predicator.CheckPredicate(condition.Predicate, tuple)
	case container.CONDITION_AND:
		boolLeft := predicator.CheckCondition(condition.ConditionL, tuple)
		boolRight := predicator.CheckCondition(condition.ConditionR, tuple)
		if boolLeft && boolRight {
			return true
		} else {
			return false
		}
	case container.CONDITION_OR:
		boolLeft := predicator.CheckCondition(condition.ConditionL, tuple)
		boolRight := predicator.CheckCondition(condition.ConditionR, tuple)
		if (!boolLeft) && (!boolRight) {
			return false
		} else {
			return true
		}
	}

	return false
}

func (predicator *Predicator) CompareInt(tupleInt int, compareInt int, CompareMark uint8) bool {
	switch CompareMark {
	case 1: // =
		if tupleInt == compareInt {
			return true
		}
	case 2: // <>
		if tupleInt != compareInt {
			return true
		}
	case 3: // <
		if tupleInt < compareInt {
			return true
		}
	case 4: // >
		if tupleInt > compareInt {
			return true
		}
	case 5: // <=
		if tupleInt <= compareInt {
			return true
		}
	case 6: // >=
		if tupleInt >= compareInt {
			return true
		}
	}
	return false
}

func (predicator *Predicator) CompareFloat(tupleFloat float64, compareFloat float64, CompareMark uint8) bool {
	switch CompareMark {
	case 1: // =
		if tupleFloat == compareFloat {
			return true
		}
	case 2: // <>
		if tupleFloat != compareFloat {
			return true
		}
	case 3: // <
		if tupleFloat < compareFloat {
			return true
		}
	case 4: // >
		if tupleFloat > compareFloat {
			return true
		}
	case 5: // <=
		if tupleFloat <= compareFloat {
			return true
		}
	case 6: // >=
		if tupleFloat >= compareFloat {
			return true
		}
	}
	return false
}

func (predicator *Predicator) CompareString(tupleString string, compareString string, CompareMark uint8) bool {
	switch CompareMark {
	case 1: // =
		if tupleString == compareString {
			return true
		}
	case 2: // <>
		if tupleString != compareString {
			return true
		}
	case 3: // <
		if tupleString < compareString {
			return true
		}
	case 4: // >
		if tupleString > compareString {
			return true
		}
	case 5: // <=
		if tupleString <= compareString {
			return true
		}
	case 6: // >=
		if tupleString >= compareString {
			return true
		}
	}
	return false
}
