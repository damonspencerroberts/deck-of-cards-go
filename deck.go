package main

import "fmt"

// Create a new type of deck
// which is a slice of strings

//these two things are 100% equal
type deck []string

//receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
