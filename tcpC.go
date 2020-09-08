package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"time"
)

type Message struct {
	To, From, Title, Content string
	Date time.Time
}

type Ack struct {
	Acknowledge int
}

func main() {
	// Check that the user included an ip when running the program
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	// Connect to the server
	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	checkError(err)

	// Get the message from the user and save it
	var to, from, title, content string
	fmt.Print("To: ")
	fmt.Scanln(&to)
	fmt.Print("From: ")
	fmt.Scanln(&from)
	fmt.Print("Title: ")
	fmt.Scanln(&title)
	fmt.Print("Content: ")
	fmt.Scanln(&content)
	date := time.Now()
	email := Message{to, from, title, content, date}

	// Encode the message
	encoder := gob.NewEncoder(c)
	encoder.Encode(email)

	// Decode the ACK
	decoder := gob.NewDecoder(c)
	var ack Ack
	decoder.Decode(&ack)
	if ack.Acknowledge == 1 {
		fmt.Println("ACK received. Exiting...")
		os.Exit(3)
	}
}

// Consolidated repeated error checks into a single function
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
