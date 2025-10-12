package main
import("fmt")

func main(){
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumIntSlice(intSlice))
}

func sumSlice[T ]