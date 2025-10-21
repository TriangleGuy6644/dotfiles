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
	args := os.Args
	//arg 0=filename, 1=operation, 2 and 3=numbers
	if args < 4{
		fmt.Println("please provide 2 numbers.")
		os.Exit(0)
	}
	switch args[1]{
	case "add":
		fmt.Println(add())
		
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
