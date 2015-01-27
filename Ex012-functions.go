// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func main(){
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
}