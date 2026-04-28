package github

import (
	"encoding/json"
	"net/http"
	"fmt"
)

func decodeResponse(resp *http.Response) (*IssuesSearchResult, error) {
	defer resp.Body.Close()

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	
	fmt.Println("DEBUG TotalCount:", result.TotalCount)
	fmt.Println("DEBUG Items len:", len(result.Items))

	return &result, nil
}
