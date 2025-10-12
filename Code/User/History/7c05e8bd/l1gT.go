package main

import (
	"fmt"
	"io/ioutil"
)

type contactInfo struct{
	Name string
	Email string
}

type purchaseInfo struct{
	Name string
	Price float32
	Amount int
}

func main(){

}

func loadJSON[T contactInfo | purchaseInfo](filepath string) []T{
	data, _ = ioutil
}