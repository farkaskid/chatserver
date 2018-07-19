package server

import (
	"bufio"
	"chatserver/client"
	"chatserver/message"
	"chatserver/utils"
	"log"
	"net"
	"strconv"
	"sync"
)

const (
	// SUCCESS Represents successful execution of a task
	SUCCESS = 0
)

// Server represents a chatserver
type Server struct {
	Clients  map[string]*client.Client
	Messages chan message.Message
	sync.Mutex
}

// Start starts the server
func (server *Server) Start(port string) {
	ln, err := net.Listen("tcp", port)
	utils.Check(err)
	log.Println("Server started on port " + port)

	go server.listenClients(ln)

	for m := range server.Messages {
		go server.handleMessage(m)
	}
}

// handleMessage handles a single message. Server reads the recipients from the message
// and push them the content of the message.
func (server *Server) handleMessage(m message.Message) {
	for _, recipient := range m.Recipients {
		client, present := server.Clients[recipient]
		if !present {
			continue
		}

		_, err := client.Conn.Write([]byte(m.Content + "\n"))
		utils.Check(err)
	}

	senderConn := server.Clients[m.Sender].Conn
	_, err := senderConn.Write([]byte(strconv.Itoa(SUCCESS) + "\n"))
	utils.Check(err)
}

// listenClients listens for new clients
func (server *Server) listenClients(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		utils.Check(err)

		go server.addClient(conn)
	}
}

// addClient adds a new client by reading a name from the connection
func (server *Server) addClient(conn net.Conn) {
	name, _, err := bufio.NewReader(conn).ReadLine()
	utils.Check(err)

	c := &client.Client{
		Name: string(name),
		Conn: conn,
	}

	server.Lock()
	server.Clients[string(name)] = c
	server.Unlock()

	log.Println("New client:" + string(name) + " is added.")
	_, err = conn.Write([]byte(strconv.Itoa(SUCCESS) + "\n"))
	utils.Check(err)

	go c.Start(server.Messages)
}
