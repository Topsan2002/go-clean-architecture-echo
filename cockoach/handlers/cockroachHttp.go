package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"go-clean-ex/cockoach/models"
	"go-clean-ex/cockoach/usecases"
)

type cockoachHttpHandler struct {
	cockoachUsecase usecases.CockroachUseCase
}

func NewCockoachHttpHandler(cockoachUsecase usecases.CockroachUseCase) CockroachHandler {
	return &cockoachHttpHandler{cockoachUsecase: cockoachUsecase}
}

// DetectCockroach implements CockroachHandler.
func (h *cockoachHttpHandler) DetectCockroach(c echo.Context) error {
	reqBody := new(models.AddCockroachData)

	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body %v", err)
		return response(c, http.StatusBadRequest, "Bad Request")
	}

	if err := h.cockoachUsecase.CockroachDataProcessing(reqBody); err != nil {
		log.Errorf("Error binding request body %v", err)
		return response(c, http.StatusInternalServerError, "Process Data Failed")
	}

	return response(c, http.StatusOK, "Success ")
}
