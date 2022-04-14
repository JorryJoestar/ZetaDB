package network

import (
	"ZetaDB/parser"
	"ZetaDB/utility"
	"fmt"
	"log"
	"net"
	"sync"
)

type Socket struct {
}

//for singleton pattern
var socketInstance *Socket
var socketOnce sync.Once

//to get kernel, call this function
func GetSocket() *Socket {
	socketOnce.Do(func() {
	})
	return socketInstance
}

func (socket *Socket) Listen() {
	tcp_addr, _ := net.ResolveTCPAddr("tcp4", utility.DEFAULT_SERVER_ADDRESS)

	// listen port 40320
	listener, _ := net.ListenTCP("tcp", tcp_addr)
	for {
		log.Println("[server] listening", tcp_addr.String())
		// wait for client connection
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	buffer := make([]byte, 256)
	conn.Read(buffer)

	Parser := parser.GetParser()
	astNode, err := Parser.ParseSql(string(buffer))

	fmt.Println(string(buffer))
	fmt.Println(astNode)
	fmt.Println(err)

	var astString string
	if err != nil { //parser failed
		astString = "syntax error"
	} else {
		astString = parser.ASTToString(astNode)
	}

	conn.Write([]byte(astString))
	log.Println("[server] response to:", conn.RemoteAddr().String())
}
