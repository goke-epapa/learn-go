package main

func main() {
	cards := newDeck()

	// Append does not modify a slice, it only creates a new slice
	// Slices are immutable
	cards = append(cards, "Six of spades")

	hand, remainingDeck := deal(cards, 5)

	hand.print()

	remainingDeck.print()
}
