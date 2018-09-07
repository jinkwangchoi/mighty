package mighty

import (
	"fmt"
)

type Deck struct {
	Cards []*Card
}

func (b Deck) Count() int {
	return len(b.Cards)
}

func (b *Deck) Pop() (Card, error) {
	if b.Count() == 0 {
		return Card{}, fmt.Errorf("bundle empty")
	}
	pop := *b.Cards[0]
	b.Cards = b.Cards[1:]
	return pop, nil
}

func NewDeck(r func(min, max int) int) *Deck {
	var cards []*Card
	for _, suit := range Suits {
		for _, number := range Numbers {
			if suit == Jokers {
				if number != Joker {
					continue
				}
			} else {
				if number == Joker {
					continue
				}
			}
			c := NewCard(suit, number)

			cards = append(cards, c)
		}
	}

	shuffledCards := make([]*Card, len(cards))
	for i := range cards {
		randomIndex := r(0, len(cards)-i)
		shuffledCards[i] = cards[randomIndex]
		cards = append(cards[:i], cards[i:]...)
	}
	return &Deck{Cards: shuffledCards}
}
