package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Weclome to 06 Time")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// Format
	fmt.Println("Format:", presentTime.Format("01-02-2006 15:04:05 Monday"))

	// https://pkg.go.dev/time#Time.Format

	// Created date
	createdDate := time.Date(2022, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
