package main

import "fmt"

type deck []string

func (deck deck) print() {
	for i, card := range deck {
		fmt.Println(i, card)
	}
}