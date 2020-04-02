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

//Option function to execute on Deck creation
type Option func([]Card) []Card

/*
New returns a new Deck, sorted A-K in spades, diamonds,
clubs, hearts order.
*/
func New(opts ...Option) Deck {
	cards := make([]Card, 52)

	for i := 0; i < len(cards); i++ {
		cards[i] = Card{value: (i % 13) + 1, suit: Suit((i / 13) + 1)}
	}

	for _, opt := range opts {
		cards = opt(cards)
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
func Jokers(numJokers int) Option {
	return func(cards []Card) []Card {
		for i := 0; i < numJokers; i++ {
			cards = append(cards, Card{value: 0, suit: JOKER})
		}
		return cards
	}
}

//Sort the deck of cards as specified by the provided sort.Interface
func (d *Deck) Sort(less func(cards []Card) func(i, j int) bool) {
	sort.Slice(d.cards, less(d.cards))
}

//Draw the top card of the deck
func (d *Deck) Draw() Card {
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}
