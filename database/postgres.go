package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-clean-ex/config"
)

type postgresBatabase struct {
	Db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *postgresBatabase
)

func NewPostgresDatabase(config *config.Config) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			config.Db.Host,
			config.Db.User,
			config.Db.Password,
			config.Db.DBName,
			config.Db.Port,
			config.Db.SSLMode,
			config.Db.TimeZone,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("fail connect database")
		}

		dbInstance = &postgresBatabase{Db: db}

	})

	return dbInstance
}

// GetDb implements Database.
func (p *postgresBatabase) GetDb() *gorm.DB {
	return dbInstance.Db
}
