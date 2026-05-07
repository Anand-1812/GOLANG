package internals

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// fetches html from live url
func FromURL(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching %s: %w", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetching %s: status %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing HTML from %s: %w", url, err)
	}

	return doc, nil
}

func FromStdin() (*html.Node, error) {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("parsing HTML from stdin: %w", err)
	}
	return doc, nil
}

func FromReader(r io.Reader) (*html.Node, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, fmt.Errorf("parsing HTML from reader: %w", err)
	}
	return doc, nil
}
