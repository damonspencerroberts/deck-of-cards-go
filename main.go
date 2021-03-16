package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}

//Returns of type string
func newCard() string {
	return "Five of Diamonds"
}