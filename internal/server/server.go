package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	e := echo.New()

	e.Static("/", "static")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Setup routes
	setupRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
