package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	runCmd("sudo pacman -Syu")
}

func runCmd(command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		fmt.Println("no command provided.")
		return
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	// cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(output))

}
