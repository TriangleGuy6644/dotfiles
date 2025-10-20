package main

import (
	"fmt"
	"os/exec"
)

func detectPM() []string {
	managers := []string{"apt", "dnf", "yum", "pacman", "zypper", "apk", "xbps-install", "nix-env", "emerge", "brew"}
	found := make(map[string]bool)
	list := []string{}

	for _, mgr := range managers {
		if _, err := exec.LookPath(mgr); err == nil {
			if !found[mgr] {
				found[mgr] = true
				list = append(list, mgr)
			}
		}
	}
	return list
}

func main() {
	fmt.Println(detectPM())
}
