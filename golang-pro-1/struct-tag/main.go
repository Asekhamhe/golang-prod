package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Person contains name and email
type Person struct {
	Name  string `json: "name"`
	Email string `json:"email"`
}

func main() {
	p := Person{"Lawrence", "law@example.com"}
	b, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

}
