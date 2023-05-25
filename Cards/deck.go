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

func (d deck) printCards() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// we are using d  as receiver function because this is a deck type and it's convention to define it with single letter.
