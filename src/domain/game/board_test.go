package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_createBoardInvalidWidth(t *testing.T) {
	board := Board{}
	err := board.Create(-3, 10, 5)
	assert.NotNil(t, err)
	assert.EqualValues(t, "bad_request", err.Error())
}

func Test_createBoardInvalidHeight(t *testing.T) {
	board := Board{}
	err := board.Create(10, 0, 5)
	assert.NotNil(t, err)
	assert.EqualValues(t, "bad_request", err.Error())
}

func Test_createBoardInvalidMines(t *testing.T) {
	board := Board{}
	err := board.Create(10, 20, 0)
	assert.NotNil(t, err)
	assert.EqualValues(t, "bad_request", err.Error())
}

func Test_createBoardValidValues(t *testing.T) {
	board := Board{}
	err := board.Create(10, 20, 5)
	assert.Nil(t, err)
	assert.NotNil(t, board)
	assert.EqualValues(t, 200, len(board.Cells))
}

func Test_plantMines(t *testing.T) {
	board := Board{}
	err := board.Create(10, 20, 10)

	assert.NotNil(t, board.Cells)
	assert.Nil(t, err)

	counter := 0
	for _, cell := range board.Cells {
		if cell.HasMine {
			counter++
		}
	}

	assert.EqualValues(t, 10, counter)
}
