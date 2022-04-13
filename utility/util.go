package utility

import (
	"errors"
	"fmt"
	"os"
)

/*
        type           elementType           elementLength
   CHAR                    1                      1
   INT                     2                      4
   INTEGER                 3                      4
   SHORTINT                4                      2
   FLOAT                   5                      4
   REAL                    6                      4
   DOUBLEPRECISION         7                      8
   DATE                    8                      4
   TIME                    9                      4
*/

//compare two element bytes according to their elementType
//if bytes1 < bytes2, return a negative value
//if bytes1 = bytes2, return 0
//if bytes1 > bytes2, return a positive value
//throw error if bytes1 or bytes2 length invalid
func CompareBytesElement(elementType uint32, bytes1 []byte, bytes2 []byte) (int, error) {
	//throw error if bytes1 or bytes2 length invalid
	switch elementType {
	case 1: //length 1
		if len(bytes1) != 1 || len(bytes2) != 1 {
			return 0, errors.New("bytes length invalid")
		}
	case 4: //length 2
		if len(bytes1) != 2 || len(bytes2) != 2 {
			return 0, errors.New("bytes length invalid")
		}
	case 2, 3, 5, 6, 8, 9: //length 4
		if len(bytes1) != 4 || len(bytes2) != 4 {
			return 0, errors.New("bytes length invalid")
		}
	case 7: //length 8
		if len(bytes1) != 8 || len(bytes2) != 8 {
			return 0, errors.New("bytes length invalid")
		}
	}

	switch elementType {
	case 1:
		v1, _ := BytesToCHAR(bytes1)
		v2, _ := BytesToCHAR(bytes2)
		if v1 < v2 {
			return -1, nil
		} else if v1 == v2 {
			return 0, nil
		} else {
			return 1, nil
		}
	case 2:
		v1, _ := BytesToINT(bytes1)
		v2, _ := BytesToINT(bytes2)
		if v1 < v2 {
			return -1, nil
		} else if v1 == v2 {
			return 0, nil
		} else {
			return 1, nil
		}
	case 3:
		v1, _ := BytesToInteger(bytes1)
		v2, _ := BytesToInteger(bytes2)
		if v1 < v2 {
			return -1, nil
		} else if v1 == v2 {
			return 0, nil
		} else {
			return 1, nil
		}
	case 4:
		v1, _ := BytesToSHORTINT(bytes1)
		v2, _ := BytesToSHORTINT(bytes2)
		if v1 < v2 {
			return -1, nil
		} else if v1 == v2 {
			return 0, nil
		} else {
			return 1, nil
		}
	case 5:
		v1, _ := BytesToFLOAT(bytes1)
		v2, _ := BytesToFLOAT(bytes2)
		if v1 < v2 {
			return -1, nil
		} else if v1 == v2 {
			return 0, nil
		} else {
			return 1, nil
		}
	case 6:
		v1, _ := BytesToREAL(bytes1)
		v2, _ := BytesToREAL(bytes2)
		if v1 < v2 {
			return -1, nil
		} else if v1 == v2 {
			return 0, nil
		} else {
			return 1, nil
		}
	case 7:
		v1, _ := BytesToDOUBLEPRECISION(bytes1)
		v2, _ := BytesToDOUBLEPRECISION(bytes2)
		if v1 < v2 {
			return -1, nil
		} else if v1 == v2 {
			return 0, nil
		} else {
			return 1, nil
		}
	case 8:
		v1, _ := BytesToDATE(bytes1)
		v2, _ := BytesToDATE(bytes2)
		if v1 < v2 {
			return -1, nil
		} else if v1 == v2 {
			return 0, nil
		} else {
			return 1, nil
		}
	case 9:
		v1, _ := BytesToTIME(bytes1)
		v2, _ := BytesToTIME(bytes2)
		if v1 < v2 {
			return -1, nil
		} else if v1 == v2 {
			return 0, nil
		} else {
			return 1, nil
		}
	default:
		return 0, errors.New("unmatching error")
	}
}

//insert a byte element into an ordered slice at correct position
//throw error if value is already in this slice
func InsertToOrderedSlice(elementType uint32, slice *[][]byte, element []byte) error {
	for i, v := range *slice {
		compare, comErr := CompareBytesElement(elementType, element, v)
		if comErr != nil {
			return comErr
		}
		if compare == 0 {
			return errors.New("duplicated value")
		} else if compare < 0 { //find insert position
			*slice = append(*slice, nil)
			copy((*slice)[i+1:], (*slice)[i:])
			(*slice)[i] = element
			return nil
		}
	}
	//append at the end
	*slice = append(*slice, element)
	return nil
}

//insert a byte element into an ordered slice at correct position
//delete and return the smallest one
//throw error if value is already in this slice
func InsertToOrderedSliceReturnMin(elementType uint32, slice *[][]byte, element []byte) ([]byte, error) {
	changed := false

	for i, v := range *slice {
		compare, comErr := CompareBytesElement(elementType, element, v)
		if comErr != nil {
			return nil, comErr
		}
		if compare == 0 {
			return nil, errors.New("duplicated value")
		} else if compare < 0 { //find insert position
			*slice = append(*slice, nil)
			copy((*slice)[i+1:], (*slice)[i:])
			(*slice)[i] = element
			changed = true //no need to append at the end
			break
		}
	}
	//append at the end
	if !changed {
		*slice = append(*slice, element)
	}
	min := (*slice)[0]
	*slice = (*slice)[1:]

	return min, nil
}

//insert a byte element into an ordered slice at correct position
//delete and return the biggest one
//throw error if value is already in this slice
func InsertToOrderedSliceReturnMax(elementType uint32, slice *[][]byte, element []byte) ([]byte, error) {
	changed := false

	for i, v := range *slice {
		compare, comErr := CompareBytesElement(elementType, element, v)
		if comErr != nil {
			return nil, comErr
		}
		if compare == 0 {
			return nil, errors.New("duplicated value")
		} else if compare < 0 { //find insert position
			*slice = append(*slice, nil)
			copy((*slice)[i+1:], (*slice)[i:])
			(*slice)[i] = element
			changed = true //no need to append at the end
			break
		}
	}
	//append at the end
	if !changed {
		*slice = append(*slice, element)
	}
	max := (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]

	return max, nil
}

//insert a byte element into an ordered slice at correct position
//divide the original slice into two slices, left slice contains small half, right slice contains big half
//if element number is even after insertion, two slices have the same length
//if element number is odd after insertion, right slice contains one more element than left slice
//throw error if value is already in this slice
func InsertToOrderedSliceSplit(elementType uint32, slice *[][]byte, element []byte) ([][]byte, [][]byte, error) {
	changed := false

	for i, v := range *slice {
		compare, comErr := CompareBytesElement(elementType, element, v)
		if comErr != nil {
			return nil, nil, comErr
		}
		if compare == 0 {
			return nil, nil, errors.New("duplicated value")
		} else if compare < 0 { //find insert position
			*slice = append(*slice, nil)
			copy((*slice)[i+1:], (*slice)[i:])
			(*slice)[i] = element
			changed = true //no need to append at the end
			break
		}
	}
	//append at the end
	if !changed {
		*slice = append(*slice, element)
	}

	//create left and right slices
	left := (*slice)[:len(*slice)/2]
	right := (*slice)[len(*slice)/2:]

	return left, right, nil
}

//if err is not nil, print err to standard error
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
