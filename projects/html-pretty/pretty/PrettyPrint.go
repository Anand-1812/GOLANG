package pretty

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var depth int

func PrettyPrint(n *html.Node) {
	depth = 0
	ForEachNode(n, startElement, endElement)
}

func startElement(n *html.Node) {
	switch n.Type {

	case html.ElementNode:
		indent()

		fmt.Printf("<%s", n.Data)

		// print attributes
		for _, attr := range n.Attr {
			fmt.Printf(" %s=\"%s\"", attr.Key, attr.Val)
		}

		// self-closing if no children
		if n.FirstChild == nil {
			fmt.Println("/>")
		} else {
			fmt.Println(">")
			depth++
		}

	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if text != "" {
			indent()
			fmt.Println(text)
		}

	case html.CommentNode:
		indent()
		fmt.Printf("<!-- %s -->\n", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		indent()
		fmt.Printf("</%s>\n", n.Data)
	}
}

func indent() {
	fmt.Printf("%*s", depth*2, "")
}
