package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

}

/* // declare functions
func sysUpd() {
	cmd := exec.Command("sudo", "pacman", "-Syu")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}


func pkgIns() {
	cmd := exec.Command("")
} */

func runCmd(command string) {
    parts := strings.Fields(command)
    if len(parts) == 0 {
        fmt.Println("No command provided")
        return
    }

    cmd := exec.Command(parts[0], parts[1:]...)
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Error:", err)
    }

    fmt.Print(string(output))
}
