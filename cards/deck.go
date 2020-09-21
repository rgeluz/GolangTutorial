package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string //this is like saying, type deck === []string

//newDeck()
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

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

//deal()
func deal(d deck, handSize int) (deck, deck) {
	//fmt.Print(d) //output the entire deck
	return d[:handSize], d[handSize:]
}

//toString() takes all of the strings in the slice and creates and returns a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",") //[]string(d) deck to slice of string conversion and then join all strings into a single string
}

//saveToFile() uses WriteFile
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) //0666 is a default permission
}
