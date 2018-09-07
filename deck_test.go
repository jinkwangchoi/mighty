package mighty

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeck(t *testing.T) {
	d := NewDeck(func(min, max int) int {
		return min
	})
	a := assert.New(t)
	a.Equal(d.Count(), 53)

	popCard, err := d.Pop()
	a.Equal(err, nil)
	a.Equal(popCard, *NewCard(Clubs, Two))
	a.Equal(d.Count(), 52)

	for i := 0; i < 52; i++ {
		_, err := d.Pop()
		a.Equal(err, nil)
	}

	_, err = d.Pop()
	a.NotEqual(err, nil)
}

func TestShuffledDeck(t *testing.T) {
	d := NewDeck(func(min, max int) int {
		if min == max {
			return min
		}
		return min + 1
	})

	a := assert.New(t)
	popCard, err := d.Pop()
	a.Equal(err, nil)
	a.Equal(popCard, *NewCard(Clubs, Three))
}
