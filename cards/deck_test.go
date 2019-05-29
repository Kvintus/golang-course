package main

import (
	"os"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Error("Expected deck lenght of 52, got: ", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, got: %v", deck[0])
	}

	lastCard := deck[len(deck)-1]
	if lastCard != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, got: %v", lastCard)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testingFileName := "_decktesting.txt"
	os.Remove(testingFileName)
	deck := newDeck()
	deck.shuffle()
	deck.saveToFile(testingFileName)

	loadedDeck := newDeckFromFile(testingFileName)

	if !reflect.DeepEqual(deck, loadedDeck) {
		t.Errorf("Expected decks to match, but they don't")
	}

	os.Remove(testingFileName)
}
