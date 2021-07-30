package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %s", d[0])
	}

	if d[len(d) - 1] != "King of Hearts" {
		t.Errorf("Expected last card of King of Hearts, but got %s", d[len(d) - 1])
	}
}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_deck_testing")

	deck := newDeck()
	deck.saveToFile("_deck_testing")

	loadedDeck := newDeckFromFile("_deck_testing")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %d", len(loadedDeck))
	}

	os.Remove("_deck_testing")

}