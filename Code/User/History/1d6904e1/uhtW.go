package tools

import(
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex":{
		AuthToken: "123ABC",
		Username: "alex",
	},
	"arshan":{
		AuthToken: "gIsLove",
		Username: "arshan",
	},
	"gLover":{
		AuthToken: "gIsLife",
		Username: "glover",
	},
	"trippitroppi":{
		AuthToken: "tralalero",
		Username: "trippi",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex":{
		Coins: 100,
		Username: "alex",
	},
	"arshan":{
		Coins: 10000,
		Username: "arshan",
	},
	"gLover":{
		Coins: 10e+3,
		Username: "glover",
	},
	"trippitroppi":{
		Coins: 171539955,
		Username: "trippi",
	},
}

func ()