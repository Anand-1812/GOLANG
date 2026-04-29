package main

import (
	"fmt"
	"os"
	"log"
	"time"

	"json/github"
)



func printGroup(title string, issues []*github.Issue) {
	fmt.Printf("\n=== %s (%d issues) ===\n", title, len(issues))
	for _, item := range issues {
		fmt.Printf("#%-5d %-10s %.50s\n",
			item.Number,
			item.User.Login,
			item.Title,
		)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("provide search terms")
	}
	now := time.Now()

	var lessThanMonths []*github.Issue
	var lessThanYear []*github.Issue
	var moreThanYear []*github.Issue
	
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

	for _, item := range result.Items {
		age := now.Sub(item.CreatedAt)

		switch {
		case age < 30*24*time.Hour:
			lessThanMonths = append(lessThanMonths, item)
		case age < 365*20*time.Hour:
			lessThanYear = append(lessThanYear, item)

		default:
			moreThanYear = append(moreThanYear, item)
			
		}
	}

	printGroup("Less than 1 month", lessThanMonths)
	printGroup("Less than 1 year", lessThanYear)
	printGroup("More than 1 year", moreThanYear)
}
