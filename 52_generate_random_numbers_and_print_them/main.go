package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomInt := r.Intn(10000000000)
	fmt.Printf("A Random Number (Int) : %d \n",randomInt)

	randomFloat := r.Float64()
	fmt.Printf("A Random Float64 : %f \n",randomFloat)


}