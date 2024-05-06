package handlers

import "github.com/labstack/echo/v4"

type baseResponse struct {
	Message string `json:"message"`
}

func response(e echo.Context, responseCode int, message string) error {
	return e.JSON(responseCode, &baseResponse{
		Message: message,
	})
}
