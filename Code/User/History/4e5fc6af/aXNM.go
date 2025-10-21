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
	args := os.Args[
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
func div[G number](a, b G) G{
	return a/b
}
