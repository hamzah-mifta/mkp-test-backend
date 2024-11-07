package main

import (
	"fmt"
	"net/http"

	"github.com/hamzah-mifta/mkp-test-backend/config"
	"github.com/hamzah-mifta/mkp-test-backend/infrastructure/datastore"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load config
	configApp := config.LoadConfig()

	// Setup infra
	dbInstance, err := datastore.NewDatabase(configApp.DatabaseURL)
	if err != nil {
		panic(err)
	}

	fmt.Println(dbInstance)

	// Setup repository

	// Setup service

	// Setup usecase

	// Setup app middleware

	// Setup route and middleware
	e := echo.New()
	e.Use(middleware.CORS())

	// Setup handler
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "PONG!")
	})

	// Start server
	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal("shutting down the server")
	}
}
