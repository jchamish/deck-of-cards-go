package main

func main() {
	cards := newDeck()
	cards.print()

	// deal func
	hand, remainingCards := deal(cards, 5)
	// print
	hand.print()
	remainingCards.print()
}