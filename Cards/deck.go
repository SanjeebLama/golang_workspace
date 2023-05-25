package main

import "fmt"

type deck []string

// func (newList deck) printCards() {
// 	// This function is a receiver function.
// 	// Receiver Function is a function that is associated with a type.

// 	for i, card := range newList { // this is for range loop
// 		fmt.Println(i, card)
// 	}
// }

func (d deck) displayCards() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

// we are using d  as receiver function because this is a deck type and it's convention to define it with single letter.

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds"}

	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
