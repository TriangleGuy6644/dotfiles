package main

import (
	"fmt"
	"os"

)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("First Arguement: ", os.Args[1])
	}
}



/*
func printBanner() {
	color.Red(`
   ____  __  ____   ______
  / __ \/  |/  / | / /  _/
 / / / / /|_/ /  |/ // /  
/ /_/ / /  / / /|  // /   
\____/_/  /_/_/ |_/___/   
                          
	`)
}
*/