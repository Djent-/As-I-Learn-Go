// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")
	go f("goroutine")
	
	//anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}