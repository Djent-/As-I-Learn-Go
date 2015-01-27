// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	
	//this mutex will synchronize access to state
	var mutex = &sync.Mutex{}
	
	var ops int64 = 0
	
	//start 100 goroutines that will repeatedly read the state
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				//pick a key to access
				key := rand.Intn(5)
				//Lock() the mutex to ensure exclusive access
				mutex.Lock()
				//read the value of the state
				total += state[key]
				//unlock the mutex
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				
				//explicitly yeild after each operation
				runtime.Gosched()
			}
		}()
	}
	
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	
	time.Sleep(time.Second)
	
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
	
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}