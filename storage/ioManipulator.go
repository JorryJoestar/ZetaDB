package storage

import (
	"ZetaDB/utility"
	"errors"
	"os"
	"sync"
)

type IOManipulator struct {
	dataFileLocation  string
	indexFileLocation string
	logFileLocation   string
	dataFile          *os.File
	indexFile         *os.File
	logFile           *os.File
}

//to get ioManipulator, call this function
var ioMinstance *IOManipulator
var ioOnce sync.Once

func GetIOManipulator() (*IOManipulator, error) {
	ioOnce.Do(func() {
		ioMinstance = &IOManipulator{
			dataFileLocation:  utility.DEFAULT_DATAFILE_LOCATION,
			indexFileLocation: utility.DEFAULT_INDEXFILE_LOCATION,
			logFileLocation:   utility.DEFAULT_LOGFILE_LOCATION}
	})

	//try to open data file from location dfl
	dFile, dfOpenError := os.OpenFile(utility.DEFAULT_DATAFILE_LOCATION, os.O_RDWR|os.O_SYNC, 0)
	if dfOpenError != nil {
		return nil, dfOpenError
	} else {
		ioMinstance.dataFile = dFile
	}

	//try to open index file from location ifl
	iFile, ifOpenError := os.OpenFile(utility.DEFAULT_INDEXFILE_LOCATION, os.O_RDWR|os.O_SYNC, 0)
	if ifOpenError != nil {
		return nil, ifOpenError
	} else {
		ioMinstance.indexFile = iFile
	}

	//try to open log file from location lfl
	lFile, lfOpenError := os.OpenFile(utility.DEFAULT_LOGFILE_LOCATION, os.O_RDWR|os.O_SYNC, 0)
	if lfOpenError != nil {
		return nil, lfOpenError
	} else {
		ioMinstance.logFile = lFile
	}

	return ioMinstance, nil
}

//always close this iom when all its work is done
func (ioM *IOManipulator) CloseIOM() error {
	dErr := ioM.dataFile.Close()
	if dErr != nil {
		return dErr
	}

	iErr := ioM.indexFile.Close()
	if iErr != nil {
		return iErr
	}

	lErr := ioM.logFile.Close()
	if lErr != nil {
		return lErr
	}

	return nil
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

//push byte slice into log file at byte index pos
func (ioM *IOManipulator) BytesToLogFile(bytes []byte, pos uint32) error {

	//convert pos from uint32 to int64
	var pos64 int64 = int64(0x00000000ffffffff & pos)

	//write bytes at pos64
	n, err := ioM.logFile.WriteAt(bytes, pos64)
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

//fetch bytes of length len from log file
func (ioM *IOManipulator) BytesFromLogFile(pos uint32, len int) ([]byte, error) {

	//slice ready to return
	bytes := make([]byte, len)

	//convert pos from uint32 to int64
	var pos64 int64 = int64(0x00000000ffffffff & pos)

	//read bytes at pos64
	n, err := ioM.logFile.ReadAt(bytes, pos64)
	if err != nil {
		return nil, err
	}
	if n != len {
		return nil, errors.New("bytes number from file len error")
	}

	return bytes, nil
}

//make data file empty
func (ioM *IOManipulator) EmptyDataFile() error {

	//if data file already exists, empty it
	f, err := os.Create(ioM.dataFileLocation)
	if err != nil {
		return err
	}
	ioM.dataFile = f
	return nil
}

//make index file empty
func (ioM *IOManipulator) EmptyIndexFile() error {

	//if index file already exists, empty it
	f, err := os.Create(ioM.indexFileLocation)
	if err != nil {
		return err
	}
	ioM.indexFile = f
	return nil
}

//make log file empty
func (ioM *IOManipulator) EmptyLogFile() error {

	//if index file already exists, empty it
	f, err := os.Create(ioM.logFileLocation)
	if err != nil {
		return err
	}
	ioM.logFile = f
	return nil
}
