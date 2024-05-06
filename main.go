package main

import (
	"go-clean-ex/config"
	"go-clean-ex/database"
	"go-clean-ex/server"
)

func main() {
	config := config.GetConfig()
	db := database.NewPostgresDatabase(config)
	// server.NewEchoServer(config, db).Start()

	// go func() {
	// 	server.NewEchoServer(config, db).Start()

	// }()
	server.NewGinServerDebug(db, config).Start()

}
