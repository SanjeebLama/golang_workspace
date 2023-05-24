package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {

	// Random Number from math
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("Random Number", rand.Intn(25))

	// Random Number from crypto
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(25))

	fmt.Println(myRandomNumber)

}
