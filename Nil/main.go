package main

import "fmt"

func main() {
	var i interface{}

	fmt.Println(i == nil)
	fmt.Printf("%T, %v\n", i, i)

	var s *string
	fmt.Println("s === nil:", s == nil)

	i = s
	fmt.Println("i == nil:", i == nil)
	fmt.Printf("%T, %v\n", i, i)
}
