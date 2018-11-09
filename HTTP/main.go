package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// byte slice with 99999 elements
	// ensure the size of the byte slice is large enough
	// read function doesn't resize byte slice
	bs := make([]byte, 99999)

	resp.Body.Read(bs)

	fmt.Println(string(bs))

	fmt.Println(resp)
}
