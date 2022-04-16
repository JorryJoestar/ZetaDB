package container

import "net"

type Request struct {
	Connection net.Conn
	Sql        string
	UserSql    string
}
