// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic (e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("test.txt", d1, 0644)
	check(err)
	
	//granular writes
	f, err := os.Create("test1.txt")
	check(err)
	
	defer f.Close()
	
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)
	
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)
	
	//issue a Sync to flush writes to stable storage
	f.Sync()
	
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	
	//use Flush to ensure all buffered operations have been applied
	w.Flush()
}