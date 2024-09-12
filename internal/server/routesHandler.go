package server

import (
	"gowiki/internal/filemanager"
	"gowiki/internal/templates"
	"gowiki/internal/themes"
	"github.com/labstack/echo/v4"
)


func handleHome(c echo.Context) error {
	component, err := themes.GetCurrentTheme().Home()
	if err != nil {
		return err
	}
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func handlePlayground(c echo.Context) error {
	content := filemanager.ParseMarkdownToHtml("example_markdown.md")

	templates.Playground(content).Render(c.Request().Context(), c.Response().Writer)

	return nil
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

func handleDocsRoot(c echo.Context) error {
	component := templates.DocsRoot()
	return _render(c, component)
}

func handleDocs(c echo.Context) error {
	title := c.Param("title")
	component := templates.Docs(title)
	return _render(c, component)
}
