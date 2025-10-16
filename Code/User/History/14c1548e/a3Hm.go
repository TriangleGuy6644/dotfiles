package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("arguements: ", os.Args)
	fmt.Formatter
}

func printBanner() {
	color.Red(`
   ____  __  ____   ______
  / __ \/  |/  / | / /  _/
 / / / / /|_/ /  |/ // /  
/ /_/ / /  / / /|  // /   
\____/_/  /_/_/ |_/___/   
                          
	`)
}
