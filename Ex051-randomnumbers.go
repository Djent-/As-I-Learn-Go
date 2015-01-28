// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"
import "math/rand"

func main() {
	//rand.Intn returns a random int n, 0 <= n <= 100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
	
	//rand.Float64 returns a float64 f, 0.0 <= f < 1.0
	fmt.Println(rand.Float64())
	
	//this can be used to generate random floats in other ranges
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()
	
	//to make the pseudorandom number generator deterministic
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)
	
	//call the resulting rand.Source just like the functions on rand
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()
}