package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gowiki/internal/filemanager"
	"gowiki/internal/templates"
	"log"
)

func setupRoutes(e *echo.Echo) {
	e.GET("/", handleHome)
	e.GET("/help", handleHelp)
	e.GET("/settings", handleSettings)
	e.GET("/search", handleSearch)
	e.GET("/docs/*", handleDocs)
	e.GET("/playground", handlePlayground)
}

func handleHome(c echo.Context) error {
	content, err := filemanager.ParseMarkdownToHtml("example_markdown.md")

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

func handlePlayground(c echo.Context) error {
	all_files, err := filemanager.GetAllFiles()
	fmt.Println("all_files: ", all_files, " err: ", err)
	component := templates.Playground()
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func handleHelp(c echo.Context) error {
	component := templates.Help()
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func handleSettings(c echo.Context) error {
	component := templates.Settings()
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func handleSearch(c echo.Context) error {
	component := templates.Search()
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func handleDocs(c echo.Context) error {
	component := templates.Docs()
	return component.Render(c.Request().Context(), c.Response().Writer)
}
