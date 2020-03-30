package deck

import (
	"fmt"
	"math/rand"
	"sort"
)

//Deck represents a deck of standard playing cards
type Deck struct {
	cards []Card
}

/*
New returns a new Deck, sorted A-K in spades, diamonds,
clubs, hearts order.
*/
func New() Deck {
	cards := make([]Card, 52)

	for i := 0; i < len(cards); i++ {
		cards[i] = Card{value: (i % 13) + 1, suit: Suit((i / 13) + 1)}
	}

	return Deck{
		cards: cards,
	}
}

func (d Deck) String() string {
	return fmt.Sprint(d.cards)
}

//Shuffle the cards in the deck
func (d *Deck) Shuffle() {
	newCards := make([]Card, len(d.cards))
	for i, v := range rand.Perm(len(d.cards)) {
		newCards[i] = d.cards[v]
	}
	d.cards = newCards
}

//Jokers are added to the deck
func (d *Deck) Jokers(numJokers int) {
	for i := 0; i < numJokers; i++ {
		d.cards = append(d.cards, Card{value: 0, suit: JOKER})
	}
}

//Sort the deck of cards as specified by the provided sort.Interface
func (d *Deck) Sort(less func(cards []Card) func(i, j int) bool) {
	sort.Slice(d.cards, less(d.cards))
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		if cards[i].value == cards[j].value {
			return cards[i].suit < cards[j].suit
		}
		return cards[i].value < cards[j].value
	}
}
