package main

func main() {
	cards := newDeck()
	// fmt.Println(deal(cards, 4))
	// cards.print()
	hand, remainingDeck := deal(cards, 4)
	hand.print()
	remainingDeck.print()
}
