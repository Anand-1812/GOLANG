package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	const name = "writetcp"
	// so, the log will appear with tab spaced
	log.SetPrefix(name + "\t")

	port := flag.Int("p", 8080, "port to connect to")
	flag.Parse()

	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{Port: *port})
	if err != nil {
		log.Fatalf("Error connecting to localhost:%d : %v", *port, err)
	}

	log.Printf("connected to %s: will forward stdin", conn.RemoteAddr())

	defer conn.Close()

	// spawns a goroutine to read incoming lines from the server
	go func() {
		connScanner := bufio.NewScanner(conn)

		for connScanner.Scan() {
			fmt.Printf("%s\n", connScanner.Text())
		}

		if err := connScanner.Err(); err != nil {
			log.Fatalf("error reading from %s: %v", conn.RemoteAddr(), err)
		}
	}()

	stdinScanner := bufio.NewScanner(os.Stdin)
	for stdinScanner.Scan() {
		log.Printf("sent: %s\n", stdinScanner.Text())

		if _, err := conn.Write(stdinScanner.Bytes());err != nil {
			log.Fatalf("error writing to %s: %v", conn.RemoteAddr(), err)
		}

		if _, err := conn.Write([]byte("\n"));err != nil {
			log.Fatalf("error writing to %s: %v", conn.RemoteAddr(), err)
		}
	}

	if err := stdinScanner.Err(); err != nil {
		log.Fatalf("stdin error: %v", err)
	}

}
