package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	for i := 0; i < 1000000001; i++ {
		fmt.Println(i)
	}
	fmt.Println("Done! Elapsed: ", time.Since(now))
}
