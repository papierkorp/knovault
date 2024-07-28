package server

import (
	"github.com/labstack/echo/v4"
	"gowiki/internal/filemanager"
	"gowiki/internal/templates"
	"log"
)

type Article struct {
	path  String
	title String
}

func handleHome(c echo.Context) error {
	content, err := filemanager.ParseMarkdownToHtml("/mnt/c/develop/gowiki/data/example_markdown.md")

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
	component := templates.Playground()
	return _render(c, component)
}

func handleHelp(c echo.Context) error {
	component := templates.Help()
	return _render(c, component)
}

func handleSettings(c echo.Context) error {
	component := templates.Settings()
	return _render(c, component)
}

func handleSearch(c echo.Context) error {
	component := templates.Search()
	return _render(c, component)
}

func handleDocs(c echo.Context, article Article) error {
	component := templates.Docs(article)
	return _render(c, component)
}
