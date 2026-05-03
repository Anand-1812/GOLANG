package main

import (
	"fmt"
	"net/http"
	"os"

	"html-pretty/pretty"
	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <url>")
		os.Exit(1)
	}

	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Parse error:", err)
		os.Exit(1)
	}

	fmt.Println("----------Pretty HTML----------")
	// write the pretty function
	pretty.PrettyPrint(doc)
}
