package main

import "fmt"

func printCard() {
	cards := getCard()

	newList := append(cards, "Six of Spades")

	for i, card := range newList { // this is for range loop
		fmt.Println(i, card)
	}
}

func getCard() []string {
	cardList := []string{"Ace of Diamonds", newCard()}
	return cardList
}

func newCard() string {
	return "Ace of Spades"
}
