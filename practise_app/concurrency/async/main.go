package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)

	<-done
	<-done
	<-done
	<-done

	fmt.Println(" Exiting Main Function")
}

func task1(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 1 - 100ms")
	done <- struct{}{}
}

func task2(done chan struct{}) {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Task 2 - 10ms")
	done <- struct{}{}
}

func task3(done chan struct{}) {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Task 3 - 300ms")
	done <- struct{}{}
}

func task4(done chan struct{}) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Task 4 - 200ms")
	done <- struct{}{}
}
