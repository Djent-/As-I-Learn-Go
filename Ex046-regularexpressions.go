// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "bytes"
import "fmt"
import "regexp"

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	//for other regexp tasks you need to Compile an optimized struct
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	//returns the match indexes
	fmt.Println(r.FindStringIndex("peach punch"))
	//will return information for p([a-z]+)ch and ([a-z]+)
	fmt.Println(r.FindStringSubmatch("peach punch"))
	//return information about the indexes of matches and submatches
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	//apply to all matches in the input
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	//we can provide []byte arguments and drop String from function name
	fmt.Println(r.Match([]byte("peach")))
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	
	in := []byte("a peach")
	//the Func variant allows you to transform matched text with a 
	//given function
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}