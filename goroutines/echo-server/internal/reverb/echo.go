package reverb

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func Echo(conn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(conn, "t", strings.ToUpper(shout))
	time.Sleep(delay)

	fmt.Fprintln(conn, "t", shout)
	time.Sleep(delay)

	fmt.Fprintln(conn, "t", strings.ToLower(shout))
}

func HandleConn(conn net.Conn) {
	defer conn.Close()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		text := input.Text()

		go Echo(conn, text, 1*time.Second)
	}
}

