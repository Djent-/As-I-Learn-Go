// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "strconv"
import "fmt"

func main() {
	//64 tells ParseFloat how many bits of precision
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	
	//0 means infer the base from the string
	//64 means the result needs to fit in 64 bits
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	
	//ParseInt will recognize hex-formatted numbers
	d, _ := strconv.ParseInt("ox1c8", 0, 64)
	fmt.Println(d)
	
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	
	//Atoi is a convenience function for basic base-10 int parsing
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	
	//Parse functions return an error on bad input
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}