package mighty

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBundle(t *testing.T) {
	b := NewBundle(func(min, max int) int {
		return min
	})
	a := assert.New(t)
	a.Equal(b.Count(), 53)

	popCard, err := b.Pop()
	a.Equal(err, nil)
	a.Equal(popCard, *NewCard(Clubs, Two))
	a.Equal(b.Count(), 52)

	for i := 0; i < 52; i++ {
		_, err := b.Pop()
		a.Equal(err, nil)
	}

	_, err = b.Pop()
	a.NotEqual(err, nil)
}

func TestShuffledBundle(t *testing.T) {
	b := NewBundle(func(min, max int) int {
		if min == max {
			return min
		}
		return min + 1
	})

	a := assert.New(t)
	popCard, err := b.Pop()
	a.Equal(err, nil)
	a.Equal(popCard, *NewCard(Clubs, Three))
}
