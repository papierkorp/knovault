package server

import (
    "github.com/labstack/echo/v4"
    "knovault/internal/globals"
)

func handleHome(c echo.Context) error {
    component, err := globals.GetAssetManager().GetCurrentTheme().Home()
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handlePlayground(c echo.Context) error {
    component, err := globals.GetAssetManager().GetCurrentTheme().Playground()
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handleHelp(c echo.Context) error {
    component, err := globals.GetAssetManager().GetCurrentTheme().Help()
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handleSettings(c echo.Context) error {
    component, err := globals.GetAssetManager().GetCurrentTheme().Settings()
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handleSearch(c echo.Context) error {
    component, err := globals.GetAssetManager().GetCurrentTheme().Search()
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handleDocsRoot(c echo.Context) error {
    component, err := globals.GetAssetManager().GetCurrentTheme().DocsRoot()
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handleDocs(c echo.Context) error {
    title := c.Param("title")
    component, err := globals.GetAssetManager().GetCurrentTheme().Docs(title)
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handlePlugins(c echo.Context) error {
    component, err := globals.GetAssetManager().GetCurrentTheme().Plugins()
    if err != nil {
        return err
    }
    return _render(c, component)
}