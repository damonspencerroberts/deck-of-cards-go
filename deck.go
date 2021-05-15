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

// deck, deck tells go we return 2 values both of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// joins all the cards into a single string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//uses io utils to save to a file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	//returns two elements byte slice and the error
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1 - log the error and return a call to newDeck()
		// Option 2 - log the error and quit program

		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// turns to deck to use all of the above functions
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := range d {
		np := r.Intn(len(d) - 1)
		d[i], d[np] = d[np], d[i]
	}
}