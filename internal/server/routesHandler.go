package server

import (
	"gowiki/internal/filemanager"
	"gowiki/internal/themes"
	"net/http"
	"github.com/labstack/echo/v4"
)


func handleHome(c echo.Context) error {
	component, err := themes.GetCurrentTheme().Home()
	if err != nil {
		return err
	}
	return _render(c, component)
}

func handlePlayground(c echo.Context) error {
	content := filemanager.ParseMarkdownToHtml("example_markdown.md")

    component, err := themes.GetCurrentTheme().Playground(content)
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handleHelp(c echo.Context) error {
	component, err := themes.GetCurrentTheme().Help()
	if err != nil {
		return err
	}
	return _render(c, component)
}

func handleSettings(c echo.Context) error {
	component, err := themes.GetCurrentTheme().Settings()
	if err != nil {
		return err
	}
	return _render(c, component)
}

func handleSearch(c echo.Context) error {
	component, err := themes.GetCurrentTheme().Search()
	if err != nil {
		return err
	}
	return _render(c, component)
}

func handleDocsRoot(c echo.Context) error {
	component, err := themes.GetCurrentTheme().DocsRoot()
	if err != nil {
		return err
	}
	return _render(c, component)
}

func handleDocs(c echo.Context) error {
	title := c.Param("title")
	component, err := themes.GetCurrentTheme().Docs(title)
	if err != nil {
		return err
	}
	return _render(c, component)
}

func handleChangeTheme(c echo.Context) error {
	var request struct {
		Theme string `json:"theme"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]bool{"success": false})
	}

	err := themes.SetCurrentTheme(request.Theme)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]bool{"success": false})
	}

	return c.JSON(http.StatusOK, map[string]bool{"success": true})
}