package card

import "testing"

func TestSuits(t *testing.T) {
	assertEqual(Clubs, 1, t)
	assertEqual(Diamonds, 2, t)
	assertEqual(Hearts, 3, t)
	assertEqual(Spades, 4, t)
}

func assertEqual(s Suit, v uint8, t *testing.T) {
	if s != Suit(v) {
		t.Fail()
	}
}
