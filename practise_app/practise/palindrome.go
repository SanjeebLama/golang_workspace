package main

import "fmt"

func main() {
	var str string = "MOMAM"
	data := false
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			data = false
		} else {
			data = true
		}
	}
	fmt.Println(data)
	fmt.Println(len(str) / 2)
}
