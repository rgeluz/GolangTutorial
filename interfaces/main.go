package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

/*
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}*/

/*
func printGreeting(sb spanishBot) {  //cannot do function overloading in go
	fmt.Println(sb.getGreeting())
}
*/

//reduce the two functions above into a single function
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string { //note: can omit eb, since we are not using it. for example func (englishBot) getGreeting() string {}
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
