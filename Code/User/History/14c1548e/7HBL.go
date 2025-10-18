package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	sysUpd()
}

// declare functions
func sysUpd() {
	cmd := exec.Command("sudo", "pacman", "-Syu")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nik {
		fmt.Println("Error: ", err)
	}
}
