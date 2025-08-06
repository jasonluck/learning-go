package deck

import (
	"testing"
)

// New Deck should contain 52 cards and be sorted A-K, in Spades, Diamonds, Clubs, Hearts order
func TestNew(t *testing.T) {
	deck := New()

	if len(deck.cards) != 52 {
		t.Errorf("Deck only contained %d cards. It was expected to contain 52 cards.", len(deck.cards))
	}

	for suitCount, suit := range Suits() {
		for i := range 13 {
			c := deck.cards[(suitCount*13)+i]
			expected := Card{
				value: i + 1,
				suit:  suit,
			}
			if c != expected {
				t.Errorf("Card at position %d in the deck was expected to be %q, but was %q", i, expected, c)
			}
		}
	}
}

func TestDeckString(t *testing.T) {
	deck := New()
	result := deck.String()
	expected := "[AS 2S 3S 4S 5S 6S 7S 8S 9S 10S JS QS KS AD 2D 3D 4D 5D 6D 7D 8D 9D 10D JD QD KD AC 2C 3C 4C 5C 6C 7C 8C 9C 10C JC QC KC AH 2H 3H 4H 5H 6H 7H 8H 9H 10H JH QH KH]"
	if result != expected {
		t.Errorf("Expected %q, but was %q", expected, result)
	}
}

func TestShuffle(t *testing.T) {
	a := New()
	b := New()
	b.Shuffle()
	t.Log(a)
	t.Log(b)
	if len(a.cards) != len(b.cards) {
		t.Error("After shuffling the deck was not the same size.")
	}
}

func TestNewJokers(t *testing.T) {
	numJokers := 3
	deck := New(Jokers(numJokers))

	if len(deck.cards) != 52+numJokers {
		t.Errorf("After adding %d Jokers, the deck size should have been %d but was %d", numJokers, 52+numJokers, len(deck.cards))
	}
	if deck.cards[52+numJokers-1].suit != JOKER {
		t.Log(deck)
		t.Errorf("Expected last card to be a Joker, but was %q", deck.cards[52+numJokers-1])
	}
}

func TestDraw(t *testing.T) {
	d := New()
	numCards := 1
	size := len(d.cards)
	expected := d.cards[0:numCards]
	result, err := d.Draw(numCards)
	if err != nil {
		t.Errorf("Unexpected error when drawing a card: %v", err)
	}
	if len(result) != numCards {
		t.Errorf("Expected to draw %d card(s), but got %d", numCards, len(result))
	}
	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected drawn card to be %q, but was %q", expected, result)
		}
	}
	if len(d.cards) != size-numCards {
		t.Error("After drawing, cards were not removed from the deck.")
	}
}

func TestDrawTooMany(t *testing.T) {
	d := New()
	numCards := 100 // More than the deck size
	expected := d.cards[0:len(d.cards)]
	result, err := d.Draw(numCards)

	if err != ErrNotEnoughCards {
		t.Errorf("Expected error %v, but got %v", ErrNotEnoughCards, err)
	}
	if len(result) != len(expected) {
		t.Errorf("Expected to draw %d card(s), but got %d", len(expected), len(result))
	}
	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected drawn card to be %q, but was %q", expected, result)
		}
	}
	if len(d.cards) != 0 {
		t.Error("After drawing, it was expected that all cards would be removed from the deck, but some were left.")
	}
}

func TestPutOnBottom(t *testing.T) {
	d := New()
	initialSize := len(d.cards)

	expected := Card{value: 1, suit: SPADES}
	d.PlaceOnBottom(expected)
	resultSize := len(d.cards)
	result := d.cards[resultSize-1]
	if expected != result {
		t.Errorf("Expected card to be placed on bottom of deck to be %q, but was %q", expected, result)
	}
	if resultSize != initialSize+1 {
		t.Error("After placing a card on the bottom, the deck size did not increase by 1.")
	}
}

func TestPutOnTop(t *testing.T) {
	d := New()
	initialSize := len(d.cards)

	expected := Card{value: 5, suit: HEARTS}
	d.PlaceOnTop(expected)
	resultSize := len(d.cards)
	result := d.cards[0]
	if expected != result {
		t.Errorf("Expected card to be placed on top of deck to be %q, but was %q", expected, result)
	}
	if resultSize != initialSize+1 {
		t.Error("After placing a card on the top, the deck size did not increase by 1.")
	}
}
