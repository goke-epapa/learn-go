package main

import (
	"fmt"
)

func main() {
	// first map syntax
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	fmt.Println(colours)

	// var syntax
	var newColours map[string]string
	fmt.Println(newColours)

	// make syntax
	someColours := make(map[string]string)
	someColours["white"] = "#ff0000"
	fmt.Println(someColours)

	// deleting values in maps
	delete(colours, "red")
	fmt.Println(colours)

	printMap(someColours)
}

func printMap(m map[string]string) {
	for colour, hex := range m {
		fmt.Printf("Hex code for %s is %s", colour, hex)
	}
}
