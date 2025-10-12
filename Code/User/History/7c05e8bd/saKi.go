package main
import("fmt")

func main(){
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice(intSlice))
}

func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _, v:=range slice{
		sum += v
	}
	return sum
}