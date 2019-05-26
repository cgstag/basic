package main

import (
	"basic/api"
	"basic/config"
	"basic/pkg/account"
	"basic/pkg/dynamo"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func main() {

	// Load Config
	configuration := config.MustLoadConfig()

	// Load DB
	db, err := dynamo.NewDB(configuration.DB)
	if err != nil {
		log.Panic(err)
	}

	// Initialize Logger
	log = zap.NewExample().Sugar()
	defer log.Sync()

	// Initialize Echo
	e := echo.New()

	// Initialize Middleware
	e.Use(middleware.Recover())
	router := e.Group("/v1")

	env := &api.Env{Db: db, Log: log}

	// Serve Routes
	account.ServeResources(env, router)

	// Healthcheck
	e.GET("/", func(c echo.Context) error {
		log.Infow("Calling Hello World...")
		return c.JSON(http.StatusOK, "Hello, World!")
	})
	e.Match([]string{"GET", "HEAD"}, "/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok")
	})

	// Start server
	address := fmt.Sprintf("%v:%v", configuration.Host, configuration.Port)
	e.Logger.Fatal(e.Start(address))
}
