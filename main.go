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
func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}
func (h Hand) Score() int {
	minScore := h.MinScore()
	// if minScore is > 11, then we can't use the ace as 11
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			// ace is worth 1, we are changing it to 11
			return minScore + 10
		}
	}
	return minScore
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
	// if dealer's score <= 16, then dealer has to hit
	// if dealer has a soft 17, then dealer has to hit

	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, cards = draw(cards)
		dealer = append(dealer, card)
	}
	pScore, dScore := player.Score(), dealer.Score()
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", player, "\nScore:", pScore)
	fmt.Println("Dealer:", dealer, "\nScore:", dScore)
	switch {
	case pScore > 21:
		fmt.Println("BUST!")
	case dScore > 21:
		fmt.Println("DEALER BUSTS!")
	case pScore > dScore:
		fmt.Println("PLAYER WINS!")
	case dScore > pScore:
		fmt.Println("DEALER WINS!")
	case dScore == pScore:
		fmt.Println("DRAW!")
	}
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
