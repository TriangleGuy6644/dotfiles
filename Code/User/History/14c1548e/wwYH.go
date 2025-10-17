package main

import (
	"fmt"
	"os/exec"
)

func main() {
	sysUpd()
}

// declare functions
func sysUpd() {
	cmd := exec.Command("eza", "--icons")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(string(output))
}
