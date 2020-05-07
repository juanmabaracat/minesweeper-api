package services

import (
	"errors"
	"github.com/juanmabaracat/minesweeper-api/src/domain/game"
	"github.com/juanmabaracat/minesweeper-api/src/repository/db/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewService(t *testing.T) {
	dbMock := &mocks.Repository{}
	service := NewGameService(dbMock)
	assert.NotNil(t, service)
}

func TestService_CreateError(t *testing.T) {
	dbMock := &mocks.Repository{}
	dbMock.On("Create").Return(errors.New("database error"))

	service := NewGameService(dbMock)
	assert.NotNil(t, service)

	newGame, err := service.Create(game.CreateRequest{1, 10, 20, 5})
	assert.Nil(t, newGame)
	assert.NotNil(t, err)
	assert.EqualValues(t, "database error", err.Error())
}
