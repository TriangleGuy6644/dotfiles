package main

import (
	"fmt"
	"math/rand"
	"time"
)

var max_monster_price float32 = 4

func main() {
	var monsterChannel = make(chan string)
	var websites = []string{"walmart.com", "dollarama.com", "costco.com"}
	for i := range websites {
		go checkMonsterPrices(websites[i], monsterChannel)
	}
	sendMessage(monsterChannel)
}

func checkMonsterPrices(website string, monsterChannel chan string){
	for {
		time.Sleep(time.Second*1)
		var monsterPrice = rand.Float32()*20
		if monsterPrice <= max_monster_price{
			monsterChannel <- website
			break
		}
	}
}

func sendMessage(monsterChannel chan string){
	fmt.Println("\nFound a deal on monster at %s", <- monsterChannel)
}