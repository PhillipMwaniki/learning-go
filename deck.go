package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
