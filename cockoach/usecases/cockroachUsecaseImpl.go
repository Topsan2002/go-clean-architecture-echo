package usecases

import (
	"time"

	"go-clean-ex/cockoach/entity"
	"go-clean-ex/cockoach/models"
	"go-clean-ex/cockoach/repository"
)

type cockoachUserCaseImp struct {
	cockoachRepository repository.CockroachRepository
	cockroachMessaging repository.CockroachMessaging
}

func NewCockoachUsecaseImp(cockoachRepository repository.CockroachRepository, cockroachMessaging repository.CockroachMessaging) CockroachUseCase {
	return &cockoachUserCaseImp{cockoachRepository: cockoachRepository, cockroachMessaging: cockroachMessaging}
}

// CockroachDataProcessing implements CockroachUseCase.
func (c *cockoachUserCaseImp) CockroachDataProcessing(in *models.AddCockroachData) error {
	// panic("unimplemented")

	insertData := &entity.InsertCockroachDto{Amount: in.Amount}
	if err := c.cockoachRepository.InsertCockoachData(insertData); err != nil {
		return err
	}

	pushData := &entity.CockroachPushNotificatinDto{
		Title:        "Cockroach Detected 🪳 !!!",
		Amount:       in.Amount,
		ReportedTime: time.Now().Local().Format("2006-01-02 15:04:00"),
	}

	if err := c.cockroachMessaging.PushNotification(pushData); err != nil {
		return err
	}

	return nil

}
