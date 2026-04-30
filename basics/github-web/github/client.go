package github

import (
	"encoding/json"
	"net/http"
)

const IssueURL = "https://api.github.com/search/issue?q="

func FetchIssue(query string) (*IssuesSearchResult, error) {
	res, err := http.Get(IssueURL + query)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var result IssuesSearchResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
