package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	fmt.Println("TestNewDeck: OK")
}

func TestShuffleDeck(t *testing.T) {

	d := newDeck()
	d.shuffleDeck()
	s := newDeck()
	s.shuffleDeck()
	if EqualStringSlices(d, s) {
		t.Errorf("Expected deck to be shuffled, but are identical")
	}

	fmt.Println("TestShuffleDeck: OK")
}

func TestSaveToDeckAndLoadNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := loadFromFile("_decktesting")

	if len(loadedDeck) > 52 && len(loadedDeck) <= 0 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

	fmt.Println("TestSaveToDeckAndNewDeckFromFile: OK")
}

// ######### HELPER FUNCTIONS #########

func EqualStringSlices(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}

	return true
}
