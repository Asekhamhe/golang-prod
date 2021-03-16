// Variadic version of strings.Join
package main

import "fmt"

func characters(s ...string) string {
	word := ""
	// var word string
	for _, v := range s {
		word += v
	}
	return word
}

func main() {
	s := []string{"a", "b", "c", "d", "e"}
	word := characters(s...)
	fmt.Println(word)
}
