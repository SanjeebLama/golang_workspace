package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

/**
** Task: Create a program that reads the contents of a text file then prints its contents to the terminal.
**/

func main() {
	fmt.Println("Hello World")

	fmt.Println()

	file, err := os.Open(os.Args[1]) // For read access.
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
