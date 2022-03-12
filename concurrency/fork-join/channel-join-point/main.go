package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})

	go func() {

		work()
		done <- struct{}{}
	}() //fork point

	//JOIN point - wait group or channel
	<-done
	fmt.Println("Elapsed : ", time.Since(now))


	fmt.Println("Done waiting -  Exiting Main Function")
}
func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Inside Work - 500ms")
}
