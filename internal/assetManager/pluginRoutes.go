// pluginRoutes.go

package assetManager

import (
    "github.com/a-h/templ"
    "github.com/labstack/echo/v4"
    "knovault/internal/types"
)

func (am *AssetManager) ApplyPluginRoutes(e *echo.Echo) {
    am.mutex.RLock()
    defer am.mutex.RUnlock()

    for _, p := range am.plugins {
        if routePlugin, ok := p.(types.PluginWithRoute); ok {
            route := routePlugin.Route()
            e.Add(route.Method, route.Path, route.Handler)
        }
    }
}

func (am *AssetManager) GetPluginTemplateExtensions(templateName string) []templ.Component {
    am.mutex.RLock()
    defer am.mutex.RUnlock()

    var extensions []templ.Component
    for _, p := range am.plugins {
        if extendablePlugin, ok := p.(types.PluginWithTemplateExtensions); ok {
            if extension, err := extendablePlugin.ExtendTemplate(templateName); err == nil {
                extensions = append(extensions, extension)
            }
        }
    }
    return extensions
}