package internals

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func Print(n *html.Node) {
	walk(nil, n, 0)
}

func walk(stack []string, n *html.Node, depth int) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		indent := strings.Repeat(" ", depth)
		fmt.Printf("%s<%s>\n", indent, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walk(stack, c, depth+1)
	}
}

// to see the exact slice at each level
func PrintRaw(n *html.Node) {
	walkRaw(nil, n)
}

func walkRaw(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walkRaw(stack, c)
	}
}
