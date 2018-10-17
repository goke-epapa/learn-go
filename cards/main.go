package main

func main() {
	cards := NewDeck()

	cards.Shuffle()

	cards.Print()
}
