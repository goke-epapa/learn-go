package main

func main() {
	cards := newDeck()

	// Append does not modify a slice, it only creates a new slice
	// Slices are immutable
	cards = append(cards, "Six of spades")

	cards.print()
}
