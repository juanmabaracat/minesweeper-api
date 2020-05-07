package app

import (
	"github.com/gin-gonic/gin"
	"github.com/juanmabaracat/minesweeper-api/src/http"
	"github.com/juanmabaracat/minesweeper-api/src/repository/db"
	"github.com/juanmabaracat/minesweeper-api/src/services"
	"math/rand"
	"time"
)

var (
	router = gin.Default()
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func StartApplication() {
	gameHandler := http.NewGameHandler(services.NewGameService(db.NewDbRepository()))
	userHandler := http.NewUserHandler()

	router.GET("/ping", http.Ping)
	router.GET("/minesweeper/api/game", gameHandler.GetAll)
	router.GET("/minesweeper/api/game/:game_id", gameHandler.GetById)
	router.POST("/minesweeper/api/game", gameHandler.Create)
	router.PUT("/minesweeper/api/game/:game_id", gameHandler.Update)
	router.PATCH("/minesweeper/api/game/:game_id", gameHandler.Update)

	router.POST("/minesweeper/api/user/login", userHandler.Login)

	router.Run(":8080")
}
