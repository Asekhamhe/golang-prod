package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lorezi/golang-prod/golang-pro-2/ch4/omdb"
)

func main() {
	res, err := omdb.SearchMovies(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, movie := range res.Search {
		fmt.Printf("%s\t%s\t%s\t%s\n", movie.Title, movie.Year, movie.Type, movie.Poster)
	}
}
