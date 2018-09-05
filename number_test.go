package mighty

import "testing"

func TestNumber(t *testing.T) {
	assertEqualNumber(Two, 2, t)
	assertEqualNumber(Three, 3, t)
	assertEqualNumber(Four, 4, t)
	assertEqualNumber(Five, 5, t)
	assertEqualNumber(Six, 6, t)
	assertEqualNumber(Seven, 7, t)
	assertEqualNumber(Eight, 8, t)
	assertEqualNumber(Nine, 9, t)
	assertEqualNumber(Ten, 10, t)
	assertEqualNumber(Jack, 11, t)
	assertEqualNumber(Queen, 12, t)
	assertEqualNumber(King, 13, t)
	assertEqualNumber(Ace, 14, t)
}

func assertEqualNumber(s Number, v uint8, t *testing.T) {
	if s != Number(v) {
		t.Fail()
	}
}
