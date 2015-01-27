// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "time"
import "fmt"

func main() {
	
	//timers provide a channel that will be notified when
	// the time is reached
	timer1 := time.NewTimer(time.Second * 2)
	
	//blocks on the timer's channel C until it sends a value
	<-timer1.C
	fmt.Println("Timer 1 expired")
	
	//timers can be cancelled before they expire
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}