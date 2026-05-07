package internals

import "golang.org/x/net/html"

type Tags struct {
	Type string
	URL string
}

var tagAttrMap = map[string]struct {
	attr string
	resType string
}{
	"a": {attr: "href", resType: "link"},
	"img": {attr: "src", resType: "image"},
	"script": {attr: "src", resType: "script"},
	"link": {attr: "href", resType: "stylesheet"},
	"video": {attr: "src", resType: "video"},
	"source": {attr: "src", resType: "video"},
}

func Extract(n *html.Node) []Tags {
	var result []Tags
	visit(n, &result)

	return result
}

func visit(n *html.Node, result *[]Tags) {
	if n.Type == html.ElementNode {
		if cfg, ok := tagAttrMap[n.Data]; ok {
			for _, attr := range n.Attr {
				if attr.Key == cfg.attr {
					*result = append(*result, Tags{
						Type: cfg.resType,
						URL: attr.Val,
					})
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, result)
	}
}
