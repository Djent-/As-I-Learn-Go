// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//error helper
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))
	
	f, err := os.Open("/tmp/dat")
	check(err)
	
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1m string(b1))
	
	//seek to a known location in a file and Read from there
	os, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
	
	//reads like the ones above can be more robustly implemented with
	//ReadAtLeast
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	
	//there's no built in rewind, but Seek(0, 0) accomplishes this
	_, err = f.Seek(0, 0)
	check(err)
	
	
}