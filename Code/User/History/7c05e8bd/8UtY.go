package main

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
	
}