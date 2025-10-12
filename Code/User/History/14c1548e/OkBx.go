package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3}
	fmt.Println("Original slice:", mySlice)

	// Append individual elements
	mySlice = append(mySlice, 4, 5)
	fmt.Println("After appending individual elements:", mySlice)
}
