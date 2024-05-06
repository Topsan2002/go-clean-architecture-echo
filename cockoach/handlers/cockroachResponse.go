package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type baseResponse struct {
	Message string `json:"message"`
}

func response(e echo.Context, responseCode int, message string) error {
	return e.JSON(responseCode, &baseResponse{
		Message: message,
	})
}

func responseGin(g *gin.Context, responseCode int, message string) {
	g.JSON(responseCode, gin.H{
		"message": message,
	})

}
