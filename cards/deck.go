package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// which is a slice of strings

// Deck : A deck of cards
type Deck []string

//NewDeck : Creates a new deck
func NewDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Print : prints the deck of cards
func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal : deals a hand
func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

// ToString : returns a string representation of the deck
func (d Deck) ToString() string {
	return strings.Join([]string(d), ",")
}

// SaveToFile : saves the deck to a file
func (d Deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0666)
}

// NewDeckFromFile : creates a deck from a file
func NewDeckFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	return Deck(strings.Split(string(bs), ","))
}

// Shuffle : shuffles the deck
func (d Deck) Shuffle() {
	for i := range d {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)

		// Go uses the exact same seed by default
		newPosition := r.Intn(len(d) - 1)

		// Swap items
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
