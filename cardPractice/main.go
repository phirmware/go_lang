package main

import "fmt"

func main() {
	var cards card
	cards = newDeck()
	first, second := cards.cut(5)
	fmt.Println(first, second)
}
