package main

func printCard() {
	cards := getCard()

	newList := append(cards, "Six of Spades")

	// for i, card := range newList { // this will throw error as we are not using index
	// 	fmt.Println(card)
	// }

	newList.printCards() // We are passing newList to printCards() printCards is inside deck.go
}

func getCard() deck { // deck is a custom type
	cardList := deck{"Ace of Diamonds", newCard()}
	return cardList
}

func newCard() string {
	return "Ace of Spades"
}
