package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 9 {
		t.Error("Expected deck length of 9 but got", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Error("Expected first card of Ace of Spades but got", d[0])
	}

	if d[len(d)-1] != "Three of Diamonds" {
		t.Error("Expected last card of Three of Clubs but got", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting") // remove file if exists before test starts because we want to start with a clean slate

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 9 {
		t.Error("Expected 9 cards in deck, got", len(loadedDeck))
	}

	// os.Remove("_decktesting")
}
