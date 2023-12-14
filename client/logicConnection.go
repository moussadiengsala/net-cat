package client

import (
	"fmt"
	"net"
	"regexp"

	"learn.zone01dakar.sn/net-cat/utils"
)

// this function has as purpose to register a new user
func New(listen net.Listener) {
	var numberOfConnection = 10
	var connections = make(chan struct{}, numberOfConnection)

	var users = Users{Clients: make(map[string]User)}
	for {
		var conn, err = listen.Accept()
		if err != nil {
			break
		}

		if len(connections) == cap(connections) {
			conn.Write([]byte("This chat is full press Enter to exit..."))
			conn.Close()
		} else {
			defer close(connections)
			connections <- struct{}{}

			go func() {
				// Displaying The pinguin logo.
				utils.NewUserUI(conn)

				// keep asking the current user his name until he put something.
				var username string = utils.GetUsername(conn)

				// Here we have to assign the current user a unique id based on the identity on the server.
				var id = regexp.MustCompile(`\.|:`).ReplaceAllString(conn.RemoteAddr().String(), "")
				var newClient = User{Id: id, Connection: conn, Name: username}

				/* Before adding the new user on our list of users we noticed all active users that there is a new
				user who joined the chat
				*/
				UserActivitiesNotifications(newClient, users.Clients, "has joined our chat...")

				users.Clients[newClient.Id] = newClient

				/* display the previous broadcasted
				messages from the channel before he logged in
				*/
				newClient.GetPreviousMessages()
				go HandleConn(conn, newClient, &users, connections)
			}()
		}
	}
}

func HandleConn(conn net.Conn, user User, users *Users, out chan struct{}) {
	defer conn.Close()

	for {
		var err = user.Chat()
		if err != nil {
			UserActivitiesNotifications(user, users.Clients, "has left our chat...")
			fmt.Printf("%s has disconnected.\n", user.Name)
			user.logOut(users, out)
			break
		}
		user.BroadcastChat(users)
	}
}
