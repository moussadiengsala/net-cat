package utils

import (
	"fmt"
	"net"
	"os"
)

// For displaying The pinguin logo.
func NewUserUI(conn net.Conn) {
	// to notice the server a new connection
	fmt.Println("A new user connected")

	pinguinDraw, _ := os.ReadFile("pinguin.txt")
	conn.Write([]byte("Welcome to TCP-Chat!"))
	conn.Write([]byte("\n"))
	conn.Write([]byte(pinguinDraw))
	conn.Write([]byte("\n"))
}
