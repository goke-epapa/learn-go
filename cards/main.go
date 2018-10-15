package main

func main() {
	cards := deck{newCard(), newCard(), "Ace of diamonds"}

	// Append does not modify a slice, it only creates a new slice
	// Slices are immutable
	cards = append(cards, "Six of spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
