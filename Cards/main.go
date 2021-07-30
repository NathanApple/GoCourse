package main

func main() {
	// This is a slice
	cards := newDeck()
	cards.shuffle()
	cards.print()

	//
	//cards.saveToFile("my_cards.txt")
	//hand, remainingDeck := deal(cards, 5)
	//
	//hand.print()
	//remainingDeck.print()

	//cards := newDeckFromFile("my_cards.txt")
	//cards.print()
}
