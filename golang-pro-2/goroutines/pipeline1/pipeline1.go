package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; x <= 10; x++ {
			// sending values to the channels
			naturals <- x
		}
		close(naturals)
	}()

	// square
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}
