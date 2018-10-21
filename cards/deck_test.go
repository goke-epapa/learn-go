package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck lenght of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, bug got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of Ace of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	fn := "_decktesting"
	os.Remove(fn)

	d := NewDeck()
	d.SaveToFile(fn)

	loadedDeck := NewDeckFromFile(fn)

	if len(loadedDeck) != len(d) {
		t.Errorf("Expected %v in deck, got %v", len(d), len(loadedDeck))
	}

	os.Remove(fn)
}
