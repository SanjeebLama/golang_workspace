package main

import "fmt"

func fib(n uint) uint {
	if n <= 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	n := uint(10)
	// fmt.Println(fib(n))
	fmt.Println(fib(n - 1))

}
