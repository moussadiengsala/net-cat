package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetPort() string{
	var arguments = os.Args[1:]
	if len(arguments) == 1 {
		var port, err = strconv.Atoi(arguments[0])
		if err != nil {
			fmt.Println("[USAGE]: ./TCPChat $port")
			os.Exit(0)
		}
		return strconv.Itoa(port)
	} else if len(arguments) >= 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(0)
	}
	return "8989"
}