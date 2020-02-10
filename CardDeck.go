package main

import "fmt"
import "math/rand"
import "time"

type card struct {
	name string
	color string
	value []int
}

type deck struct {
	cards []card
}

func newDeck() *deck {
	returnDeck := new(deck)
	farben := []string{"Kreuz","Pik","Herz","Karo"}
	for _,farbe := range farben {
		for i :=1 ; i <= 10; i++ {

			newCard := new(card)
			newCard.color = farbe
			newCard.value = append(newCard.value,i)
			newCard.name = fmt.Sprintf("%s %d", farbe, i)
			returnDeck.cards = append(returnDeck.cards, *newCard)
		}
	}
	returnDeck.showDeck()
	return returnDeck
}

func (deck *deck)shuffleDeck() {

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.cards), func(i, j int) { deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i] })

}
//todo get card and delete from desk
func (deck *deck) getCard() card {
	deck.cards[]
}

func (deck *deck) showDeck() {
	fmt.Println(deck)
}



func (deck *deck) addNewDeck() {

}

