package storage

import (
	. "ZetaDB/container"
	. "ZetaDB/utility"
	"errors"
)

/*
									log page structure
    -------------------------------------------------------------------------------------
    |     logPageId      |    logPrePageId    |    logNextPageId   |       logNum       |
    -------------------------------------------------------------------------------------
    |                              log 0                           |        log 1       |
	-------------------------------------------------------------------------------------
    |                 log 1                   |                  log 2                  |
    -------------------------------------------------------------------------------------
    |       log 2        |                            log 3                             |
	-------------------------------------------------------------------------------------
    |                                    . . . . . .                                    |
	-------------------------------------------------------------------------------------
    |                           log N                              |    padding bytes   |
	-------------------------------------------------------------------------------------

	~logPageId
		-uint32, 4 bytes
		-page id of this page in log file

	~logPrePageId
		-uint32, 4 bytes
		-page id of the previous page

	~logNextPageId
		-uint32, 4 bytes
		-page id of the next page

	~logNum
		-int32, 4 bytes
		-log number in this page

*/

type logPage struct {
	logPageId     uint32
	logPrePageId  uint32
	logNextPageId uint32
	logNum        int32
	logs          []*Log
}

//create a new log page
func NewLogPage(logPageId uint32, logPrePageId uint32, logNextPageId uint32) *logPage {
	lp := &logPage{
		logPageId:     logPageId,
		logPrePageId:  logPrePageId,
		logNextPageId: logNextPageId,
		logNum:        0}

	return lp
}

//create a new log page from bytes
//throw error if bytes length invalid
func NewLogPageFromBytes(bytes []byte) (*logPage, error) {
	//throw error if bytes length invalid
	if len(bytes) != DEFAULT_PAGE_SIZE {
		return nil, errors.New("bytes length invalid")
	}

	lp := &logPage{}

	//get logPageId
	logPageId, _ := BytesToUint32(bytes[:4])
	lp.logPageId = logPageId
	bytes = bytes[4:]

	//get logPrePageId
	logPrePageId, _ := BytesToUint32(bytes[:4])
	lp.logPrePageId = logPrePageId
	bytes = bytes[4:]

	//get logNextPageId
	logNextPageId, _ := BytesToUint32(bytes[:4])
	lp.logNextPageId = logNextPageId
	bytes = bytes[4:]

	//get logNum
	logNum, _ := BytesToINT(bytes[:4])
	lp.logNum = logNum
	bytes = bytes[4:]

	//get logs
	var i int32
	for i = 0; i < logNum; i++ {
		log, _ := NewLogFromBytes(bytes[:9])
		lp.logs = append(lp.logs, log)
		bytes = bytes[9:]
	}

	return lp, nil
}

//convert this log page to bytes, ready to be pushed into disk
func (lp *logPage) LogPageToBytes() []byte {
	var bytes []byte

	//logPageId
	bytes = append(bytes, Uint32ToBytes(lp.logPageId)...)

	//logPrePageId
	bytes = append(bytes, Uint32ToBytes(lp.logPrePageId)...)

	//logNextPageId
	bytes = append(bytes, Uint32ToBytes(lp.logNextPageId)...)

	//logNum
	bytes = append(bytes, INTToBytes(lp.logNum)...)

	//logs
	for _, log := range lp.logs {
		bytes = append(bytes, log.LogToBytes()...)
	}

	//padding bytes
	var i int32
	for i = 0; i < lp.LogPageVacantByteNum(); i++ {
		bytes = append(bytes, byte(0))
	}

	return bytes
}

//logPageId getter
func (lp *logPage) LogPageGetLogPageId() uint32 {
	return lp.logPageId
}

//logPrePageId getter
func (lp *logPage) LogPageGetLogPrePageId() uint32 {
	return lp.logPrePageId
}

//logPrePageId setter
func (lp *logPage) LogPageSetLogPrePageId(id uint32) {
	lp.logPrePageId = id
}

//logNextPageId getter
func (lp *logPage) LogPageGetLogNextPageId() uint32 {
	return lp.logNextPageId
}

//logNextPageId setter
func (lp *logPage) LogPageSetLogNextPageId(id uint32) {
	lp.logNextPageId = id
}

//logNum getter
func (lp *logPage) LogPageGetLogNum() int32 {
	return lp.logNum
}

//logs getter
func (lp *logPage) LogPageGetLogs() []*Log {
	return lp.logs
}

//insert a log into this page
//throw error if no enough space to complete this insertion
func (lp *logPage) LogPageInsertLog(l *Log) error {
	//throw error if no enough space to complete this insertion
	if lp.LogPageVacantByteNum() < 9 {
		return errors.New("not enough space")
	}

	lp.logNum++
	lp.logs = append(lp.logs, l)

	return nil
}

//return how many bytes this logPage can hold
func (lp *logPage) LogPageVacantByteNum() int32 {
	var size int32
	size = int32(DEFAULT_PAGE_SIZE)

	//head takes 16 bytes
	size -= 16

	//each log takes 9 bytes
	size -= 9 * lp.logNum

	return size
}

//if this logPage can hold no more log, return true
func (lp *logPage) LogPageIsFull() bool {
	if lp.LogPageVacantByteNum() < 9 {
		return true
	} else {
		return false
	}
}
