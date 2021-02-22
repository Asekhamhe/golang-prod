package main

import "fmt"

// Rectangle struct
type Rectangle struct {
	Length  int
	Breadth int
}

// Square struct
type Square struct {
	Length int
}

// Triangle struct
type Triangle struct {
	Length int
	Height int
}

// Perimeter interface function
func (r Rectangle) Perimeter() int {
	return r.Breadth * r.Length
}

// Perimeter interface function
func (s Square) Perimeter() int {
	return s.Length * s.Length
}

// Perimeter interface function
func (t Triangle) Perimeter() int {
	return (t.Height * t.Length) / 2
}

// Shapes interface
type Shapes interface {
	Perimeter() int
}

// Display function
func Display(s Shapes) int {
	return s.Perimeter()
}

func main() {
	r := Rectangle{10, 3}
	s := Square{10}
	t := Triangle{2, 5}

	fmt.Println(Display(r))
	fmt.Println(Display(s))
	fmt.Println(Display(t))

}
