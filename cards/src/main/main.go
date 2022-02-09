package main

func newCard() string {
	return "Five of Diamonds"
}

// no need to add a receiver because not comes from already existing

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
