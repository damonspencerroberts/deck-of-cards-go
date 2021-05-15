package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
}

// Byte slice
// fmt.Println([]byte(greeting))