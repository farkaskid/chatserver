package main

import (
	"chatserver/client"
	"chatserver/message"
	"chatserver/server"
)

func main() {
	s := &server.Server{
		Clients:  make(map[string]*client.Client),
		Messages: make(chan message.Message, 100),
	}

	s.Start(":8000")
}
