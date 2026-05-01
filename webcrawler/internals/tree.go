package internals

import (
	"fmt"

	"golang.org/x/net/html"
)

// summary of each node type found
type Summary struct {
	Elements int
	Text     int
	Comments int
	Doctype  int
	Other    int
}

// traverse enitre html node tree and return summary
func Analyze(n *html.Node) Summary {
	var s Summary
	analyze(n, &s)

	return s
}

// recursive worker, for all the recursion call to update the same struct
func analyze(n *html.Node, s *Summary) {
	switch n.Type {
	case html.ElementNode:
		s.Elements++
	case html.TextNode:
		s.Text++
	case html.CommentNode:
		s.Comments++
	case html.DoctypeNode:
		s.Doctype++

	default:
		s.Other++
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		analyze(child, s)
	}
}

func PrintSummary(s Summary) {
	fmt.Println("=== Document Summary ===")
	fmt.Printf("  Element nodes : %d\n", s.Elements)
	fmt.Printf("  Text nodes    : %d\n", s.Text)
	fmt.Printf("  Comment nodes : %d\n", s.Comments)
	fmt.Printf("  Doctype nodes : %d\n", s.Doctype)
	fmt.Printf("  Other nodes   : %d\n", s.Other)
}
