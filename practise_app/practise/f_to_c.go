package main

import "fmt"

func main() {
	farenheit := 98.6
	var celsius float32
	celsius = float32(farenheit-32) * 5 / 9
	fmt.Println(celsius)
}
