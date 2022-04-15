package network

import (
	"ZetaDB/container"
	"ZetaDB/utility"
	"log"
	"net"
)

func Listen(requestChannel chan container.Request) {
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

		//fetch sql
		buffer := make([]byte, 256)
		conn.Read(buffer)
		sqlString := string(buffer)

		//generate a request
		newRequest := container.Request{
			Connection: conn,
			Sql:        sqlString,
		}

		//push the request into channel
		requestChannel <- newRequest
	}
}

func Reply(connection net.Conn, result string) {
	connection.Write([]byte(result))
	log.Println("[server] response to:", connection.RemoteAddr().String())
}
