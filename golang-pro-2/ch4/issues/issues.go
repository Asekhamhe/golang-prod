package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lorezi/golang-prod/golang-pro-2/ch4/github"
)

func main() {
	month, _ := time.Parse("02-01-2006", time.Now().AddDate(0, 0, -30).Format("02-01-2006"))
	year, _ := time.Parse("02-01-2006", time.Now().AddDate(-1, 0, 0).Format("02-01-2006"))
	res, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range res.Items {
		// report the results in age categories of less than a month
		if d, _ := time.Parse("02-01-2006", v.CreatedAt.Format("02-01-2006")); d.After(month) {
			fmt.Printf("%s #%-5d %9.9s %.55s\n", v.CreatedAt.Format("02-01-2006"), v.Number, v.User.Login, v.Title)
		}

		// report the result in age categories of less than a year

		if y, _ := time.Parse("02-01-2006", v.CreatedAt.Format("02-01-2006")); y.After(year) {
			fmt.Printf("%s #%-5d %9.9s %.55s\n", v.CreatedAt.Format("02-01-2006"), v.Number, v.User.Login, v.Title)
		}

	}
}
