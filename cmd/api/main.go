package main

import (
	"net/http"

	"github.com/hamzah-mifta/mkp-test-backend/config"
	httpDelivery "github.com/hamzah-mifta/mkp-test-backend/delivery/http"
	"github.com/hamzah-mifta/mkp-test-backend/infrastructure/datastore"
	pgsqlRepository "github.com/hamzah-mifta/mkp-test-backend/repository/pgsql"
	"github.com/hamzah-mifta/mkp-test-backend/usecase"
	"github.com/hamzah-mifta/mkp-test-backend/utils/crypto"
	"github.com/hamzah-mifta/mkp-test-backend/utils/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load config
	configApp := config.LoadConfig()

	// Setup infra
	db, err := datastore.NewDatabase(configApp.DatabaseURL)
	if err != nil {
		panic(err)
	}

	// Setup repository
	userRepo := pgsqlRepository.NewPgsqlUserRepository(db)

	// Setup service
	cryptoSvc := crypto.NewCryptoService()
	jwtSvc := jwt.NewJWTService(configApp.JWTSecretKey)

	// Setup usecase
	authUsecase := usecase.NewAuthUsecase(userRepo, cryptoSvc, jwtSvc)

	// Setup app middleware

	// Setup route and middleware
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	// Setup handler
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "PONG!")
	})
	httpDelivery.NewAuthHandler(e, authUsecase)

	// Start server
	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal("shutting down the server")
	}
}
