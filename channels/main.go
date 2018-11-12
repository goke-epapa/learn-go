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

	for _, link := range links {
		go checkLinkStatus(link)
	}
}

func checkLinkStatus(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Printf("Link %v might be down\n", link)
		return
	}

	fmt.Printf("%v is up\n", link)
}
