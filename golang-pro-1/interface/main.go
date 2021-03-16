package main

import (
	"fmt"
	"time"
)

// Entry interface
type Entry interface {
	Title() string
}

// Book struct
type Book struct {
	Name      string
	Author    string
	Published time.Time
}

// Title function returns book details
func (b Book) Title() string {
	return fmt.Sprintf("%s by %s (%s)", b.Name, b.Author, b.Published.Format("Ja\n 2006"))
}

// Movie struct
type Movie struct {
	Name     string
	Director string
	Year     int
}

// Title function is
func (m Movie) Title() string {
	return fmt.Sprintf("%s (%d)", m.Name, m.Year)
}

// Display using interface and polymorphism
func Display(e Entry) string {
	return e.Title()
}

func main() {
	b := Book{
		Name: "John Doe", Author: "David Mcloskey", Published: time.Date(2001, time.May, 22, 0, 0, 0, 0, time.UTC),
	}
	m := Movie{Name: "The Godfather", Director: "Francis Ford Coppola", Year: 1972}

	fmt.Println(Display(b))
	fmt.Println(Display(m))
}

// interface
// https://play.golang.org/p/T0GnoRkicOz
