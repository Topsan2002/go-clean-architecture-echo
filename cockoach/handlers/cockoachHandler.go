package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type CockroachHandler interface {
	DetectCockroach(c echo.Context) error
}

type CockroachGinHandler interface {
	CreateCockroch(c *gin.Context)
}
