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
		-1: from logPageId to data file at objectPageId
		-2: from logPageId to index file at objectPageId

	~logPageId
		the snapshot id in log file

	~objectPageId
		the id where this snapshot page should be stored at
*/

type log struct {
	fileMode     int32
	logPageId    uint32
	objectPageId uint32
}

//create a new log
//throw error if fileMode is not 1 or 2
func NewLog(fileMode int32, logPageId uint32, objectPageId uint32) (*log, error) {
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
//throw error if bytes length is not 12
func NewLogFromBytes(bytes []byte) (*log, error) {
	//throw error if bytes length is not 12
	if len(bytes) != 12 {
		return nil, errors.New("byte slice length invalid")
	}

	fileMode, _ := BytesToINT(bytes[0:4])
	logPageId, _ := BytesToUint32(bytes[4:8])
	objectPageId, _ := BytesToUint32(bytes[8:12])

	l := &log{
		fileMode:     fileMode,
		logPageId:    logPageId,
		objectPageId: objectPageId}

	return l, nil
}

//convert this log to byte slice, ready to push into disk
func (l *log) LogToBytes() []byte {
	var bytes []byte

	bytes = append(bytes, INTToBytes(l.fileMode)...)
	bytes = append(bytes, Uint32ToBytes(l.logPageId)...)
	bytes = append(bytes, Uint32ToBytes(l.objectPageId)...)

	return bytes
}

//fileMode getter
func (l *log) LogGetFileMode() int32 {
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
