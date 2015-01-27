// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	
	kvs := map[string]string{"a": "apple", "b": "banana"}
	// this seems to be a bit like lua
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}