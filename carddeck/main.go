package main

import (
	"fmt"

	"github.com/jasonluck/learning-go/carddeck/deck"
)

func main() {
	deck := deck.New()
	fmt.Print(deck)
}
