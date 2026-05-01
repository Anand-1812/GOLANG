package main

import (
	"fmt"
	"os"

	"webcrawler/internals"
)

func main() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	url := os.Args[2]

	doc, err := internals.FromURL(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	switch command {
	case "links":
		fmt.Printf("=== Links found in %s ===\n", url)
		found := internals.Extract(doc)
		if len(found) == 0 {
			fmt.Println("No links found.")
		}
		for i, link := range found {
			fmt.Printf("  [%d] %s\n", i+1, link)
		}

	case "outline":
		fmt.Printf("=== HTML Tree (indented) for %s ===\n", url)
		internals.Print(doc)

	case "raw":
		fmt.Printf("=== HTML Tree (book style) for %s ===\n", url)
		internals.PrintRaw(doc)

	case "summary":
		summary := internals.Analyze(doc)
		internals.PrintSummary(summary)

	case "all":
		// Run all four tools on the same document
		fmt.Printf("=== Links found in %s ===\n", url)
		found := internals.Extract(doc)
		for i, link := range found {
			fmt.Printf("  [%d] %s\n", i+1, link)
		}

		fmt.Println()
		summary := internals.Analyze(doc)
		internals.PrintSummary(summary)

		fmt.Println()
		fmt.Println("=== HTML Tree (indented) ===")
		internals.Print(doc)

	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}

}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  crawler links   <url>   — extract all links")
	fmt.Println("  crawler outline <url>   — print indented HTML tree")
	fmt.Println("  crawler raw     <url>   — print HTML tree (book style)")
	fmt.Println("  crawler summary <url>   — print node type counts")
	fmt.Println("  crawler all     <url>   — run all of the above")
}
