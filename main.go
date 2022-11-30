// Objective.
// learn how to use go os to issue command line commands
// - example: run python script from same folder

package main

	

import (
	"log"
	"os/exec"
	)

	

func main() {

	arg1 := "python3"

	arg2 := "./helper.py"
    
	cmd := exec.Command(arg1, arg2)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
