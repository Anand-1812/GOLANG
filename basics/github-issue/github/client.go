package github

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func buildQuery(terms []string) string {
	return url.QueryEscape(strings.Join(terms, " "))
}

func fetch(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()

		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}

	return resp, nil
}
