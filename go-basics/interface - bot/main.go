package main

import "fmt"

type bot interface {
	getGreeting() string // getGreeting Inside bot
}

func main() {
	fmt.Println("Interface")

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
