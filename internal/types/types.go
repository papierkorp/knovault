// internal/types/types.go
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

// ThemeInfo structure for theme metadata
type ThemeInfo struct {
    Name        string   `json:"name"`
    Description string   `json:"description"`
    Tags        []string `json:"tags"`
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

// Manager interfaces
type PluginManager interface {
    GetPlugin(name string) (Plugin, bool)
    ListPlugins() []PluginInfo
    Initialize() error
}

type ThemeManager interface {
    GetCurrentTheme() Theme
    SetCurrentTheme(name string) error
    GetCurrentThemeName() string
    GetAvailableThemes() []string
    Initialize() error
}

// Extension interfaces for additional functionality
type PluginManagerWithExtensions interface {
    PluginManager
    GetPluginTemplateExtensions(templateName string) []templ.Component
}

type ThemeManagerWithRoutes interface {
    ThemeManager
    ApplyThemeRoutes(e *echo.Echo)
}

// Configuration types
type ThemeMetadata struct {
    Name    string   `json:"name"`
    Path    string   `json:"path"`
    Enabled bool     `json:"enabled"`
    Tags    []string `json:"tags"`
}

type PluginMetadata struct {
    Name    string   `json:"name"`
    Path    string   `json:"path"`
    Enabled bool     `json:"enabled"`
    Tags    []string `json:"tags"`
}