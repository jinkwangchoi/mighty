package mighty

import (
	"fmt"
)

type Bundle struct {
	Cards []*Card
}

func (b Bundle) Count() int {
	return len(b.Cards)
}

func (b *Bundle) Pop() (Card, error) {
	if b.Count() == 0 {
		return Card{}, fmt.Errorf("bundle empty")
	}
	pop := *b.Cards[0]
	b.Cards = b.Cards[1:]
	return pop, nil
}

func NewBundle(r func(min, max int) int) *Bundle {
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
	return &Bundle{Cards: shuffledCards}
}
