package entity

import "time"

type (
	InsertCockroachDto struct {
		Id       uint32 `gorm:"primaryKey;autoIncrement" json:"id"`
		Amount   uint32 `json:"amount"`
		CreateAt time.Time
	}

	Cockroach struct {
		Id       uint32 `json:"id"`
		Amount   uint32 `json:"amount"`
		CreateAt time.Time
	}

	CockroachPushNotificatinDto struct {
		Title        string `json:"title"`
		Amount       uint32 `json:"amount"`
		ReportedTime string `json:"createdAt"`
	}
)
