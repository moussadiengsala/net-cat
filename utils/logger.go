package utils

import (
	"fmt"
	"os"
)

// This function has as purpose to log all broadcasteed messages into a file
func MessageLogger(message string, parameter int) {
	// Open the file in append mode with write-only permissions
	file, err := os.OpenFile("messagesLogs.txt", parameter, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(message + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

/*
	* This has purpose to remove all contents in the discussion file after each restart of the server.
	* The zero 0 in os.Turncat is specify the new size of the discussion file, so it will be empty.
*/
func ClearMessageLogger() {
	var err = os.Truncate("messagesLogs.txt", 0)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
}