package main

import "fmt"

func main() {
	// cards := newDeck()
	cards := readFile("my_cards")
	firstItem := cards[0]
	fmt.Println(firstItem)
}
