package main

func main() {
	cards := newDeckFromFile("cards")

	cards.print()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("cards")
}
