// internal/server/server.go

package server

import (
    "log"
    "knovault/internal/assetManager"
    "knovault/internal/globals"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func Start() {
    e := echo.New()

    manager := assetManager.NewAssetManager()
    globals.SetAssetManager(manager)

    if err := manager.LoadAssets(); err != nil {
        log.Fatalf("Failed to load assets: %v", err)
    }

    err := manager.SetCurrentTheme("defaultTheme")
    if err != nil {
        log.Fatalf("Failed to set default theme: %v", err)
    }

    e.Static("/static", "static")
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    setupRoutes(e)

    e.Logger.Fatal(e.Start(":1323"))
}