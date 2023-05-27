package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello HTTP")

	res, err := http.Get("http://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Response : ", res)

	bs := make([]byte, 99999)
	res.Body.Read(bs)

	fmt.Println(string(bs))
}
