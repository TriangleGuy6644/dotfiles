package main


func main(){
	var c = make(chan int)
	c <- 1
}