package main

import (
	"fmt"

	"new_path.com/point"
	emoji "new_path.com/quotes"
)

func main() {
	pt1 := point.Point{X: 2, Y: 3}
	fmt.Println(pt1)
	fmt.Println(pt1.Length())
	fmt.Println(emoji.GetEmoji("smile"))
}
