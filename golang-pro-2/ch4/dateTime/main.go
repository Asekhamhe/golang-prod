package main

import (
	"fmt"
	"time"
)

func main() {
	start, _ := time.Parse("02-01-2006", time.Now().AddDate(0, 0, -30).Format("02-01-2006"))

	fmt.Println(start.Format("02-01-2006"))

	dates := []string{
		time.Date(2021, 01, 10, 0, 0, 0, 0, time.Now().UTC().Location()).Format("02-01-2006"),
		time.Date(2021, 03, 12, 0, 0, 0, 0, time.Now().UTC().Location()).Format("02-01-2006"),
		time.Date(2021, 02, 11, 0, 0, 0, 0, time.Now().UTC().Location()).Format("02-01-2006"),
		time.Date(2021, 03, 19, 0, 0, 0, 0, time.Local).Format("02-01-2006"),
	}
	for _, v := range dates {
		if v, _ := time.Parse("02-01-2006", v); v.After(start) {
			fmt.Println(v.Format("02-01-2006"))
		}
	}

}
