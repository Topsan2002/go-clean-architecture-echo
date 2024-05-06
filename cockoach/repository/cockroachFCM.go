package repository

import (
	"github.com/labstack/gommon/log"

	"go-clean-ex/cockoach/entity"
)

type cockoachFCMMessaging struct{}

func NewCockroachFCMMessaging() CockroachMessaging {
	return &cockoachFCMMessaging{}
}

// PushNotification implements CockroachMessaging.
func (c *cockoachFCMMessaging) PushNotification(m *entity.CockroachPushNotificatinDto) error {

	log.Debugf("Pushed FCM notificatin with data : %v", m)
	return nil
}
