package main

import (
	"os/exec"
)

func main() {

}

// declare functions
func sysUpd() {
	cmd := exec.Command("sudo", "pacman", "-syu")
}
