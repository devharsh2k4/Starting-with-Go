package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	randomNumber := rand.Intn(6)
	rand.Seed(time.Now().Unix())
	fmt.Println("you will get a", randomNumber)
}
