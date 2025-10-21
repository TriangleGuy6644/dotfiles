package main
import(
	"math/rand"
	"fmt"
	"os"
)

type number interface{
	int | int64 | float32 | float64
}

func main(){

}


func add[T number](a, b T) T {
	return a + b
}
func sub[T number](a, b T) T{
	return a - b
}
func mult[T number](a, b T) T{
	return a * b
}
