package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	fmt.Println("Welcome to Struct")

	// tony := person{"Tony", "Stark"} // donot use it like this

	// tony := person{firstName: "Tony", lastName: "Stark"}
	//
	// fmt.Println(tony)

	// fmt.Println(tony.firstName)
	// fmt.Println(tony.lastName)
	// fmt.Printf("%T", tony.firstName)

	// var ironMan person

	// fmt.Println(ironMan)

	// fmt.Println(ironMan.firstName)

	// fmt.Println(ironMan.lastName)

	// fmt.Printf("%+v", ironMan)

	// ironMan.firstName = "Tony"
	// ironMan.lastName = "Stark"
	// ironMan.contact.email = "tony@stark.industries"

	// fmt.Println(ironMan)

	var tony = person{
		lastName:  "Stark",
		firstName: "Tony",
		contactInfo: contactInfo{
			zip:   123,
			email: "tony@stark.com",
		},
	}

	tonyPointer := &tony
	tonyPointer.updateFirstName("IronMan") // This won't work because we are not using pointer

	tony.print()
}

func (p person) print() {
	fmt.Printf(" %+v", p)
}

func (pointerToPersonValue *person) updateFirstName(newFirstName string) {
	(*pointerToPersonValue).firstName = newFirstName
}
