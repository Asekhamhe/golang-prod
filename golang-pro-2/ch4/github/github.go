// Package github provides a GO API for the Github issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// IssuesURL is
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult is
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue is
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAr time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// User is
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the Github issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	res, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close res.Body on all execution paths
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", res.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		res.Body.Close()
		return nil, err
	}

	res.Body.Close()
	return &result, nil
}
