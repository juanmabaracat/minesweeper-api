package game

import (
	"github.com/juanmabaracat/minesweeper-api/src/errors"
	"math/rand"
)

type Board struct {
	Width  int   `json:"width"`
	Height int   `json:"height"`
	Mines  int   `json:"mines"`
	Cells  Cells `json:"cells"`
}

func (b *Board) Create(width int, height int, mines int) error {
	b.Width = width
	b.Height = height
	b.Mines = mines

	err := b.createCells()
	if err != nil {
		return err
	}

	if err := b.plantMines(); err != nil {
		return err
	}

	return nil
}

func (b *Board) createCells() error {
	if b.Width <= 0 || b.Height <= 0 {
		return errors.NewBadRequestError("invalid Board size")
	}

	b.Cells = Cells{}
	for i := 0; i < b.Width; i++ {
		for j := 0; j < b.Height; j++ {
			b.Cells = append(b.Cells, Cell{
				X:           i,
				Y:           j,
				HasMine:     false,
				WasRevealed: false,
				Mark:        NONE,
			})
		}
	}
	return nil
}

func (b *Board) plantMines() error {
	if b.Mines <= 0 || b.Mines > len(b.Cells) {
		return errors.NewBadRequestError("invalid mines quantity")
	}

	for i := 0; i < b.Mines; i++ {
		planted := false
		for !planted {
			randomCell := rand.Intn(len(b.Cells))
			if !b.Cells[randomCell].HasMine {
				b.Cells[randomCell].HasMine = true
				planted = true
			}
		}
	}
	return nil
}
