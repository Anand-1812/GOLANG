package main

import (
	"fmt"
	"net"
	"strings"
)

type Client struct {
	Name string
	Conn net.Conn
}

type Server struct {
	Address string
	clients map[string]*Client
}

func CreateServer(addr string) *Server {
	return &Server{
		Address: addr,
		clients: make(map[string]*Client),
	}
}

func CreateClient(username string, conn net.Conn) *Client {
	return &Client{
		Name: username,
		Conn: conn,
	}
}

func (s *Server) HandleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Println("New user connected")

	_, err := conn.Write([]byte("Enter username:"))
	if err != nil {return}

	// take the user name and proceed for storing
	name := make([]byte, 1024)
	n, err := conn.Read(name)
	if err != nil {
		return
	}

	username := strings.TrimSpace(string(name[:n]))
	client := CreateClient(username, conn)
	s.clients[username] = client

	fmt.Printf("%s joins the chat\n", username)

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("%s disconnected\n", username)
			return
		}

		msg := strings.TrimSpace(string(buffer[:n]))
		fmt.Printf("[%s] : %s\n", username, msg)
	}
}

func (s *Server) Start() error {
	// create a listener
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error during connection:%v", err)
	}

	fmt.Printf("Server listening on %s\n", s.Address)
	for {
		// creating conection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error during connection:%v", err)
			continue
		}

		// spwans a routine for handelling client
		go s.HandleClient(conn)
	}
}

func main() {
	server := CreateServer(":8080")
	server.Start()
}
