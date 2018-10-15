package main

import (
	"fmt"
)

func main() {
	cards := []string{newCard(), newCard(), "Ace of diamonds"}

	// Append does not modify a slice, it only creates a new slice
	// Slices are immutable
	cards = append(cards, "Six of spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
