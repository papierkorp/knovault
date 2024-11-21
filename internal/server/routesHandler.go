package server

import (
	"knovault/internal/themes"
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
    component, err := themes.GetCurrentTheme().Playground()
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


func handlePlugins(c echo.Context) error {
	component, err := themes.GetCurrentTheme().Plugins()
	if err != nil {
		return err
	}
	return _render(c, component)
}