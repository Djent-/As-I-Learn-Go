// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"
import "time"

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)
	
	//there is no UnixMillis so you need to divide
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)
	
	//can convert integer seconds or nanoseconds since the epoch
	//into the corresponding time
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}