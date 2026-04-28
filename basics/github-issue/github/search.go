package github

import "fmt"

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := buildQuery(terms)

	fullURL := IssuesURL + "?q=" + q
	fmt.Println("DEBUG URL:", fullURL)

	res, err := fetch(fullURL)
	if err != nil {
		return nil, err
	}

	return decodeResponse(res)
}
