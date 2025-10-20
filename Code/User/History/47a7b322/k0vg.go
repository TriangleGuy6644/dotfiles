package main

import (
	"fmt"
	"os/exec"
	"runtime"
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
func detectPackageManagers() []string {
	managers := []string{"apt", "dnf", "yum", "pacman", "zypper", "apk", "xbps-install", "nix-env", "emerge", "brew"}
	found := []string{}

	for _, mgr := range managers {
		if _, err := exec.LookPath(mgr); err == nil {
			found = append(found, mgr)
		}
	}
	return found
}

func preferredManager() string {
	found := detectPackageManagers()

	if len(found) == 0 {
		return "unknown"
	}

	if runtime.GOOS == "darwin" {
		return "brew"
	}

	for _, native := range []string{"apt", "dnf", "yum", "pacman", "zypper", "apk", "xbps-install", "nix-env", "emerge"} {
		for _, mgr := range found {
			if mgr == native {
				return mgr
			}
		}
	}

	return found[0]
}

func main() {
	fmt.Println(detectPM(), ", ", preferredManager())
}
