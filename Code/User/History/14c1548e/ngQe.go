package main

import(
	"fmt"
)

func main(){
	tasks := []string{}
	fmt.Println("The current tasks are: ", tasks)
	var newTask string
	fmt.Print("Add a new task: ")
	fmt.Scan(&newTask)
}
