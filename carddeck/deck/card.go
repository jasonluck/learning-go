package deck

import (
	"strconv"
)

//Suit represents the category (or suit) of playing card
type Suit int

//Types of suits
const (
	SPADES Suit = 1 + iota
	DIAMONDS
	CLUBS
	HEARTS
	JOKER
)

var suits = []string{
	"S",
	"D",
	"C",
	"H",
	"Joker",
}

func (s Suit) String() string {
	return suits[s-1]
}

//Card represents a standard playing card
type Card struct {
	value int
	suit  Suit
}

func (c Card) String() string {
	if c.suit == JOKER {
		return c.suit.String()
	}

	var faceVal string
	switch c.value {
	case 11:
		faceVal = "J"
	case 12:
		faceVal = "Q"
	case 13:
		faceVal = "K"
	case 1:
		faceVal = "A"
	default:
		faceVal = strconv.Itoa(c.value)
	}

	return faceVal + c.suit.String()
}

//Suits returns a set of all possible card suits
func Suits() []Suit {
	suits := make([]Suit, 4)
	suits[0] = SPADES
	suits[1] = DIAMONDS
	suits[2] = CLUBS
	suits[3] = HEARTS
	return suits
}
