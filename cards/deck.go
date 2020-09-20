package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string //this is like saying, type deck === []string

//newDeck()
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Act", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//print()
func (d deck) print() { //notice the receiver (d deck), variable d of type deck, d is similar to this. By convention the variable receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}
