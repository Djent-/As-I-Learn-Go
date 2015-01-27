// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"

//By default channels are unbuffered, meaning that they will 
//only accept sends (chan <-) if there is a corresponding
//receive (<- chan) ready to receive the sent value. Buffered 
//channels accept a limited number of values without a
//corresponding receiver for those values.

func main() {
	//channel of strings, buffering up to two values
	messages := make(chan string, 2)
	
	messages <- "buffered"
	messages <- "channel"
	
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}