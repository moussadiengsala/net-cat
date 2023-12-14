package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetPort() string{
	var arguments = os.Args[1:]

	if len(arguments) == 0 {
		return "8989"
	}

	var port, err = strconv.Atoi(arguments[0])
	if err != nil || len(arguments) >= 2 {
		fmt.Println("ℹ️ [USAGE]: ./TCPChat $port")
		os.Exit(0)
	}

	return strconv.Itoa(port)
}