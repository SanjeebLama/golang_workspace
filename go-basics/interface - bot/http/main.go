package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	fmt.Println("Hello HTTP")

	res, err := http.Get("http://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Response : ", res)

	/*******
	** Method 1: Read Response body
	**/

	// bs := make([]byte, 99999) // This will allocate memory for the response
	// fmt.Println(len(bs))      // returns 99999

	// fmt.Printf("Type of bs :  %T\n", bs) // returns []uint8

	// res.Body.Read(bs)

	// fmt.Println(string(bs))

	/*******
	** Method 2: Read Response body
	**/

	io.Copy(os.Stdout, res.Body)

	/*******
	** Creating custom Writer
	**/

	lw := logWriter{}
	io.Copy(lw, res.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes : ", len(bs))
	return len(bs), nil
}
