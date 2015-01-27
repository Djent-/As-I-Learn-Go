// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)
	
	const n = 500000000
	
	const d = 3e20 / n
	fmt.Println(d)
	
	fmt.Println(int64(d))
	
	fmt.Println(math.Sin(n))
}