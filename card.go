package mighty

type Card struct {
	Suit   Suit
	Number Number
}

func (c Card) IsJoker() bool {
	return c.Suit == Jokers
}

func NewJoker() *Card {
	return &Card{
		Jokers,
		Joker,
	}
}

func NewCard(suit Suit, number Number) *Card {
	return &Card{
		suit,
		number,
	}
}
