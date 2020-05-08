package services

import (
	"errors"
	"github.com/juanmabaracat/minesweeper-api/src/domain/game"
	"github.com/juanmabaracat/minesweeper-api/src/repository/db/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewService(t *testing.T) {
	dbMock := &mocks.Repository{}
	service := NewGameService(dbMock)
	assert.NotNil(t, service)
}

func TestService_CreateError(t *testing.T) {
	dbMock := &mocks.Repository{}
	dbMock.On("Create", mock.Anything).Return(errors.New("database error"))

	service := NewGameService(dbMock)
	assert.NotNil(t, service)

	newGame, err := service.Create(game.CreateRequest{1, 10, 20, 5})
	assert.Nil(t, newGame)
	assert.NotNil(t, err)
	assert.EqualValues(t, "database error", err.Error())
}

func TestService_Create(t *testing.T) {
	dbMock := &mocks.Repository{}
	dbMock.On("Create", mock.Anything).Return(nil)

	service := NewGameService(dbMock)
	assert.NotNil(t, service)

	newGame, err := service.Create(game.CreateRequest{1, 10, 20, 5})
	assert.Nil(t, err)
	assert.NotNil(t, newGame)
	assert.EqualValues(t, game.PLAYING, newGame.Status)
	assert.EqualValues(t, 1, newGame.PlayerId)
	assert.EqualValues(t, 10, newGame.Board.Width)
	assert.EqualValues(t, 20, newGame.Board.Height)
	assert.EqualValues(t, 5, newGame.Board.Mines)
}
