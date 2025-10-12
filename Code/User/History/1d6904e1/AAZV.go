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