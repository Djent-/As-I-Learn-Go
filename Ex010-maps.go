// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"

// To make an empty map, use the built-in make:
	// make(map[key-type]val-type)
	
func main() {
	m := make(map[string]int)
	
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)
	
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))
	
	delete(m, "k2")
	fmt.Println("map:", m)
	
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}