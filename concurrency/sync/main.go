package main

import (
	"fmt"
	"time"
)

func main() {
	task1()
	task2()
	task3()
	task4()
	fmt.Println("Exiting Main function")
}

func task1() {
	time.Sleep(100 * time.Microsecond)
	fmt.Println("Task 1 - 100ms")
}

func task2() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Task 2 - 10ms")
}

func task3() {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Task 3 - 300ms")
}

func task4() {
	time.Sleep(0 * time.Millisecond)
	fmt.Println("Task 4 - 0 ms")
}
