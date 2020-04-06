package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 30 {
		t.Errorf("Expected deck of length 30 but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Six of Cloves" {
		t.Errorf("Expected Ace of Spades but got %v", d[len(d)-1])
	}
}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 30 {
		t.Errorf("Expected deck of length 30 but got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
