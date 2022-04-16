package network

import (
	"ZetaDB/container"
	"ZetaDB/utility"
	"log"
	"net"
)

func Listen(requestChannel chan container.Session) {
	//listen from this tcp address
	tcp_addr, _ := net.ResolveTCPAddr("tcp4", utility.DEFAULT_SERVER_ADDRESS)

	listener, _ := net.ListenTCP("tcp", tcp_addr)

	for {
		log.Println("[server] listening", tcp_addr.String())

		// wait for client connection
		conn, err := listener.Accept()
		if err != nil {
			log.Println("[server] listening error", err)
			continue
		}

		//fetch request
		buffer := make([]byte, 256)
		conn.Read(buffer)
		currentRequest := container.NewRequestFromBytes(buffer)
		requestUserId := currentRequest.UserId
		requestSql := currentRequest.Sql

		//generate a session
		newSession := container.Session{
			Connection: conn,
			Sql:        requestSql,
			UserId:     requestUserId,
		}

		//push the request into channel
		requestChannel <- newSession
	}
}

func Reply(connection net.Conn, message string, stateCode int32) {

	newResponse := container.NewResponse(stateCode, message)
	responseBytes := newResponse.ResponseToBytes()

	connection.Write([]byte(responseBytes))
	log.Println("[server] response to:", connection.RemoteAddr().String())
}
