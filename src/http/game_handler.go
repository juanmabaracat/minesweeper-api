package http

import (
	"github.com/gin-gonic/gin"
	"github.com/juanmabaracat/minesweeper-api/src/domain/game"
	"github.com/juanmabaracat/minesweeper-api/src/errors"
	"github.com/juanmabaracat/minesweeper-api/src/services"
	"net/http"
	"strconv"
)

type GameHandler interface {
	Create(*gin.Context)
	Update(*gin.Context)
	GetById(*gin.Context)
	GetAll(*gin.Context)
}

type gameHandler struct {
	service services.GameService
}

func NewGameHandler(service services.GameService) GameHandler {
	return gameHandler{service}
}

func (handler gameHandler) Create(c *gin.Context) {
	var createRequest game.CreateRequest
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status(), restError)
		return
	}

	newGame, err := handler.service.Create(createRequest)
	if err != nil {
		restErr, ok := err.(errors.RestErr)
		if !ok {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(http.StatusCreated, newGame)
}

func (handler gameHandler) Update(c *gin.Context) {
	id, err := getGameId(c.Param("game_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var game game.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status(), restError)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch
	game.Id = id
	if err := handler.service.Update(isPartial, &game); err != nil {
		restErr, ok := err.(errors.RestErr)
		if !ok {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(http.StatusOK, game)
}

func (handler gameHandler) GetById(c *gin.Context) {
	id, err := getGameId(c.Param("game_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	game, getErr := handler.service.GetById(id)
	if getErr != nil {
		restErr, ok := getErr.(errors.RestErr)
		if !ok {
			c.JSON(http.StatusInternalServerError, restErr.Error())
			return
		}
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(http.StatusOK, game)
}

func (handler gameHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "not implemented!")
}

func getGameId(gameIdParam string) (uint64, errors.RestErr) {
	id, parseErr := strconv.ParseUint(gameIdParam, 10, 64)
	if parseErr != nil {
		return 0, errors.NewBadRequestError("invalid game id")
	}
	return id, nil
}
