package server

import (
	"github.com/labstack/echo/v4"
	"gowiki/internal/parser"
	"gowiki/internal/templates"
	"log"
)

func setupRoutes(e *echo.Echo) {
	e.GET("/", handleHome)
	e.GET("/help", handleHelp)
	e.GET("/settings", handleSettings)
	e.GET("/search", handleSearch)
	e.GET("/docs/*", handleDocs)
}

func handleHome(c echo.Context) error {
	content, err := parser.ReadMarkdownFile("example_markdown.md")
	if err != nil {
		log.Printf("Error reading markdown file: %v", err)
		return err
	}

	err = templates.Home(content).Render(c.Request().Context(), c.Response().Writer)
	if err != nil {
		log.Printf("Error rendering template: %v", err)
		return err
	}

	return nil
}

func handleHelp(c echo.Context) error {
	// Assuming you have a Help template
	component := templates.Help()
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func handleSettings(c echo.Context) error {
	// Assuming you have a Help template
	component := templates.Settings()
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func handleSearch(c echo.Context) error {
	// Assuming you have a Help template
	component := templates.Search()
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func handleDocs(c echo.Context) error {
	// Assuming you have a Help template
	component := templates.Docs()
	return component.Render(c.Request().Context(), c.Response().Writer)
}
