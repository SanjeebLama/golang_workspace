package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go printNumber(&wg)
	go printHello(&wg)
	wg.Wait()
	fmt.Println("Inside MAIN function")

}

func printNumber(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		println("Number :", i)
	}
}
func printHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello world")
}
