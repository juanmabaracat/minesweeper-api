package game

import (
	"github.com/juanmabaracat/minesweeper-api/src/utils"
)

type Status string

const (
	WON      Status = "won"
	PLAYING  Status = "playing"
	DEFEATED Status = "defeated"
)

func NewGame() *Game {
	return &Game{}
}

type Game struct {
	Id           uint64 `json:"id"`
	PlayerId     uint64 `json:"player_id"`
	Status       Status `json:"status"`
	Board        *Board `json:"board"`
	DateCreated  string `json:"date_created"`
	DateFinished string `json:"date_finished"`
}

func (g *Game) Create(playerId uint64, width int, height int, mines int) error {
	board := Board{}
	if err := board.Create(width, height, mines); err != nil {
		return err
	}

	g.PlayerId = playerId
	g.Status = PLAYING
	g.Board = &board
	g.DateCreated = utils.GetNowString()
	return nil
}
