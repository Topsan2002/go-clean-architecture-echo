package main

import (
	"go-clean-ex/cockoach/entity"
	"go-clean-ex/config"
	"go-clean-ex/database"
)

func main() {
	config := config.GetConfig()
	db := database.NewPostgresDatabase(config)
	cockroachMigrate(db)
}

func cockroachMigrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entity.Cockroach{})
	db.GetDb().CreateInBatches([]entity.Cockroach{
		{Amount: 1},
		{Amount: 2},
		{Amount: 3},
		{Amount: 4},
	}, 10)
}
