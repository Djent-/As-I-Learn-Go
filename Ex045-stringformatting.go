// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"
//import "os"

type point struct {
	x, y int
}

func main() {
	//print value
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	
	//%+v will include the struct's field names
	fmt.Printf("%+v\n", p)
	
	//prints Go syntax representation of the value (source code)
	fmt.Printf("%#v\n", p)
	
	var pf = fmt.Printf
	
	//prints the type of the value
	pf("%T\n", p)
	
	//formatting booleans
	pf("%t\n", true)
	
	//formatting integers
	pf("%d\n", 123)
	
	//binary representation
	pf("%b\n", 14)
	
	//corresponding character
	pf("%c\n", 33)
	
	//hex encoding
	pf("%x\n", 456)
	
	//floats
	pf("%f\n", 78.9)
	
	//...
}