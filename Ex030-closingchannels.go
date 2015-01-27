// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/


package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	
	go func() {
		for {
			j, more := <- jobs
			//more will be false if jobs has been closed
			//and all values in the channel have already been
			//received.
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	
	<-done
}