package server

import (
	"pewitima/internal/themes"
	_ "pewitima/internal/themes/defaultTheme"
	_ "pewitima/internal/themes/secondTheme"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pewitima/internal/plugins"
	_ "pewitima/internal/plugins/core"
)

func Start() {
    e := echo.New()

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