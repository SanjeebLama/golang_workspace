package main

import "fmt"

func main() {
	const age int = 1

	if age >= 16 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You cannot vote.")
	}

}
