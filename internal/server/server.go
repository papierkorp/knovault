// internal/server/server.go
package server

import (
    "log"
    "knovault/internal/pluginManager"
    "knovault/internal/themeManager"
    "knovault/internal/globals"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func Start() {
    e := echo.New()

    // Initialize plugin manager
    pm := pluginManager.NewPluginManager()
    globals.SetPluginManager(pm)
    if err := pm.Initialize(); err != nil {
        log.Fatalf("Failed to initialize plugin manager: %v", err)
    }

    // Initialize theme manager
    tm := themeManager.NewThemeManager()
    globals.SetThemeManager(tm)
    if err := tm.Initialize(); err != nil {
        log.Fatalf("Failed to initialize theme manager: %v", err)
    }

    // Set default theme
    err := tm.SetCurrentTheme("defaultTheme")
    if err != nil {
        log.Fatalf("Failed to set default theme: %v", err)
    }

    e.Static("/static", "static")
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    setupRoutes(e)

    e.Logger.Fatal(e.Start(":1323"))
}

