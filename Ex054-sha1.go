// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "crypto/sha1"
import "fmt"

func main() {
	s := "sha1 this string"
	
	//start with a new hash
	h := sha1.New()
	
	//write expects bytes
	//if you have a string s, use []bytes(s) to coerce it to bytes
	h.Write([]byte(s))
	
	//this gets the finalized hash result as a byte slice
	//the argument to Sum can be used to append to an existing
	//byte slice. it isn't usually needed
	bs := h.Sum(nil)
	
	fmt.Println(s)
	//use the %x format verb to convert a hash result to a hex string
	fmt.Printf("%x\n", bs)
}