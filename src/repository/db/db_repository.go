package db

import (
	"fmt"
	"github.com/juanmabaracat/minesweeper-api/src/domain/game"
	"github.com/juanmabaracat/minesweeper-api/src/errors"
	"math/rand"
)

type Repository interface {
	Create(*game.Game) error
	Update(*game.Game) error
	GetById(uint64) (*game.Game, error)
}

func NewDbRepository() Repository {
	gamesTable := make(map[uint64]*game.Game)
	playerGamesTable := make(map[uint64][]uint64)
	return &repository{gamesTable, playerGamesTable}
}

type repository struct {
	gamesTable       map[uint64]*game.Game
	playerGamesTable map[uint64][]uint64
}

func (r repository) Create(game *game.Game) error {
	game.Id = generateId()
	r.gamesTable[game.Id] = game
	return nil
}

func (r repository) Update(g *game.Game) error {
	r.gamesTable[g.Id] = g
	return nil
}

func (r repository) GetById(id uint64) (*game.Game, error) {
	game := r.gamesTable[id]
	if game == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("game with id %d not found", id))
	}
	return game, nil
}

func generateId() uint64 {
	// just a dump id generator for the exercise, could be an UUID generator or another approach
	return rand.Uint64()
}
