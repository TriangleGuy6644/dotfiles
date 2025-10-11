package main

import "fmt"

func main() {
	var c = make(chan int)
	c <- 1
	var i = <-c
	fmt.Println(i)
	math.
}
