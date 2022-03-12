// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	now := time.Now()

// 	var wg sync.WaitGroup
// 	wg.Add(1) //wait for one go routine

// 	go func() {
// 		defer wg.Done()
// 		work()
// 	}() //fork point

// 	//JOIN point - wait group or channel
// 	wg.Wait()
// 	fmt.Println("Elapsed : ", time.Since(now))

// 	fmt.Println("Done waiting -  Exiting Main Function")
// }
// func work() {
// 	time.Sleep(500 * time.Millisecond)
// 	fmt.Println("Inside Work - 500ms")
// }
