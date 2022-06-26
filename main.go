package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osamamosaad/paytabs/config"
	"github.com/osamamosaad/paytabs/controllers"
	"github.com/osamamosaad/paytabs/migrations"
	"github.com/osamamosaad/paytabs/storage"
)

func main() {
	migrations.LoadAccounts()

	// Run storage
	storage := storage.New()

	// Configure frameworks
	e := echo.New()
	e.HideBanner = true

	// Register routes
	registerRoutes(e, storage)

	// Run server
	e.Logger.Fatal(e.Start(
		fmt.Sprintf("%v:%v",
			config.SERVER_URL,
			config.SERVER_PORT,
		),
	))
}

func registerRoutes(e *echo.Echo, storage *storage.Storage) {
	// Greeting message
	e.GET("/", hello())

	// Account routes
	e.GET("/accounts", controllers.ListAccounts(storage))
	e.GET("/accounts/:id", controllers.GetAccount(storage))

	// transactions routes
	e.POST("/transactions", controllers.Transaction(storage))
}

func hello() func(echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to PayTabs Api")
	}
}
