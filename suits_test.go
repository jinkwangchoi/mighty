package mighty

import "testing"

func TestSuits(t *testing.T) {
	assertEqualSuit(Clubs, 1, t)
	assertEqualSuit(Diamonds, 2, t)
	assertEqualSuit(Hearts, 3, t)
	assertEqualSuit(Spades, 4, t)
}

func assertEqualSuit(s Suit, v uint8, t *testing.T) {
	if s != Suit(v) {
		t.Fail()
	}
}
