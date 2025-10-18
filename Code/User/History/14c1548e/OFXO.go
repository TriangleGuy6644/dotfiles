package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/creack/pty"
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
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error!\n", err)
	}
}

func runCmdIn