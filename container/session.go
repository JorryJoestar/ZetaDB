package container

import "net"

type Session struct {
	Connection net.Conn
	UserId     int32
	Sql        string
}
