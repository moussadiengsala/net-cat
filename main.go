package main

import (
	"learn.zone01dakar.sn/net-cat/client"
	"learn.zone01dakar.sn/net-cat/server"
	"learn.zone01dakar.sn/net-cat/utils"
)

func main() {
	server := server.ChatServer{IP: "127.0.0.1", Port: utils.GetPort()}
	listener, _ := server.Start()
	defer listener.Close()

	// Clear all recent content in the discussion file before adding any new user.
	utils.ClearMessageLogger()

	// This allows our server to continue accepting new user.
	client.New(listener)
}
