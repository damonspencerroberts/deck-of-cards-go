package main

import "fmt"

// Create a new type of deck
// which is a slice of strings

//these two things are 100% equal
type deck []string

func newDeck() deck {
	// returns a value of type deck
	cards := deck{}

	cardSuits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, vals := range cardValues {
			newString := fmt.Sprintf("%v of %s", vals, suit)
			cards = append(cards, newString)
		}
	}

	return cards
}

//receive
// Any variable of type deck gets acces to the print method
// d work as this, self in this case
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
