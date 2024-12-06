// assetManager.go
package assetManager

import (
    "fmt"
    "log"
    "sync"

    "github.com/a-h/templ"
    "github.com/labstack/echo/v4"
    "knovault/internal/types"
)

var _ types.AssetManager = (*AssetManager)(nil)  // Compile-time interface check

type AssetManager struct {
    plugins      map[string]types.Plugin
    themes       map[string]types.Theme
    currentTheme types.Theme
    pluginInfo   map[string]types.PluginMetadata
    themeInfo    map[string]types.ThemeMetadata
    mutex        sync.RWMutex
}

func NewAssetManager() *AssetManager {
    return &AssetManager{
        plugins:    make(map[string]types.Plugin),
        themes:     make(map[string]types.Theme),
        pluginInfo: make(map[string]types.PluginMetadata),
        themeInfo:  make(map[string]types.ThemeMetadata),
    }
}

func (am *AssetManager) LoadAssets() error {
    // Load plugins
    if err := am.loadPlugins(); err != nil {
        log.Printf("Warning: Error loading plugins: %v", err)
    }

    // Load themes
    if err := am.loadThemes(); err != nil {
        log.Printf("Warning: Error loading themes: %v", err)
    }

    // Verify we have at least one theme
    if len(am.themes) == 0 {
        return fmt.Errorf("no themes were loaded")
    }

    return nil
}

// Theme management methods
func (am *AssetManager) GetCurrentTheme() types.Theme {
    am.mutex.RLock()
    defer am.mutex.RUnlock()
    return am.currentTheme
}

func (am *AssetManager) SetCurrentTheme(name string) error {
    am.mutex.Lock()
    defer am.mutex.Unlock()

    theme, ok := am.themes[name]
    if !ok {
        return fmt.Errorf("theme %s not found", name)
    }

    am.currentTheme = theme
    log.Printf("Current theme set to: %s", name)
    return nil
}

func (am *AssetManager) GetCurrentThemeName() string {
    am.mutex.RLock()
    defer am.mutex.RUnlock()

    for name, theme := range am.themes {
        if theme == am.currentTheme {
            return name
        }
    }
    return ""
}

func (am *AssetManager) GetAvailableThemes() []string {
    am.mutex.RLock()
    defer am.mutex.RUnlock()

    var names []string
    for name := range am.themes {
        names = append(names, name)
    }
    return names
}

// Plugin management methods
func (am *AssetManager) GetPlugin(name string) (types.Plugin, bool) {
    am.mutex.RLock()
    defer am.mutex.RUnlock()
    p, ok := am.plugins[name]
    return p, ok
}

func (am *AssetManager) ListPlugins() []types.PluginInfo {
    am.mutex.RLock()
    defer am.mutex.RUnlock()

    var plugins []types.PluginInfo
    for name, p := range am.plugins {
        info := types.PluginInfo{
            Name:        name,
            Description: p.Description(),
        }
        if metadata, ok := am.pluginInfo[name]; ok {
            info.Tags = metadata.Tags
        }
        plugins = append(plugins, info)
    }
    return plugins
}

// Template extension methods
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

// Route management methods
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