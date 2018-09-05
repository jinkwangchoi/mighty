package card

import "testing"
import "github.com/stretchr/testify/assert"

func TestCard(t *testing.T) {
	joker := NewJoker()

	assert := assert.New(t)
	assert.Equal(joker.IsJoker(), true)

	clubTwo := NewCard(Clubs, Two)
	assert.Equal(clubTwo.IsJoker(), false)
	assert.Equal(clubTwo.Suit, Clubs)
	assert.Equal(clubTwo.Number, Two)
}
