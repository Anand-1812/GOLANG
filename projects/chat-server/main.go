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

func (s *Server) Broadcast(sender string, msg string) {
  for name, client := range s.clients {
    if name == sender {
      continue
    }

    _, err := client.Conn.Write([]byte(msg))
    if err != nil {
      fmt.Printf("Failed to send to %s: %v\n", name, err)
    }
  }
}

func (s *Server) HandleClient(conn net.Conn) {
  defer conn.Close()

  fmt.Println("New user connected")

  var username string

  for {
    _, err := conn.Write([]byte("Enter username: "))
    if err != nil {
      return
    }

    nameBuf := make([]byte, 1024)

    n, err := conn.Read(nameBuf)
    if err != nil {
      return
    }

    username = strings.TrimSpace(string(nameBuf[:n]))

    if username == "" {
      conn.Write([]byte("Username cannot be empty\n"))
      continue
    }

    if _, exists := s.clients[username]; exists {
      conn.Write([]byte("Username already taken\n"))
      continue
    }

    break
  }

  client := CreateClient(username, conn)

  s.clients[username] = client

  fmt.Printf("%s joined the chat\n", username)

  s.Broadcast(
    username,
    fmt.Sprintf("%s joined the chat\n", username),
  )

  buffer := make([]byte, 1024)

  for {
    n, err := conn.Read(buffer)
    if err != nil {
      fmt.Printf("%s left the chat\n", username)

      delete(s.clients, username)

      s.Broadcast(
        username,
        fmt.Sprintf("%s left the chat\n", username),
      )

      return
    }

    msg := strings.TrimSpace(string(buffer[:n]))

    if msg == "" {
      continue
    }

    fmt.Printf("[%s]: %s\n", username, msg)

    s.Broadcast(
      username,
      fmt.Sprintf("[%s]: %s\n", username, msg),
    )
  }
}

func (s *Server) Start() {
  listener, err := net.Listen("tcp", s.Address)
  if err != nil {
    panic(err)
  }

  fmt.Printf("Server listening on %s\n", s.Address)

  for {
    conn, err := listener.Accept()
    if err != nil {
      fmt.Printf("Accept error: %v\n", err)
      continue
    }

    go s.HandleClient(conn)
  }
}

func main() {
  server := CreateServer(":8080")
  server.Start()
}
