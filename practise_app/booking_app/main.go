package main

import "fmt"

func main() {
	var firstName string
	var lastName string

	var email string
	var ticket uint

	const totalTicket = 50

	fmt.Println("Welcome to Booking App")
	fmt.Println("Enter your First name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("How any ticket do you want to book?")
	fmt.Scan(&ticket)

	fmt.Printf("Dear %s%s, you have booked %d tickets. Your ticket confirmation is sent to your email: %s \n Thank you\n", firstName, lastName, ticket, email)

	fmt.Println("Remaining Ticket :", totalTicket-ticket)

	// fmt.Println(`"Hello ", checking `, firstName)

	// test := "onetwothree"
	// fmt.Println(len(test))
}
