// internal/server/routes.go
package server

import (
    "github.com/labstack/echo/v4"
    "knovault/internal/globals"
    "knovault/internal/types"
)

func setupRoutes(e *echo.Echo) {
    e.GET("/", handleHome)
    e.GET("/home", handleHome)
    e.GET("/help", handleHelp)
    e.GET("/settings", handleSettings)
    e.GET("/search", handleSearch)
    e.GET("/docs", handleDocsRoot)
    e.GET("/docs/:title", handleDocs)
    e.GET("/playground", handlePlayground)
    e.GET("/plugins", handlePlugins)

    // Apply plugin routes
    applyPluginRoutes(e)
}

func applyPluginRoutes(e *echo.Echo) {
    pm := globals.GetPluginManager()
    for _, info := range pm.ListPlugins() {
        if plugin, ok := pm.GetPlugin(info.Name); ok {
            if routePlugin, ok := plugin.(types.PluginWithRoute); ok {
                route := routePlugin.Route()
                e.Add(route.Method, route.Path, route.Handler)
            }
        }
    }
}

