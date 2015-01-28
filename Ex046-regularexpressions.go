// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "bytes"
import "fmt"
import "regexp"

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	//for other regexp tasks you need to Compile an optimized struct
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	//returns the match indexes
	fmt.Println(r.FindStringIndex("peach punch"))
	
}