package main

import "fmt"

func main() {

	ch := make(chan string)

	go func(ch chan string) {
		ch <- "hello, world 1"
	}(ch)

	go func(ch chan string) {
		ch <- "hello, world 2"
	}(ch)

	go func(ch chan string) {
		ch <- "hello, world 3"
	}(ch)

	a, b, c := <-ch, <-ch, <-ch

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
