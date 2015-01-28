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
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	
	//the json package can encode your custom data types
	//it will only include exported fields in the encoded output
	//by default will use those names as json keys
	res1D := &Response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	
	res2D := &Response2{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
	
	//decoding JSON data into Go values
	//example for a generic data structure
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	
	//we need to provide a variable where the JSON package can put the
	//encoded data. this will hold a map of strings to arbitrary data types
	var dat map[string]interface{}
	
	//the actual decoding and error checking
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	
	//in order to use the values in the decoded map, we'll need to cast them
	//to their appropriate type. here we cast num to float64
	num := dat["num"].(float64)
	fmt.Println(num)
	
	//accessing nested data requires a series of casts
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)
	
	//decode JSON into custom data types
	//advantage: type safety and eliminating the need to type assertions
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
	
	//we have used bytes and strings as intermediates between data and json
	//we can also stream JSON encodings directly to os.Writers like
	//os.Stdout or even HTTP response bodies
	enc := json.NewEncoder(os.Stdout)
	d:= map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}