package main

import (
	"basic/config"
	"basic/pkg/account"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func main() {

	// Load Config
	config := config.MustLoadConfig()

	// Initialize Logger
	log = zap.NewExample().Sugar()
	defer log.Sync()

	// Initialize Echo
	e := echo.New()

	// Initialize Middleware
	e.Use(middleware.Recover())
	router := e.Group("/v1")

	// Serve Routes
	account.ServeResources(router, log)

	// Healthcheck
	e.GET("/", func(c echo.Context) error {
		log.Infow("Calling Hello World...")
		return c.JSON(http.StatusOK, "Hello, World!")
	})
	e.Match([]string{"GET", "HEAD"}, "/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok")
	})

	// Start server
	address := fmt.Sprintf("%v:%v", config.Host, config.Port)
	e.Logger.Fatal(e.Start(address))
}
