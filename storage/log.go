package storage

import (
	. "ZetaDB/utility"
	"errors"
)

/*
                          log structure
   ----------------------------------------------------------------
   |     fileMode       |     logPageId      |    objectPageId    |
   ----------------------------------------------------------------

	~fileMode
		-uint8, 1 byte
		-1: from logPageId to data file at objectPageId
		-2: from logPageId to index file at objectPageId

	~logPageId
		-uint32, 4 bytes
		-the snapshot id in log file

	~objectPageId
		-uint32, 4 bytes
		-the id where this snapshot page should be stored at
*/

type log struct {
	fileMode     uint8
	logPageId    uint32
	objectPageId uint32
}

//create a new log
//throw error if fileMode is not 1 or 2
func NewLog(fileMode uint8, logPageId uint32, objectPageId uint32) (*log, error) {
	//throw error if fileMode is not 1 or 2
	if fileMode != 1 && fileMode != 2 {
		return nil, errors.New("fileMode invalid")
	}

	l := &log{
		fileMode:     fileMode,
		logPageId:    logPageId,
		objectPageId: objectPageId}

	return l, nil
}

//create a new log from bytes
//throw error if bytes length is not 9
func NewLogFromBytes(bytes []byte) (*log, error) {
	//throw error if bytes length is not 9
	if len(bytes) != 9 {
		return nil, errors.New("byte slice length invalid")
	}

	fileMode := bytes[0]
	logPageId, _ := BytesToUint32(bytes[1:5])
	objectPageId, _ := BytesToUint32(bytes[5:9])

	l := &log{
		fileMode:     fileMode,
		logPageId:    logPageId,
		objectPageId: objectPageId}

	return l, nil
}

//convert this log to byte slice, ready to push into disk
func (l *log) LogToBytes() []byte {
	var bytes []byte

	bytes = append(bytes, l.fileMode)
	bytes = append(bytes, Uint32ToBytes(l.logPageId)...)
	bytes = append(bytes, Uint32ToBytes(l.objectPageId)...)

	return bytes
}

//fileMode getter
func (l *log) LogGetFileMode() uint8 {
	return l.fileMode
}

//logPageId getter
func (l *log) LogGetlogPageId() uint32 {
	return l.logPageId
}

//objectPageId getter
func (l *log) LogGetobjectPageId() uint32 {
	return l.objectPageId
}
