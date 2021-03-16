// Variadic function for max and min
package main

import "fmt"

func min(n ...int) int {
	min := 0
	for i, v := range n {
		if n[i] < n[len(n)-(len(n)-1)] {
			min = v
		}
	}
	return min
}

func max(n ...int) int {
	max := 0
	for i, v := range n {
		if n[i] > n[len(n)-(len(n)-1)] {
			max = v
		}
	}
	return max
}

func main() {
	min := min(2, 3, 5, 7, 10, 11)
	max := max(2, 3, 5, 7, 10, 11)
	fmt.Printf("The min is %d and the max is %d", min, max)
}
