// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "os"

func main() {
	panic("a problem")
	
	_, err := os.Create("tmp/file")
	if err != nil {
		panic(err)
	}
}