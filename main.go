package main

func main() {
	cards := deck{newCard()}

	cards.print()
}

func newCard() string {
	return "Diamond"
}
