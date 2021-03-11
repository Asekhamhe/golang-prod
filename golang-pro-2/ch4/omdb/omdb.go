package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// API is
const API = "http://www.omdbapi.com/?apikey=[KEY]"

// SearchResult is
type SearchResult struct {
	Search []*Search
}

// Search object is
type Search struct {
	Title  string
	Year   string
	ImdbID string `json:"imdbID"`
	Type   string
	Poster string
}

// SearchMovies is
func SearchMovies(terms []string) (*SearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	res, err := http.Get(API + "&s=" + q)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", res.Status)
	}

	var result SearchResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		res.Body.Close()
		return nil, err
	}
	res.Body.Close()
	return &result, nil
}
