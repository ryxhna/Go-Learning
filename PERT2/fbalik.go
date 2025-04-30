package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var randomValue int

	// create and copy  3x
	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number: ", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number: ", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number: ", randomValue)
}

func randomWithRange(min, max int) int {
	var val = rand.Int() % (max - min + 1) + min
	return val
}