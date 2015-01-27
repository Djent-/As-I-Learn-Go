// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"

func main() {
	
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	
	for {
		fmt.Println("loop")
		break
	}
}