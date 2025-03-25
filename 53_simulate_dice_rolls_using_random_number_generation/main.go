package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	fmt.Printf("Role The Dice 10 Time . Let's Go !\n")
	for i:=0 ; i<10;i++ {
		fmt.Printf("Dice %d : %d\n",i+1,r.Intn(6)+1)
		// Sleep 1 second
		time.Sleep(time.Second)
	}
}