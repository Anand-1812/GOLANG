package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const IssueURL = "https://api.github.com/search/issues?q="

func FetchIssue(query string) (*IssuesSearchResult, error) {
	res, err := http.Get(IssueURL + url.QueryEscape(query))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github search failed: %s", res.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
