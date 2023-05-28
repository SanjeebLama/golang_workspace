package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Status Checker")

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://twitter.com",
		"https://linkedin.com",
	}

	/**
	** SERIAL Link checking - one at a time
	**/

	// for _, link := range links {
	// 	checkLink(link)
	// }

	/**
	** Go Routine needs Channel to communicate with each other - main routine and  go routines
	**/

	// Declare a channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	/**
	** Receiving message from channel 1:
	**/

	//Receiving the message from channel is a blocking operation
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c) // this is bad practice because it will block the main routine

	/**
	** Receiving message 2:
	**/

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	/**
	** Receiving message 3: Following infinite loop will be executed forever and will not terminate.
	**/

	for {
		go checkLink(<-c, c) // this statement will block the main routine
	}

	// Alternative way of receiving message

	// for l := range c {
	// 	go checkLink(l, c)
	// }

	// Add sleep time

	// for l := range c {
	// 	go func(link string) {
	// 		time.Sleep(5 * time.Second)
	// 		checkLink(link, c)
	// 	}(l)
	// }

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- "hey, might be down"
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	// c <- "hey, is up"
	c <- link
}
