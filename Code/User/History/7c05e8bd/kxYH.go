package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func main() {
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("")
}

func loadJSON[T contactInfo | purchaseInfo](filepath string) []T {
	data, _ = ioutil.ReadFile(filepath)
	var loaded = []T{}
	json.Unmarshal(data, &loaded)
	return loaded
}
