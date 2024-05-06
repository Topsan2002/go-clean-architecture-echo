package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	handlersCockroach "go-clean-ex/cockoach/handlers"
	repoCockroah "go-clean-ex/cockoach/repository"
	usecasesCockroach "go-clean-ex/cockoach/usecases"
	"go-clean-ex/config"
	"go-clean-ex/database"
)

type echoServer struct {
	app    *echo.Echo
	db     database.Database
	config *config.Config
}

// initializeCockroachHttpHandler implements Server.
func (e *echoServer) initializeCockroachHttpHandler() {
	repo := repoCockroah.NewCockroachPostgresRepository(e.db)
	fmc := repoCockroah.NewCockroachFCMMessaging()

	usecases := usecasesCockroach.NewCockoachUsecaseImp(repo, fmc)

	httpHandler := handlersCockroach.NewCockoachHttpHandler(usecases)

	cockroachRouters := e.app.Group("v1/cockroach")
	cockroachRouters.POST("", httpHandler.DetectCockroach)

}

func NewEchoServer(config *config.Config, db database.Database) Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)
	return &echoServer{
		app:    echoApp,
		db:     db,
		config: config,
	}
}

// Start implements Server.
func (e *echoServer) Start() {
	e.app.Use(middleware.Recover())
	e.app.Use(middleware.Logger())

	e.app.GET("v1/health", func(c echo.Context) error {
		return c.JSON(200, "OK")
	})

	e.initializeCockroachHttpHandler()

	serverUrl := fmt.Sprintf(":%d", e.config.Server.Port)
	e.app.Logger.Fatal(e.app.Start(serverUrl))
}
