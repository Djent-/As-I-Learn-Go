// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//wrapping the unbuffered os.Stdin with a buffered scanner gives us a
	//convenient Scan method that advanes the scanner to the next token;
	//which is the next line in the default scanner
	scanner := bufio.NewScanner(os.Stdin)
	
	//Text returns the current token, here the next line from the input
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}