package main
import("fmt")

func main(){
	var intSlice = []int{}
	fmt.Println(isEmpty[int](intSlice))

	var float32Slice = []float32{1,2,3}
	fmt.Println(isEmpty[float32](float32Slice))
}

func isEmpty[T any](slice []T) bool{
	return len(slice)==0
}