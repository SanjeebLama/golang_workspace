package main

import (
	"fmt"
	"time"
)

func main() {
	go work() //fork point

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done waiting -  Exiting Main Function")

	//JOIN point - wait group or channel
}
func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Inside Work - 500ms")
}
