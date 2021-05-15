package main

import (
	"os"
	"testing"
)
func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length 52 got %v", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("First card should be Ace of Diamonds got %v", d[0])
	}

	if d[len(d) -1] != "King of Spades" {
		t.Errorf("First card should be King of Spades got %v", d[len(d) -1])
	}
}

func TestDeckSave (t *testing.T) {
	os.Remove("_testfile")
	d := newDeck()
	d.saveToFile("_testfile")
	nd := newDeckFromFile("_testfile")

	if len(d) != len(nd) {
		t.Errorf("File not saved correctly, expecting length of deck (%v) got %d", len(d), len(nd))
	}
	
	os.Remove("_testfile")
}