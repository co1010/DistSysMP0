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
	// Check that the user included a port when running the program
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	// Set the port and designate as a listen server
	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	checkError(err)
	c, err := l.Accept()
	checkError(err)

	// Decode the message and print it
	decoder := gob.NewDecoder(c)
	var email Message
	decoder.Decode(&email)
	fmt.Printf("To: %s\nFrom: %s\nTitle: %s\nContent: %s\nDate: %s\n",
		email.To, email.From, email.Title, email.Content, email.Date)

	encoder := gob.NewEncoder(c)
	ack := Ack{1}
	encoder.Encode(ack)

	l.Close()

	os.Exit(3)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
