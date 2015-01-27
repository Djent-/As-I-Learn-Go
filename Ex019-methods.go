// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"

type rect struct {
	width, height int
}

// takes a rect pointer named r
// method name is area
// returns an int
func (r *rect) area() int {
	return r.width * r.height
}

// this method takes an entire rect
// "value receiver"
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
	
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}