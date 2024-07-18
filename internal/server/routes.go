package server

import (
	"github.com/labstack/echo/v4"
	"gowiki/internal/templates"
)

func setupRoutes(e *echo.Echo) {
	e.GET("/", handleHome)
	e.GET("/help", handleHelp)
}

func handleHome(c echo.Context) error {
	component := templates.Home()
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func handleHelp(c echo.Context) error {
	// Assuming you have a Help template
	component := templates.Help()
	return component.Render(c.Request().Context(), c.Response().Writer)
}