package main

import "fmt"

func main(){
	tasks := []string{}
	fmt.Println("The current tasks are: ", tasks)
	fmt.Println("Add a new task: ")
	fmt.Scan(&tasks)
	fmt.Println("The new current tasks are: ", tasks)
}
