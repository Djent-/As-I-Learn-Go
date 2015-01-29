// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "syscall"
import "os"
import "os/exec"

func main() {
	//Go requires an absolute path to the binary we want to execute
	//we'll use exec.LookPath to find it
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	
	//Exec requires arguments in slice form
	args := []string{"ls", "-a", "-l", "-h"}
	
	//Exec also needs a set of environment variables to use
	//we provide our current environment
	env := os.Environ()
	
	//if this call is successful, the execution of our process will end
	//here and be replaced by the /bin/ls -a -l -h process
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}