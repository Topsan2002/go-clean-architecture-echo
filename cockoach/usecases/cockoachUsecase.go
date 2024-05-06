package usecases

import "go-clean-ex/cockoach/models"

type CockroachUseCase interface {
	CockroachDataProcessing(in *models.AddCockroachData) error
}
