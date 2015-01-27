// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"

func main() {
	// make(chan val-type)
	messages := make(chan string)
	
	//send a value into a channel
	//from a goroutine
	go func() {messages <- "ping"}()
	
	//recieves a value from a channel
	msg := <-messages
	fmt.Println(msg)
}

//When we run the program the "ping" message is successfully 
//passed from one goroutine to another via our channel.

//By default sends and receives block until both the sender 
//and receiver are ready. This property allowed us to wait at
//the end of our program for the "ping" message without 
//having to use any other synchronization.