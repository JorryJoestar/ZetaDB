package storage

import (
	"errors"
	"os"
)

//TODO :construct a class ioManipulator with two methods below, avoid too mant file open and close

//push byte slice into file at byte index pos
func BytesToFile(bytes []byte, pos uint32, fileLocation string) error {
	//open file from fileLocation
	file, openError := os.OpenFile(fileLocation, os.O_RDWR|os.O_SYNC, 0)
	if openError != nil {
		return openError
	}
	defer file.Close()

	var pos64 int64 = int64(0x00000000ffffffff & pos)
	//write bytes at pos64
	n, writeError := file.WriteAt(bytes, pos64)
	if writeError != nil {
		return writeError
	}
	if n != len(bytes) {
		return errors.New("bytes to file len error")
	}
	return nil
}

//fetch bytes of length len from
//func BytesFromFile(pos uint32, len int, fileLocation string) ([]byte, error) {}
