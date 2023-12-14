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
	adress := fmt.Sprintf("%s:%s", s.IP, s.Port)
	fmt.Println("Server starting...")
	listener, err := net.Listen("tcp", adress)
	if err != nil {
		log.Fatal(err)
		return nil, err
	} 
	fmt.Println("Server listening on port:", s.Port)
	
	return listener, nil
}
