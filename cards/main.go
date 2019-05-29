package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	//newHand, remainingCards := deal(cards, 5)
	//newHand.print()
	//remainingCards.print()
	//fmt.Println(remainingCards.toString())
	//fromFile := newDeckFromFile("test.txt")
	//fromFile.print()
}

func newCard() string {
	return "Five of Diamonds"
}
