package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String2() string {
	return fmt.Sprintf("$%.2f", d)
}

func (d dollars) String() string {
	return fmt.Sprintf("$%.3f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {

		fmt.Fprintf(w, "%s: %s: %s\n", item, price, price)
	}
}

func main() {

	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe(":8000", db))

}
