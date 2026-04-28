package main

import (
	"fmt"
	"os"
	"log"

	"json/github"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("provide search terms")
	}

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues found:\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %-10s %0.50s\n", 
			item.Number,
			item.User.Login,
			item.Title,
		)
	}
}
