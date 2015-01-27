// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"

//non-blocking works in the following way:
//receiving only happens if there is a goroutine
	//waiting for the current goroutine to pick up
	//the message
//sending works the same way but from the other side

func main() {
	messages := make(chan string)
	signals := make(chan bool)
	
	select {
		case msg := <-messages:
			fmt.Println("received message", msg)
		default:
			fmt.Println("no message received")
	}
	
	msg := "hi"
	select {
		case messages <- msg:
			fmt.Println("sent message", msg)
		default:
			fmt.Println("no message sent")
	}
	
	select {
		case msg := <-messages:
			fmt.Println("received message", msg)
		case sig := <-signals:
			fmt.Println("received signal", sig)
		default:
			fmt.Println("no activity")
	}
}