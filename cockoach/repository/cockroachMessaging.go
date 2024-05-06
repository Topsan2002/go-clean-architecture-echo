package repository

import "go-clean-ex/cockoach/entity"

type CockroachMessaging interface {
	PushNotification(m *entity.CockroachPushNotificatinDto) error
}
