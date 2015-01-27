// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	
	<-done
}