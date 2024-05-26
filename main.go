package main

import (
	"fmt"
	"github.com/aaaaayushh/deck"
	"strings"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

// DealerString because one of the dealer's cards is hidden
func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func main() {
	// 3 decks of cards, shuffled
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	var player, dealer Hand
	// deal 2 cards to each player
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}
	var input string
	for input != "s" {
		// print the hands
		fmt.Println("Player:", player)
		fmt.Println("Dealer:", dealer.DealerString())
		fmt.Println("What will you do? (h)it or (s)tand?")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		}
	}
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", player)
	fmt.Println("Dealer:", dealer)
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
