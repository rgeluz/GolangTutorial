package main

import "fmt"

func main() {

	//variableDeclaration()
	//slicesandloops()
	testNewDeck()

	//test type conversion
	//greeting := "Hi there!"
	//fmt.Println(greeting)         //Hi there!
	//fmt.Println([]byte(greeting)) //[72 105 32 116 104 101 114 101 33]

}

//Variable Declaration Lesson
func variableDeclaration() {
	var card string = "Ace of Spades" //or var card = "Ace of Spades", data type can be inferred from right side of assignment operator. Can just write card := "Ace of Spades"
	fmt.Println(card)

	card1 := "Ace of Spades" //only use := during declaration and initializatoin of variables, no needed during reinitializaiotn
	card1 = "Five of Diamonds"

	fmt.Println(card1)

	card2 := newCard()
	fmt.Println(card2)

	card3 := newCard1()
	fmt.Println(card3)

	printState() //This is from state.go in the same package

}

func newCard() string {
	return "Five of Diamonds"
}

func newCard1() int {
	return 7
}

//SlicesAndForLoops
func slicesandloops() {
	//cards := []string{newCard(), newCard()} //or {"Ace of Diamonds", newCard}
	cards := deck{newCard(), newCard()} //using deck.go's deck type

	//add more elements to cards
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	//iterate slice of cards   //moved this to the deck type
	/* for i, card := range cards {
		fmt.Println(i, card)
	}
	*/

	cards.print()
}

//Test newDeck() function
func testNewDeck() {
	/*cards := newDeck()
	cards.print()

	hand, remainingCards := deal(cards, 5) //deal is defined in the same package main
	hand.print()
	remainingCards.print()
	*/
	cards := newDeck()
	fmt.Println(cards.toString())

	//save list of cards to harddrive
	cards.saveToFile("my_cards")
}
