package utils

import (
	"net"
	"regexp"
	"strings"
)

func GetUsername(conn net.Conn) string {
	var username string
	var pattern = regexp.MustCompile(`\s+`)
GetMessage:
	conn.Write([]byte("[ENTER YOUR NAME] : "))
	buffer := make([]byte, 4096)
	length, _ := conn.Read(buffer)
	if length != 0 {
		username = strings.TrimSpace(string(buffer[:length-1]))
		username = pattern.ReplaceAllString(username, " ")
	}
	if len(username) == 0 {
		goto GetMessage
	}
	return username
}
