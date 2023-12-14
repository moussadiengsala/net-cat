package server

import (
	"fmt"
	"log"
	"net"
)

type ChatServer struct {
	IP   string
	Port string
}

func (s *ChatServer) Start() (net.Listener, error) {
	address := fmt.Sprintf("%s:%s", s.IP, s.Port)
	fmt.Println("🌐 Server starting...")

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("❌ Error listening to: ", address, err)
		return nil, err
	} 
	fmt.Println("🌐 Server listening at: ", address)
	
	return listener, nil
}
