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
		for i := 0; i < 13; i++ {
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

func TestJokers(t *testing.T) {
	deck := New()
	numJokers := 3
	deck.Jokers(numJokers)

	if len(deck.cards) != 52+numJokers {
		t.Errorf("After adding %d Jokers, the deck size should have been %d but was %d", numJokers, 52+numJokers, len(deck.cards))
	}
	if deck.cards[52+numJokers-1].suit != JOKER {
		t.Log(deck)
		t.Errorf("Expected last card to be a Joker, but was %q", deck.cards[52+numJokers-1])
	}
}
