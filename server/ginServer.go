package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	handlersCockroach "go-clean-ex/cockoach/handlers"
	repoCockroah "go-clean-ex/cockoach/repository"
	usecasesCockroach "go-clean-ex/cockoach/usecases"
	"go-clean-ex/config"
	"go-clean-ex/database"
)

type ginServer struct {
	app    *gin.Engine
	db     database.Database
	config *config.Config
}

func NewGinServerDebug(db database.Database, config *config.Config) Server {
	gin.SetMode(gin.DebugMode)
	app := gin.Default()
	return &ginServer{
		app:    app,
		db:     db,
		config: config,
	}
}

// Start implements Server.
func (g *ginServer) Start() {

	g.app.GET("v1/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})

	g.initializeCockroachHttpHandler()

	serverUrl := fmt.Sprintf(":%d", g.config.Server.Port)
	g.app.Run(serverUrl)
}

func (g *ginServer) initializeCockroachHttpHandler() {
	repo := repoCockroah.NewCockroachPostgresRepository(g.db)
	fmc := repoCockroah.NewCockroachFCMMessaging()

	usecases := usecasesCockroach.NewCockoachUsecaseImp(repo, fmc)
	httpHandler := handlersCockroach.NewCockoachGinHttpHandler(usecases)

	cockroachRouters := g.app.Group("v1/cockroach")
	cockroachRouters.POST("", httpHandler.CreateCockroch)

	// g.app.Group("/")
}
