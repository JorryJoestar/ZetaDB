package utility

import "errors"

//convert one byte from a byte slice to CHAR(string)
//error if this byte out of range of ascii
//error if this byte slice length is not 1
func BytesToCHAR(bytes []byte) (string, error) {
	if len(bytes) != 1 {
		return "", errors.New("length of byte slice is not 1")
	}
	if bytes[0] > 0b01111111 {
		return "", errors.New("out of range of ascii")
	}
	return string(bytes), nil
}

//convert a string whose length is 1 into a byte slice
//error if string length is not 1
func CHARToBytes(c string) ([]byte, error) {
	if len(c) != 1 {
		return nil, errors.New("length of string is not 1")
	}
	return []byte(c), nil
}

//convert one byte to CHAR(string)
//error if this byte out of range of ascii
func ByteToCHAR(b byte) (string, error) {
	if b > 0b01111111 {
		return "", errors.New("out of range of ascii")
	}
	var bytes []byte
	bytes = append(bytes, b)
	return string(bytes), nil
}

//convert CHAR(string) to a byte
//error if string length is not 1
func CHARToByte(c string) (byte, error) {
	if len(c) != 1 {
		return 0, errors.New("length of string is not 1")
	}
	bytes := []byte(c)
	return bytes[0], nil
}

//convert byte slice to VARCHAR(string)
//error if any one byte in byte slice out of range of ascii
func BytesToVARCHAR(bytes []byte) (string, error) {
	for _, b := range bytes {
		if b > 0b01111111 {
			return "", errors.New("byte out of range of ascii")
		}
	}
	return string(bytes), nil
}

func VARCHARToBytes(s string) ([]byte, error) {
	bytes := []byte(s)
	for _, b := range bytes {
		if b > 0b01111111 {
			return nil, errors.New("character out of range of ascii")
		}
	}
	return bytes, nil
}
