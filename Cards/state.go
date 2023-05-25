package main

import "fmt"

func printCard() {
	cards := getCard()

	newList := append(cards, "Six of Spades")

	for i, card := range newList { // this is for range loop
		fmt.Println(i, card)
	}

	// for i, card := range newList { // this will throw error as we are not using index
	// 	fmt.Println(card)
	// }
}

func getCard() []string {
	cardList := []string{"Ace of Diamonds", newCard()}
	return cardList
}

func newCard() string {
	return "Ace of Spades"
}
