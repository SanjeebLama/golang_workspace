package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	fmt.Println("Welcome to Struct")

	// tony := person{"Tony", "Stark"} // donot use it like this

	// tony := person{firstName: "Tony", lastName: "Stark"}

	// fmt.Println(tony)

	// fmt.Println(tony.firstName)
	// fmt.Println(tony.lastName)
	// fmt.Printf("%T", tony.firstName)

	var ironMan person

	fmt.Println(ironMan)

	fmt.Println(ironMan.firstName)

	fmt.Println(ironMan.lastName)

	fmt.Printf("%+v", ironMan)

	ironMan.firstName = "Tony"
	ironMan.lastName = "Stark"

	fmt.Println(ironMan)
}
