// https://play.golang.org/p/xfHwQ-9CVxx
// Difference between pointer type receiver method and named type receiver method
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

type Man struct {
	age int
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.X)
}

// Distance returns the distance traveled along the path
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			// fmt.Printf("%T", path[i-1])
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// Method with a pointer receiver
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (m Man) Display() {
	m.age += 105
	fmt.Println(m.age)
}

func DisplayPtr(m Man) {
	fmt.Println(m.age)
}

func main() {

	lawPtr := Man{10}

	lawPtr.Display()
	DisplayPtr(lawPtr)

	p := Point{1, 2}
	q := Point{4, 6}

	r := Point{1, 2}
	r.ScaleBy(2)

	fmt.Println(r)
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
	fmt.Println(perim.Distance())
}
