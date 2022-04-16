package container

import "ZetaDB/utility"

type Response struct {
	StateCode int32
	Message   string
}

func NewResponse(stateCode int32, message string) *Response {
	return &Response{
		StateCode: stateCode,
		Message:   message,
	}
}

func NewResponseFromBytes(responseBytes []byte) *Response {
	stateCode, _ := utility.BytesToINT(responseBytes[:4])

	responseBytes = responseBytes[4:]
	message := string(responseBytes)

	return NewResponse(stateCode, message)
}

func (response *Response) ResponseToBytes() []byte {
	var bytes []byte

	//convert stateCode to bytes
	bytes = append(bytes, utility.INTToBytes(response.StateCode)...)

	//convert message to bytes
	bytes = append(bytes, []byte(response.Message)...)

	return bytes
}
