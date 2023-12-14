package client

import (
	"fmt"
	"net"
	"os"
	"time"

	"learn.zone01dakar.sn/net-cat/utils"
)

type Users struct {
	Clients map[string]User
}

type User struct {
	Id         string
	Name       string
	Connection net.Conn
	Message    string
}

/*
this function has as purpose to get the message
from a specific user and send it to all users
*/
func (currentUser User) BroadcastChat(users *Users) {

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	var message = fmt.Sprintf("[%s][%s]: %s", currentTime, currentUser.Name, currentUser.Message)
	utils.MessageLogger(message[:len(message)-1], os.O_APPEND|os.O_WRONLY)

	for _, user := range users.Clients {
		if currentUser.Connection.RemoteAddr() != user.Connection.RemoteAddr() {
			if len(currentUser.Message) != 1 {
				var _, errs = user.Connection.Write([]byte(message))
				if errs != nil {
					UserActivitiesNotifications(currentUser, users.Clients, "has left our chat...")
					break
				}
			}
		}
	}
}

func (user User) GetPreviousMessages() {
	previousMessage, _ := os.ReadFile("messagesLogs.txt")
	user.Connection.Write(previousMessage)
}

func (user *User) Chat() error {
	var buf = make([]byte, 1024)
	var len, err = user.Connection.Read(buf)
	if err != nil {
		return err
	}

	user.Message = string(buf[:len])
	return nil
}

/* When
 */
func (user User) logOut(users *Users, out chan struct{}) {
	delete(users.Clients, user.Id)
	<-out
}

/*
	Here we noticed all active users the activities that happen in the room
*/
func UserActivitiesNotifications(client User, clients map[string]User, message string) {
	for _, value := range clients {
		var str = fmt.Sprintf("%s %s\n", client.Name, message)
		value.Connection.Write([]byte(str))
	}
}
