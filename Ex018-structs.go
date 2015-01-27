// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	
	sp := &s
	//pointers are automatically dereferenced
	fmt.Println(sp.age)
	
	sp.age = 51
	fmt.Println(sp.age)
}