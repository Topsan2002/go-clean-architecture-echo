package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-clean-ex/cockoach/models"
	"go-clean-ex/cockoach/usecases"
)

type cockoachGinHttpHandler struct {
	usecase usecases.CockroachUseCase
}

func NewCockoachGinHttpHandler(usecase usecases.CockroachUseCase) CockroachGinHandler {
	return &cockoachGinHttpHandler{usecase: usecase}
}

// CreateCockroch implements CockroachGinHandler.
func (s *cockoachGinHttpHandler) CreateCockroch(c *gin.Context) {
	reqBody := new(models.AddCockroachData)

	if err := c.Bind(&reqBody); err != nil {
		responseGin(c, http.StatusBadRequest, "Bad Reqwuest")
	}

	if err := s.usecase.CockroachDataProcessing(reqBody); err != nil {
		responseGin(c, http.StatusInternalServerError, "Process Data Failed")
	}

	responseGin(c, http.StatusOK, "Success ")

}
