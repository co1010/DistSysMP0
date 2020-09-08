# Distributed Systems MP0

To run:
1. Download tcpC.go and tcpS.go
2. Open the terminal and navigate to tcpS.go
3. Type >go run tcpS.go [insert port here]
4. Open another terminal and navigate to tcpC.go
5. Type >go run tcpC.go 127.0.0.1:[port used above]

This code went through many iterations before I found the simplest way of sending a struct over TCP. The TCP connection is set up as a client/server connection, with the client connecting to the server's port. The message struct is encoded using gob and is received by the server, which decodes and prints the message. It then encodes an ACK struct which is sent back to the client. When the client receives the ACK it knows the message has been delievered and exits.
