package main

import "fmt"

func main() {

	/**
	** Map 1
	**/

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"blue":  "#0000ff",
	// 	"green": "#00ff00",
	// }

	// fmt.Println("Colors", colors)

	/**
	** Map 2
	**/

	// var colors map[int]string

	// fmt.Println("Colors", colors) // Output: Colors map[]

	/**
	** Map 3
	**/

	// colors := make(map[int]string)

	// colors[0] = "#ff0000"

	// delete(colors, 0) // Delete key 0

	// fmt.Println("Colors", colors)

	/**
	** Map 4
	**/

	colors := map[string]string{
		"red":    "#ff0000",
		"blue":   "#0000ff",
		"green":  "#00ff00",
		"orange": "#ffa500",
	}

	printColors(colors)

}

func printColors(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
