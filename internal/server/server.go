package server

import (
	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()

	// Setup routes
	setupRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}