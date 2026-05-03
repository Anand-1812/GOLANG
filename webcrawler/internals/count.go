package internals

import (
	"strings"

	"golang.org/x/net/html"
)

// multiple return
func CountWordAndImg(n *html.Node) (int, int) {
	var words, img int

	var visit func(*html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.TextNode {
			words += len(strings.Fields(n.Data))
		}

		if n.Type == html.ElementNode && n.Data == "img" {
			img++
		}

		for c := n.FirstChild;c != nil;c = c.NextSibling {
			visit(c)
		}
	}

	visit(n)
	return words, img
}
