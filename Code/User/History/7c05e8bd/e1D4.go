package main

import (
	"fmt"
)

func main() {
	//intslice
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumIntSlice(intSlice))

	//float32slice
	var float32slice = []float32{1, 2, 3}
	fmt.Println(sumFloat32Slice(float32slice))

	//float64slice
	var float64slice = []float32{1, 2, 3}
	fmt.Println(sumFloat64Slice(float64slice))
}

func sumIntSlice(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

func sumFloat32Slice(slice []float32) float32 {
	var sum float32
	for _, v := range slice {
		sum += v
	}
	return sum
}

func sumFloat64Slice(slice []float64) float64 {
	var sum float64
	for _, v := range slice {
		sum += v
	}
	return sum
}
