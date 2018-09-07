package mighty

type Suit uint8

const (
	Clubs Suit = 1 + iota
	Diamonds
	Hearts
	Spades
	Jokers
)

var Suits = []Suit{Clubs, Diamonds, Hearts, Spades, Jokers}
