// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"
import "math"

//interfaces are collections of methods
type geometry interface {
	//method name and return type
	area() float64
	perim() float64
}

//structs are collections of variables
type square struct {
	//variable name(s) and type
	width, height float64
}

type circle struct {
	radius float64
}

//area method recieving a square
func (s square) area() float64 {
	return s.width * s.height
}

//perimeter method recieving a square
func (s square) perim() float64 {
	return 2*s.width + 2*s.height
}

//area method recieving a circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//perimeter method recieving a circle
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	s := square{width: 3, height: 4}
	c := circle{radius: 5}
	
	measure(s)
	measure(c)
}