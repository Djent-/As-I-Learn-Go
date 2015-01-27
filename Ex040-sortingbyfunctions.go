// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "sort"
import "fmt"

//alias for []string
type ByLength []string

//implement these functions for our custom type so we can
//use the sort package's generic Sort function
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	//this type of assignment worksin Lua too
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}