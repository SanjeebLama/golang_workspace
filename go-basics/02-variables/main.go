package main

import "fmt"

// Variables declared with Capital letters are public by default.
const hello = "Hello"

const Hello = "Capital hello"

func main() {

	// var username string = "Tony"
	// fmt.Println(username)
	// fmt.Printf("Variable is of type: %T \n", username)

	// var isLoggedIn bool = true
	// fmt.Println(isLoggedIn)
	// fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// var smallValue int8 = 127
	// // int8 is a signed 8-bit integer
	// // 2^7 = 128 [0-127]
	// fmt.Println(smallValue)
	// fmt.Printf("Variable is of type: %T \n", smallValue)

	// var smallFloat float32 = 253235.45983556
	// // float32 is a 32-bit floating point number
	// fmt.Println(smallFloat)
	// fmt.Printf("Variable is of type: %T \n", smallFloat)

	// Default Values and alias
	// var defaultInt int // default values is 0
	// fmt.Println(defaultInt)
	// fmt.Printf("Variable is of type: %T \n", defaultInt)

	// var defaultString string
	// fmt.Println(defaultString) // default values is ""
	// fmt.Printf("Variable is of type: %T \n", defaultString)

	log := "Variable is of type: %T \n"

	// implicit types:
	var a = "initial"
	fmt.Println(a)
	fmt.Printf(log, a)

	// no-vars ** This is only possible inside of the function
	y := 20
	fmt.Println(y)
	fmt.Printf(log, y)

	z := 12312312.123123
	fmt.Println(z)
	fmt.Printf(log, z)

	z = 9090
	fmt.Println("new", log, z)
	fmt.Printf(log, z)

	// This will throw errors.
	// z = "hello world"
	// fmt.Println("new", log, z)
	// fmt.Printf(log, z)

	fmt.Println(hello)
	fmt.Println(Hello)
}
