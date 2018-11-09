package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please specify the file argument")
		os.Exit(1)
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
