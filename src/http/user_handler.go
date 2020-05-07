package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// The idea was to use JWT for Authentication.
type UserHandler interface {
	Login(*gin.Context)
}

type userHandler struct {
}

func NewUserHandler() UserHandler {
	return &userHandler{}
}

func (u *userHandler) Login(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "not implemented!")
}
