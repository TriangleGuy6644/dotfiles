package main

import(
	"fmt"
	"time"
)

func main(){
	tasks := []string{}
	fmt.Println("The current tasks are: ", tasks)
	fmt.Println("Add a new task: ")
	fmt.Scan(&tasks)
	time.Sleep(time.Second*2)
	fmt.Println("The new current tasks are: ", tasks)
}
