package github

var CachedResult *IssuesSearchResult

func LoadData(query string) error {
	result, err := FetchIssue(query)
	if err != nil {
		return err
	}

	CachedResult = result
	return nil
}
