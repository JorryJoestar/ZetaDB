package container

import "ZetaDB/utility"

type Request struct {
	UserId int32
	Sql    string
}

func NewRequest(userId int32, sql string) *Request {
	return &Request{
		UserId: userId,
		Sql:    sql,
	}
}

func NewRequestFromBytes(requestBytes []byte) *Request {
	userId, _ := utility.BytesToINT(requestBytes[:4])

	requestBytes = requestBytes[4:]
	sql := string(requestBytes)

	return NewRequest(userId, sql)
}

func (request *Request) RequestToBytes() []byte {
	var bytes []byte

	//convert userId to bytes
	bytes = append(bytes, utility.INTToBytes(request.UserId)...)

	//convert sql to bytes
	bytes = append(bytes, []byte(request.Sql)...)

	return bytes
}
