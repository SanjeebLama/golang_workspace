package main

import "fmt"

func main() {
	var a = 5

	fmt.Printf(&a)
	increase(&a)

	fmt.Println(a)
}

func increase(x *int) {
	*x = *x + 1
}
