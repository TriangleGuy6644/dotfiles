package main

import(
	"fmt"
)

func main(){
	var intSlice = []int{1,2,3}
	fmt.Println(sumSlice(intSlice))
}

func sumSlice(slice []int){
	var sum int
	for _, v := range slice{
		sum += v
	}
	return sum
}