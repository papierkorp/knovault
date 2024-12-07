// internal/server/routesHandler.go
package server

import (
    "github.com/labstack/echo/v4"
    "knovault/internal/globals"
)

func handleHome(c echo.Context) error {
    component, err := globals.GetThemeManager().GetCurrentTheme().Home()
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handlePlayground(c echo.Context) error {
    component, err := globals.GetThemeManager().GetCurrentTheme().Playground()
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handleHelp(c echo.Context) error {
    component, err := globals.GetThemeManager().GetCurrentTheme().Help()
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handleSettings(c echo.Context) error {
    component, err := globals.GetThemeManager().GetCurrentTheme().Settings()
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handleSearch(c echo.Context) error {
    component, err := globals.GetThemeManager().GetCurrentTheme().Search()
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handleDocsRoot(c echo.Context) error {
    component, err := globals.GetThemeManager().GetCurrentTheme().DocsRoot()
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handleDocs(c echo.Context) error {
    title := c.Param("title")
    component, err := globals.GetThemeManager().GetCurrentTheme().Docs(title)
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handlePlugins(c echo.Context) error {
    component, err := globals.GetThemeManager().GetCurrentTheme().Plugins()
    if err != nil {
        return err
    }
    return _render(c, component)
}