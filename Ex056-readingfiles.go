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
	dat, err := ioutil.ReadFile("Ex001-hello.go")
	check(err)
	fmt.Print(string(dat))
	
	f, err := os.Open("Ex001-hello.go")
	check(err)
	
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))
	
	//seek to a known location in a file and Read from there
	o2, err := f.Seek(6, 0)
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
	
	//the bufio package implements a buffered reader that may be useful
	//for both its efficieny with many small reads and because of the
	//additional reading methods it provides
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	
	f.Close()
}