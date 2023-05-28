package main

import "fmt"

func main() {
	fmt.Println("** Pointers **")

	var ptr *int // Here * is used to declare a pointer

	fmt.Println("Value of Pointer is ", ptr) // If you donot pass any value to pointer it will return nil

	myNum := 7

	var numPtr = &myNum // Here & is used to get the address of a variable

	fmt.Println("Value of numPtr  is ", numPtr) //  It will return the address of the variable

	fmt.Println("Value of numPtr is ", *numPtr) // It will return the value at the address i.e. 7

	*numPtr = 8 // It will change the value at the address i.e. 8

	fmt.Println("Value of myNum is ", myNum)

	fmt.Println("Value of numPtr is ", numPtr)
	fmt.Println("Value of *numPtr is ", *numPtr)

}
