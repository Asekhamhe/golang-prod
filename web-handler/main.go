package main

import (
	"log"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, req *http.Request) {
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("hello, world")
	}()
	return
}

func main() {
	http.HandleFunc("/", home)
	log.Println("Running on :8000...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
