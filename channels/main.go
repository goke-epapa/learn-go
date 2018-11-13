package main

import (
	"fmt"
	"net/http"
	"time"
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

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLinkStatus(link, c)
		}(l)
	}
}

func checkLinkStatus(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Link " + link + " might be down")
		c <- link
		return
	}

	fmt.Println(link + " is up")
	c <- link
}
