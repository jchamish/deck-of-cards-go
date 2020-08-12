package main

func main() {
	cards := newDeck()
	cards.print()

	// deal func
	hand, remainingCards := deal(cards, 5)
	// print
	hand.print()
	remainingCards.print()

	// save cards to file
	cards.saveToFile("my_cards.csv")

	// shuffle
	remainingCards.shuffle()
}