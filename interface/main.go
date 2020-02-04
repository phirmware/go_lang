package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGreeting() string
}

func main() {
	sb := englishBot{}
	eb := spanishBot{}

	printGreeting(sb)
	printGreeting(eb)
}

func printGreeting(eb bot) {
	fmt.Println(eb.getGreeting())
}

func (e englishBot) getGreeting() string {
	return "Hello!"
}

func (s spanishBot) getGreeting() string {
	return "Holla!"
}
