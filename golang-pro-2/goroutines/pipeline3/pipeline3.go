package main

import "fmt"

// sender goroutine function
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

// sender and receiver goroutine function
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

// receiver only goroutine function
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	// make a channel for counter function
	naturals := make(chan int)
	// make a channel for squarer function
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)

}
