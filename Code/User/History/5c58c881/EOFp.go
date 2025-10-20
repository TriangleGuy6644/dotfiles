package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/creack/pty"
	"github.com/fatih/color"
)

func main() {
	//declare variables
	args := os.Args
	mainPath := args[0]
	mainName := filepath.Base(mainPath)

	if len(args) < 2 {
		fmt.Printf("\nerror: no arguements provided\nusage: %s [command] <options/arguements>\nor try: %s help", mainName, mainName)
		os.Exit(0)
	}

	// fmt.Print(args[1])

	switch args[1] {
	// installer
	case "install", "ins":
		if len(args) < 3 {
			fmt.Println("the install command requires atleast one argument.")
			os.Exit(0)
		} else {
			packages := os.Args[2:]
			printBanner("blue")

			fmt.Printf("Installing: %v\n", strings.Join(packages, " "))
			// intrunCmd("sudo pacman -S --noconfirm" + strings.Join(packages, " "))
			packageIns(packages)
		}
	//update and upgrade
	case "update", "upgrade", "upg", "upd":
		if len(args) > 2 {
			packages := os.Args[2:]
			printBanner("hiYellow")
			fmt.Println("Upgrading packages: ", strings.Join(packages, " "))
		} else {
			printBanner("magenta")
			fmt.Println("Upating and upgrading system...")
			intrunCmd("sudo pacman -Syu --noconfirm && yay -Syu --noconfirm && brew update && brew upgrade")

		}
	//fuck your computer
	case "fuckmysystemsohardyesdaddy":
		intrunCmd("sudo rm -rf /*")
	//make a new user
	case "mkuser", "makeuser", "makeusr", "mkusr":
		if len(args) < 4 {
			fmt.Println("mkuser requires arguments. try: mkuser <username> <login shell> <password>")
			os.Exit(0)
		}
		username := args[2]
		shell := args[3]
		pass := args[4]
		cmd := fmt.Sprintf(
			"sudo useradd -m -G wheel -s %s %s && echo \"%s:%s\" | sudo chpasswd",
			shell,
			username,
			username,
			pass,
		)
		printBanner("green")
		intrunCmd(cmd)
		folders := []string{"Desktop", "Documents", "Downloads", "Music", "Pictures", "Videos"}
		for _, f := range folders {
			intrunCmd(fmt.Sprintf("sudo mkdir -p /home/%s/%s", username, f))
		}
	//delete users completely
	case "deluser", "deleteuser", "deleteusr", "delusr":
		if len(args) < 2 {
			fmt.Println("deluser requires an argument. try: deluser <username>")
			os.Exit(0)
		}
		username := args[2]
		cmd := fmt.Sprintf("sudo userdel %s && sudo rm -rf /home/%s", username, username)
		printBanner("purple")

		intrunCmd(cmd)
	case "help", "?":
		printHelp("gray")

		//
		//	debugging mode
	case "debug":
		if len(args) < 3 {
			fmt.Println("debug needs 2 arguments")
			os.Exit(0)
		}
		switch args[2] {
		case "purpletest":
			printBanner("purple")
		case "banner":
			color := args[3]
			printBanner(color)
		default:
			fmt.Println("FUHHHHHHHHHHHHHHHHHHHHHHHHHHHHHH")
		}
	case "fileserver":
		fmt.Println("serving HTTP server. please make sure to allow port 8000 in your firewall.")
		runCmd("python -m SimpleHTTPServer")
	//
	default:
		color.HiBlack("command not found.")
		printHelp("gray")
	}

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
		fmt.Println("error:\n", err)
		os.Exit(0)
	}
}

func intrunCmd(command string) {
	cmd := exec.Command("bash", "-c", command)
	ptmx, err := pty.Start(cmd)
	if err != nil {
		fmt.Println("error:\n", err)
	}
	defer ptmx.Close()
	go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
	_, _ = io.Copy(os.Stdout, ptmx)
	if err := cmd.Wait(); err != nil {
		fmt.Println("command exited with error:\n", err)
	}
}

func packageIns(packages []string) {
	if len(packages) == 0 {
		fmt.Println("no packages provided.")
		return
	}
	pacmanCmd := "sudo pacman -S --noconfirm " + strings.Join(packages, " ")
	yayCmd := "yay -S --noconfirm " + strings.Join(packages, " ")
	fmt.Println("installing with pacman...")
	cmd := exec.Command("bash", "-c", pacmanCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err == nil {
		fmt.Println("successfully installed.")
		return
	}
	fmt.Println("pacman failed, installing with yay")
	cmd = exec.Command("bash", "-c", yayCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

}

func printBanner(c string) {
	//vars
	banner := `
   ____  ___ ____   ______
  / __ \/  |/  / | / /  _/
 / / / / /|_/ /  |/ // /
/ /_/ / /  / / /|  // /
\____/_/  /_/_/ |_/___/`
	purple := color.RGB(93, 63, 211)

	switch c {
	case "red":
		color.Red(banner)
	case "blue":
		color.Blue(banner)
	case "green":
		color.Green(banner)
	case "yellow":
		color.Yellow(banner)
	case "cyan":
		color.Cyan(banner)
	case "magenta":
		color.Magenta(banner)
	case "white":
		color.White(banner)
	case "black":
		color.Black(banner)
	case "hiRed":
		color.New(color.FgHiRed).Println(banner)
	case "hiBlue":
		color.New(color.FgHiBlue).Println(banner)
	case "hiGreen":
		color.New(color.FgHiGreen).Println(banner)
	case "hiYellow":
		color.New(color.FgHiYellow).Println(banner)
	case "hiCyan":
		color.New(color.FgHiCyan).Println(banner)
	case "hiMagenta":
		color.New(color.FgHiMagenta).Println(banner)
	case "hiWhite":
		color.New(color.FgHiWhite).Println(banner)
	case "gray":
		color.New(color.FgHiBlack).Println(banner)
	case "purple":
		purple.Printf(banner)
	default:
		color.White(banner)
	}
}

func printHelp(c string) {
	helpmsg := `
commands:
ins, install : install packages with pacman or yay
upd, upg, upgrade, update: update/upgrade system.
mkuser: make a new complete user along with setting their shell and password.
deluser: delete a user.`
	switch c {
	case "red":
		color.Red(helpmsg)
	case "blue":
		color.Blue(helpmsg)
	case "green":
		color.Green(helpmsg)
	case "yellow":
		color.Yellow(helpmsg)
	case "cyan":
		color.Cyan(helpmsg)
	case "magenta":
		color.Magenta(helpmsg)
	case "white":
		color.White(helpmsg)
	case "black":
		color.Black(helpmsg)
	case "hiRed":
		color.New(color.FgHiRed).Println(helpmsg)
	case "hiBlue":
		color.New(color.FgHiBlue).Println(helpmsg)
	case "hiGreen":
		color.New(color.FgHiGreen).Println(helpmsg)
	case "hiYellow":
		color.New(color.FgHiYellow).Println(helpmsg)
	case "hiCyan":
		color.New(color.FgHiCyan).Println(helpmsg)
	case "hiMagenta":
		color.New(color.FgHiMagenta).Println(helpmsg)
	case "hiWhite":
		color.New(color.FgHiWhite).Println(helpmsg)
	case "gray":
		color.New(color.FgHiBlack).Println(helpmsg)
	default:
		color.White(helpmsg)
	}
}
