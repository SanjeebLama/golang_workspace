package main

import "fmt"

func printCard() {
	cards := newDeck()

	// newList := append(cards, "Six of Spades")

	// for i, card := range newList { // this will throw error as we are not using index
	// 	fmt.Println(card)
	// }

	// newList.displayCards() // We are passing newList to printCards() printCards is inside deck.go

	hand, remainingCards := deal(cards, 5)

	// fmt.Println("Hand : ", hand, "\nRemaining Cards : ", remainingCards)

	fmt.Println("Hand : ")
	hand.displayCards()

	fmt.Println("Remaining Cards : ")
	remainingCards.displayCards()

}

// func getCard() deck { // deck is a custom type
// 	cardList := deck{"Ace of Diamonds", newCard()}
// 	return cardList
// }

// func newCard() string {
// 	return "Ace of Spades"
// }
