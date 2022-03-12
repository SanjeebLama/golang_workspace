package main

import (
	"fmt"
	"sync"
)

var mu sync.Locker
var num int = 0

func main() {
	wg := &sync.WaitGroup{}

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go updateMap(wg, i)
	}

	wg.Wait()
	fmt.Println("Value of ", num)

	fmt.Println("Exiting Main functionß")
}
func updateMap(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	defer mu.Unlock()

	mu.Lock()
	num = i
}
