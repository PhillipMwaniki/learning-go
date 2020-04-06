package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	suites := []string{"Spades", "Diamonds", "Hearts", "Clubs", "Cloves"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six"}

	for _, suite := range suites {
		for _, val := range cardValues {
			cards = append(cards, suite+" of "+val)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
