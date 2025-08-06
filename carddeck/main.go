package main

import (
	"fmt"

	"github.com/jasonluck/learning-go/carddeck/deck"
)

func main() {
	deck := deck.New()
	fmt.Println("Fresh deck out of the box:")
	fmt.Println(deck)
	fmt.Println("Shuffling the deck...")
	deck.Shuffle()
	fmt.Println(deck)
}
