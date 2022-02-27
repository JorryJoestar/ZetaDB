package storage

import (
	"errors"
	"os"
	"sync"
)

type IOManipulator struct {
	dataFileLocation  string
	indexFileLocation string
	dataFile          *os.File
	indexFile         *os.File
}

//to get ioManipulator, call this function
var ioMinstance *IOManipulator
var ioOnce sync.Once

func GetIOManipulator(dfl string, ifl string) (*IOManipulator, error) {
	ioOnce.Do(func() {
		ioMinstance = &IOManipulator{
			dataFileLocation:  dfl,
			indexFileLocation: ifl}
	})

	//try to open data file from location dfl
	dFile, dfOpenError := os.OpenFile(dfl, os.O_RDWR|os.O_SYNC, 0)
	if dfOpenError != nil {
		return nil, dfOpenError
	} else {
		ioMinstance.dataFile = dFile
	}

	//try to open index file from location ifl
	iFile, ifOpenError := os.OpenFile(ifl, os.O_RDWR|os.O_SYNC, 0)
	if ifOpenError != nil {
		return nil, ifOpenError
	} else {
		ioMinstance.indexFile = iFile
	}

	return ioMinstance, nil
}

//always close this iom when all its work is done
func (ioM *IOManipulator) CloseIOM() {
	ioM.dataFile.Close()
	ioM.indexFile.Close()
}

//push byte slice into data file at byte index pos
func (ioM *IOManipulator) BytesToDataFile(bytes []byte, pos uint32) error {

	//convert pos from uint32 to int64
	var pos64 int64 = int64(0x00000000ffffffff & pos)

	//write bytes at pos64
	n, err := ioM.dataFile.WriteAt(bytes, pos64)
	if err != nil {
		return err
	}
	if n != len(bytes) {
		return errors.New("bytes to file len error")
	}
	return nil
}

//push byte slice into index file at byte index pos
func (ioM *IOManipulator) BytesToIndexFile(bytes []byte, pos uint32) error {

	//convert pos from uint32 to int64
	var pos64 int64 = int64(0x00000000ffffffff & pos)

	//write bytes at pos64
	n, err := ioM.indexFile.WriteAt(bytes, pos64)
	if err != nil {
		return err
	}
	if n != len(bytes) {
		return errors.New("bytes number to file len error")
	}
	return nil
}

//fetch bytes of length len from data file
func (ioM *IOManipulator) BytesFromDataFile(pos uint32, len int) ([]byte, error) {

	//slice ready to return
	bytes := make([]byte, len)

	//convert pos from uint32 to int64
	var pos64 int64 = int64(0x00000000ffffffff & pos)

	//read bytes at pos64
	n, err := ioM.dataFile.ReadAt(bytes, pos64)
	if err != nil {
		return nil, err
	}
	if n != len {
		return nil, errors.New("bytes number from file len error")
	}

	return bytes, nil
}

//fetch bytes of length len from index file
func (ioM *IOManipulator) BytesFromIndexFile(pos uint32, len int) ([]byte, error) {

	//slice ready to return
	bytes := make([]byte, len)

	//convert pos from uint32 to int64
	var pos64 int64 = int64(0x00000000ffffffff & pos)

	//read bytes at pos64
	n, err := ioM.indexFile.ReadAt(bytes, pos64)
	if err != nil {
		return nil, err
	}
	if n != len {
		return nil, errors.New("bytes number from file len error")
	}

	return bytes, nil
}
