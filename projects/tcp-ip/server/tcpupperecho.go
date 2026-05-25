package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func echoUpper(w io.Writer, r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		log.Printf("received: %s\n", line)

		fmt.Fprintf(w, "%s\n", strings.ToUpper(line))
	}

	if err := scanner.Err();err != nil {
		log.Printf("error: %s", err)
	}
}

func main() {
	const name = "tcpupperecho"
	// so, the log will appear with tab spaced
	log.SetPrefix(name + "\t")

	port := flag.Int("p", 8080, "port to connect on")
	flag.Parse()

	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: *port})
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	log.Printf("listening to localhost: %s", listener.Addr())

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			panic(err)
		}

		log.Printf("received: %s\n", conn.RemoteAddr())

		go func() {
			defer conn.Close()
			echoUpper(conn, conn)
		}()
	}
}
