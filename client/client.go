package client

import (
	"bufio"
	"chatserver/message"
	"chatserver/utils"
	"encoding/json"
	"net"
)

// Client represents a chatserver's client
type Client struct {
	Name string
	Conn net.Conn
}

// Start starts the client
func (client *Client) Start(messages chan message.Message) {
	reader := bufio.NewReader(client.Conn)

	for {
		line, _, err := reader.ReadLine()
		utils.Check(err)

		var m message.Message
		err = json.Unmarshal(line, &m)
		utils.Check(err)
		m.Sender = client.Name

		messages <- m
	}
}
