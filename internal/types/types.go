package types

import (
    "github.com/a-h/templ"
    "github.com/labstack/echo/v4"
)

// Theme interface defines required methods for themes
type Theme interface {
    Home() (templ.Component, error)
    Help() (templ.Component, error)
    Settings() (templ.Component, error)
    Search() (templ.Component, error)
    DocsRoot() (templ.Component, error)
    Docs(content string) (templ.Component, error)
    Playground() (templ.Component, error)
    Plugins() (templ.Component, error)
}

// Plugin interface defines required methods for plugins
type Plugin interface {
    Name() string
    Description() string
    Help() string
    TemplResponse() (templ.Component, error)
    JsonResponse() ([]byte, error)
    Execute(params map[string]string) (interface{}, error)
}

// PluginWithRoute interface for plugins that provide routes
type PluginWithRoute interface {
    Plugin
    Route() PluginRoute
}

// PluginWithTemplateExtensions interface for plugins that extend templates
type PluginWithTemplateExtensions interface {
    Plugin
    ExtendTemplate(templateName string) (templ.Component, error)
}

// PluginInfo structure for plugin metadata
type PluginInfo struct {
    Name        string   `json:"name"`
    Description string   `json:"description"`
    Tags        []string `json:"tags"`
}

// PluginRoute structure for plugin routes
type PluginRoute struct {
    Method  string
    Path    string
    Handler echo.HandlerFunc
}

// AssetManager interface defines methods for managing assets
type AssetManager interface {
    GetPlugin(name string) (Plugin, bool)
    ListPlugins() []PluginInfo
    GetCurrentTheme() Theme
    SetCurrentTheme(name string) error
    GetCurrentThemeName() string
    GetAvailableThemes() []string
    GetPluginTemplateExtensions(templateName string) []templ.Component
    LoadAssets() error
    ApplyPluginRoutes(e *echo.Echo)
}