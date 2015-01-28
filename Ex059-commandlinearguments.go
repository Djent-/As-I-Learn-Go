// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "os"
import "fmt"

func main() {
	//os.Args proves access to raw command-line arguments
	//note that the first value in this slice is the path to the program
	//and os.Args[1:] holds the arguments to the program
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	
	arg := os.Args[3]
	
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}