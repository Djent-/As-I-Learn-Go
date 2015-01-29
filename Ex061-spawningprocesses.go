// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {
	//exec.Command creates an object to represent the external process
	dateCmd := exec.Command("date")
	
	//.Output collects the output of the command after it executes
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))
	
	//pipe data into the external process on its stdin and collect results
	grepCmd := exec.Command("grep", "hello")
	
	//explicitly grab input/output pipes
	//start the process
	//write some input
	//read output
	//exit process
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))
	
	//when spawning commands we need to provide an explicitly delineated
	//command and argument array
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}