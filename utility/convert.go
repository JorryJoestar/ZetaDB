package utility

import "errors"

func BytesToChar(bytes []byte) (string, error) {
	if len(bytes) != 1 || bytes[0] > 0b01111111 {
		return "", errors.New("BytesToChar: bytes invalid")
	}
	return string(bytes), nil
}
