// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	
	for w := 1; w <= 3; w++ {
		//create a worker goroutine
		go worker(w, jobs, results)
	}
	
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	
	for a := 1; a <= 9; a++ {
		<-results
	}
}