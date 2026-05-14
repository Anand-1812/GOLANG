package main

import (
	"fmt"
	"os"
	"log"
	"net"
	"bufio"
	"strings"
)

type Clock struct {
	name string
	address string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("clockwall NewYork=localhost:8010 Tokyo=localhost:8020")
		return
	}

	var clocks []Clock

	for _, arg := range os.Args[1:] {
		parts := strings.Split(arg, "=")

		if len(parts) != 2 {
			log.Fatalf("invalid argument: %s", arg)
		}

		clocks = append(clocks, Clock{
			name: parts[0],
			address: parts[1],
		})
	}

	for _, clock := range clocks {
		go fetchTime(clock)
	}

	select {}
}


func fetchTime(clock Clock) {

	conn, err := net.Dial("tcp", clock.address)
	if err != nil {
		log.Printf("could not connect to %s\n", clock.name)
		return
	}

	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		fmt.Printf("%-10s -> %s\n", clock.name, scanner.Text())
	}
}
