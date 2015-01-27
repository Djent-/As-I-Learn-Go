// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"

//channels can be read-only or write-only

//only accepts a channel for sending values
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//takes one read-only channel and one write-only channel
//note the difference in placement of arrow

// "<-chan" read-only
// "chan<-" write-only
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}