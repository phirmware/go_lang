package main

type card []string

func newDeck() card {
	cardName := card{"ace", "queen", "king", "joker"}
	cardTitles := card{"hearts", "spade", "diamond"}
	var newDeck = card{}

	for _, c := range cardName {
		for _, t := range cardTitles {
			newDeck = append(newDeck, c+" of "+t)
		}
	}

	return newDeck
}

func (c card) cut(number int64) (card, card) {
	return c[0:number], c[number:]
}

func saveToFile() {

}
