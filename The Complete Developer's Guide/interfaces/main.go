package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	printGreeting(eb)

	sb := spanishBot{}
	printGreeting(sb)
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}


func (englishBot) getGreeting() string {
	// VERY CUSTOM logic
	return "Hi, there!"
}

func (spanishBot) getGreeting() string {
	// VERY CUSTOM logic
	return "Hola!"
}
