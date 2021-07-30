package main

import "fmt"

type bot interface {
	// Check if there a type with a function 'getGreeting' with an argument of 'blank' and return 'string'
	// They became a honorary member of type bot
	// So they could be considered to be type 'bot'
	getGreeting() string
	// getGreeting(string, int) (string, error)
}

// struct is a concrete type
type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//func printGreeting(eb englishBot) {
//	fmt.Println(eb.getGreeting())
//}

//func printGreeting(sb englishBot) {
//	fmt.Println(sb.getGreeting())
//}

// For example! GetGreeting will probably have very different logic
func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}