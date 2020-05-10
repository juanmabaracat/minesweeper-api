package services

import (
	"github.com/juanmabaracat/minesweeper-api/src/domain/game"
	"github.com/juanmabaracat/minesweeper-api/src/repository/db"
	"github.com/juanmabaracat/minesweeper-api/src/utils"
)

type GameService interface {
	Create(game.CreateRequest) (*game.Game, error)
	Update(bool, *game.Game) error
	GetById(uint64) (*game.Game, error)
}

type gameService struct {
	repository db.Repository
}

func NewGameService(repo db.Repository) GameService {
	return &gameService{repository: repo}
}

func (s *gameService) Create(request game.CreateRequest) (*game.Game, error) {
	newGame := game.NewGame()
	err := newGame.Create(request.PlayerId, request.Width, request.Height, request.Mines)
	if err != nil {
		return nil, err
	}

	if createErr := s.repository.Create(newGame); createErr != nil {
		return nil, createErr
	}

	return newGame, nil
}

func (s *gameService) Update(isPartial bool, gameRequest *game.Game) error {
	current, getError := s.repository.GetById(gameRequest.Id)
	if getError != nil {
		return getError
	}

	if isPartial {
		if gameRequest.Status != "" {
			current.Status = gameRequest.Status
		}
		if gameRequest.Board != nil {
			current.Board = gameRequest.Board
		}
	} else {
		current.Status = gameRequest.Status
		current.Board = gameRequest.Board
	}

	if gameRequest.Status == game.WON || gameRequest.Status == game.DEFEATED {
		current.DateFinished = utils.GetNowString()
	}

	*gameRequest = *current
	return s.repository.Update(current)
}

func (s *gameService) GetById(id uint64) (*game.Game, error) {
	return s.repository.GetById(id)
}
