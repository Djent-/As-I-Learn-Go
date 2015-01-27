// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	// sort of like the forward slash in perl I suppose
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	
	zeroval(i)
	fmt.Println("zeroval:", i)
	
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	
	fmt.Println("pointer:", &i)
}