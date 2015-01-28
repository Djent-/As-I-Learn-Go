// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "flag"
import "fmt"

func main() {
	//declare a string flag 'word' with default value 'foo' and description
	//returns a string pointer
	wordPtr := flag.String("word", "foo", "a string")
	
	numPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	
	//it is possible to declare an option that uses an existing var
	//declared elsewhere in the program
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	
	//execute command line parsing
	flag.Parse()
	
	//* to dereference the pointers
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}