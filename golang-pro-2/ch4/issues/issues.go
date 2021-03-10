package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lorezi/golang-prod/golang-pro-2/ch4/github"
)

func main() {
	res, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", res.TotalCount)
	for _, v := range res.Items {
		fmt.Printf("#%-5d %9.9s %.55s \n", v.Number, v.User.Login, v.Title)
	}
}
