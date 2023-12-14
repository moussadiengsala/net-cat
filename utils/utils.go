package utils

import (
	"net"
)

func GetUsername(conn net.Conn) string {
	var username string
GetMessage:
	conn.Write([]byte("[ENTER YOUR NAME] : "))
	buffer := make([]byte, 4096)
	length, _ := conn.Read(buffer)
	if length != 0 {
		username = string(buffer[:length-1])
	}
	if len(username) == 0 {
		goto GetMessage
	}
	return username
}
