package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	myCards, _ := deal(cards, 7)
	myCards.print()
}

// Byte slice
// fmt.Println([]byte(greeting))