package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var (
	host, path, method string
	port int
)

func main() {
	flag.StringVar(&method, "method", "GET", "http method to use")
	flag.StringVar(&host, "host", "localhost", "host to connect to")
	flag.IntVar(&port, "port", 8080, "port to connect to")
	flag.StringVar(&path, "path", "/", "path to request")
	flag.Parse()

	ip, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d\n", host, port))
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, ip)
	if err != nil {
		panic(err)
	}

	log.Printf("connected to %s (@ %s)", host, conn.RemoteAddr())
	defer conn.Close()

	var reqFields = []string{
		fmt.Sprintf("%s %s HTTP/1.1", method, path),
		"Host: " + host,
		"User-Agent: httpget",
		"",
	}

	request := strings.Join(reqFields, "\r\n") + "\r\n"

	conn.Write([]byte(request))
	log.Printf("sent request:\n%s", request)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Bytes()
		if _, err := fmt.Fprintf(os.Stdout, "%s\n", line); err != nil {
			log.Printf("error writting to connection: %s", err)
		}
	}

	if scanner.Err() != nil {
		log.Printf("error reading from connection: %s", err)
		return
	}
}
