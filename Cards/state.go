package main

import "fmt"

func printCard() {
	card := getCard()
	fmt.Println(card)
}

func getCard() string {
	return "Five of Diamonds"
}
