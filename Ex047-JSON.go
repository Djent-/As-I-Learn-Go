// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "encoding/json"
import "fmt"
import "os"

//we'll use these two structs to demonstrate encoding and decoding
//custom types
type Response1 struct {
	Page int
	Fruits []string
}
type Response2 struct {
	//graves? I already don't like this
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	//encoding basic types to JSON strings
	//here are some examples for atomic values
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))
	
	//here are some for slices and maps which encode to json arrays
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(sclD)
	fmt.Println(string(slcB))
	
	//the json package can encode your custom data types
	//it will only include exported fields in the encoded output
	//by default will use those names as json keys
	res1D := &Response1{
		page: 1,
		Fruits: []string{"apple", "peach", "pear"}
	}
	res1B, _ := json.Marshal(res1D)
	
	res2D := &Response2{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
	
	
	
	
	
	
	
	
	
	
	
	
}