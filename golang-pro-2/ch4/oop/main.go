package main

import "github.com/lorezi/golang-prod/golang-pro-2/ch4/oop/person"

func main() {
	p1 := person.New()

	p1.Name = "Lawrence"
	p1.Age = 23
	p1.Address = "Silicon Valley USA"
	p1.Properties()

	p1.Code([]string{"Golang", "Nodejs"})

}
