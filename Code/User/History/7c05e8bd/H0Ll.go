package main

import (
	"fmt"
	"math/rand"
	"time"
)

var max_monster_price float32 = 5
var max_celsius_price float32 = 4

func main() {
	var monsterChannel = make(chan string)
	var celciusChannel = make(chan string)
	var websites = []string{"walmart.com", "dollarama.com", "costco.com"}
	for i := range websites {
		go checkMonsterPrices(websites[i], monsterChannel)
		go checkCelciusPrices(websites[i], celciusChannel)
	}
	sendMessage(monsterChannel, celciusChannel)
}

func checkCelciusPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var celcius_price = rand.Float32() * 20
		if celcius_price < max_celsius_price {
			c <- website
			break
		}
	}

}

func checkMonsterPrices(website string, monsterChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var monsterPrice = rand.Float32() * 20
		if monsterPrice <= max_monster_price {
			monsterChannel <- website
			break
		}
	}
}

func sendMessage(monsterChannel chan string, celciusChannel chan string) {
	select {
	case website := <-monsterChannel:
		fmt.Printf("\nText Sent: found deal on monster at %v", website)
		
	}
}
