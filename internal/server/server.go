package server

import (
	"gowiki/internal/plugins"
	"gowiki/internal/plugins/core"
	"gowiki/internal/themes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	e := echo.New()

	pluginManager := plugins.NewManager()
    themes.SetPluginManager(pluginManager)
    pluginManager.RegisterPlugin(&core.ThemeSwitcherPlugin{})
    pluginManager.RegisterPlugin(&core.DarkModePlugin{})

	e.Static("/", "static")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	setupRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
