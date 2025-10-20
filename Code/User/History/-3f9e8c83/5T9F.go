package main

import (
	"fmt"
	"time"
)

func maing() {
	now := time.Now()
	for i := 0; i < 1000001; i++ {
		fmt.Println(i)
	}
	fmt.Println("Done! Elapsed: ", time.Since(now))
}
