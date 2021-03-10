package person

import "fmt"

// struct object

// Person is
type Person struct {
	Name    string // name property
	Age     int    // age property
	Address string // address property
}

// constructor

// New is
func New() *Person {
	return &Person{}
}

// methods

// Drive is
func (p Person) Drive(c string) {
	fmt.Println("Drive method")
}

// Code is
func (p Person) Code(lang []string) {
	for _, v := range lang {
		fmt.Printf("%v programming language\n", v)
	}

}

// Sports is
func (p Person) Sports(sport string) {
	fmt.Println("Sports method")
}

// Properties is
func (p Person) Properties() {
	fmt.Println(p.Name, "\t", p.Age, "\t", p.Address)
}
