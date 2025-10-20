package main

import (
	"fmt"
	"os/exec"
)

func detectPM() string {
	pms := []string{"apt", "pacman", "dnf", "yum", "zypper", "xbps-install", "brew"}
	for _, mgr := range pms {
		if _, err := exec.LookPath(mgr); err == nil {
			return mgr
		}
	}
	return "unknown"
}

func main() {
	fmt.Println(detectPM())
}
