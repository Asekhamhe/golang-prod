package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello from the other side using goroutine!")
}

func main() {
	go hello()
	time.Sleep(1 * time.Second)
}
