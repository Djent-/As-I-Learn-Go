// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"
import "time"

func main() {
	p := fmt.Println
	
	now := time.Now()
	p(now)
	
	//build a time struct
	then := time.Date(
		2009, 11 ,17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	
	//extract various components
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	
	p(then.Weekday())
	
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))
	
	//the Sub method returns a Duration
	diff := now.Sub(then)
	
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	
	//use Add to advance a time by a given duration
	p(then.Add(diff))
	p(then.Add(-diff))
}