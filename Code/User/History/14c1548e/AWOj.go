package main

import (
	"fmt"
	"os/exec"
)

func main() {

}

// declare functions
func sysUpd() {
	cmd := exec.Command("ls", "-la")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(string(output))
}
