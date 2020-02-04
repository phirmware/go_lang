package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := deck{"spades", "diamond", "heart"}
	cardValues := deck{"ace", "two", "three", "four"}
	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func readFile(fileName string) deck {
	strBuffer, error := ioutil.ReadFile(fileName)
	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}
	r := strings.Split(string(strBuffer), ",")
	return deck(r)
}
