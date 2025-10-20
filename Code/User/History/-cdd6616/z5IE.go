package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("no args")
		os.Exit(0)
	}
	fmt.Printf("Arg 0: %v\nArg 1: %v\n Arg 2: %v\n Arg 3: %v\n", args[0], args[1], args[2], args[3])
	fmt.Println("All args: ", args)
}
