// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {
	var ops int64 = 0
	
	reads := make(chan *readOp)
	writes := make(chan *writeOp)
	
	//this goroutine owns the state
	go func() {
		var state = make(map[int]int)
		for {
			select {
				case read := <-reads:
					read.resp <- state[read.key]
				case write := <-writes:
					state[write.key] = write.val
					write.resp <- true
			}
		}
	}()
	
	for r := 0; r < 100; r++ {
		go func() {
			for {
				//construct a read operation
				read := &readOp{
					key: rand.Intn(5),
					resp: make(chan int)}
				//send read operation through reads channel
				reads <- read
				//receive reply from state-owning goroutine
				<-read.resp
				//increment the operations counter
				atomic.AddInt64(&ops, 1)
			}
		}()
	}
	
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}
	
	time.Sleep(time.Second)
	
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
}