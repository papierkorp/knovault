package server

import (
	"pewito/internal/themes"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pewito/internal/plugins"
	_ "pewito/internal/plugins/core"
)

func Start() {
    e := echo.New()

    if err := themes.LoadThemes(); err != nil {
        log.Fatalf("Failed to load themes: %v", err)
    }
    
    err := themes.SetCurrentTheme("defaultTheme")
    if err != nil {
        log.Fatalf("Failed to set default theme: %v", err)
    }

    if err := plugins.LoadCommonPlugins(); err != nil {
        log.Printf("Error loading common plugins: %v", err)
    }


    e.Static("/static", "static")

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    setupRoutes(e)

    e.Logger.Fatal(e.Start(":1323"))
}

