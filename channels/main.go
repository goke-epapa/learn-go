package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"http://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string) // Creates a new channel

	for _, link := range links {
		go checkLinkStatus(link, c) // Create a child go routine
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLinkStatus(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Test")
		c <- "Link " + link + " might be down"
		return
	}

	c <- link + " is up"
}
