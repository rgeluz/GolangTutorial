package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string //this is like saying, type deck === []string

//newDeck(). Create and return a list of playing cards. Essentially an array of strings
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

//print(). Log out the contents of a deck of cards
func (d deck) print() { //notice the receiver (d deck), variable d of type deck, d is similar to this. By convention the variable receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//deal(). Create a 'hand' of cards.
func deal(d deck, handSize int) (deck, deck) {
	//fmt.Print(d) //output the entire deck
	return d[:handSize], d[handSize:]
}

//toString() takes all of the strings in the slice and creates and returns a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",") //[]string(d) deck to slice of string conversion and then join all strings into a single string
}

//saveToFile() uses WriteFile. Save a list of cards to a file on the  local machine
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) //0666 is a default permission
}

//newDeckFromFile(). Load a list of cards from the local machine
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1) //O is a success, anything not 0 is an error and will exit the program
	}

	s := strings.Split(string(bs), ",") //byte slice to string conversion
	return deck(s)
}

//shuffle(). Shuffles all the cards in a deck
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) //generate a new seed for random number generator
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
