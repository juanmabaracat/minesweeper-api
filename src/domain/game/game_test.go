package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameConstants(t *testing.T) {
	assert.EqualValues(t, "playing", PLAYING)
	assert.EqualValues(t, "won", WON)
	assert.EqualValues(t, "defeated", DEFEATED)
}

func TestCreateGame(t *testing.T) {
	game := NewGame()
	err := game.Create(123, 10, 20, 5)

	assert.Nil(t, err)
	assert.NotNil(t, game)

	assert.NotNil(t, game.Board)
	assert.EqualValues(t, 200, len(game.Board.Cells))
	assert.EqualValues(t, 123, game.PlayerId)
	assert.EqualValues(t, PLAYING, game.Status)
}
