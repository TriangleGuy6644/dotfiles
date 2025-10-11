package main

import(
	"fmt"
	"time"
	"math/rand"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main(){
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		dbCall(i)

	}
	
}

func dbCall(i int){
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
}