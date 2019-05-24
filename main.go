package main

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
)

func main() {

	// Initialize Router
	router := echo.New()

	// Initialize Logger
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	// Construct Routes
	router.GET("/", func(c echo.Context) error {
		greeting := "Hello, World "
		logger.Infof("Route '/' found, returning Greetings %s", greeting )
		return c.String(http.StatusOK, greeting )
	})

	// Start Router
	router.Logger.Fatal(router.Start(":8080"))
}
