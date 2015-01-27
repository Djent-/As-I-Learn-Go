// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "time"
import "fmt"

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	
	//this will be the regulator in the limiting scheme
	limiter := time.Tick(time.Millisecond * 200)
	
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	
	//buffering the limiter channel will allow bursts of events
	burstyLimiter := make(chan time.Time, 3)
	
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()
	
	burstyRequests := make(chan int, 5)
	//the first 3 of the five will benifit from the burst
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}