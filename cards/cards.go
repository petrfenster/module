package cards

func cards() {
	cards := NewDeck()
	cards.shuffle()
	cards.Print()
}
