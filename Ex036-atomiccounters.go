// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {
	//unsigned 64 bit integer
	var ops uint64 = 0
	
	//start 50 goroutines that increment the counter
	for i:= 0; i < 50; i++ {
		go func() {
			for {
				//to atomically increment the counter, we use AddUint64
				//giving it the memory address of our ops counter
				atomic.AddUint64(&ops, 1)
				
				//allow other goroutines to proceed
				runtime.Gosched()
			}
		}()
	}
	
	time.Sleep(time.Second)
	
	//in order to safely use the counter while it's being updated by other
	//goroutines, we extract a copy of the current value into opsFinal
	//via LoadUint64
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}