package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}

	return cards
}

func newDeckFromFile(fileName string) deck {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fileStringContent := string(contents)
	stringSlice := strings.Split(fileStringContent, ",")
	return deck(stringSlice)
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (deck deck) print() {
	for i, card := range deck {
		fmt.Println(i, card)
	}
}

func (deck deck) shuffle() {
	source := rand.NewSource(time.Now().Unix())
	generator := rand.New(source)
	for i := range deck {
		newPosition := generator.Intn(len(deck) - 1)
		deck[i], deck[newPosition] = deck[newPosition], deck[i]
	}
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
