package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Run: go run cmd/main.go <file_name.torrnet>")
		return
	}


	fmt.Println("Torrent Info")
	// write further
}
