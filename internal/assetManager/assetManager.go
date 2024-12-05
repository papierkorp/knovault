package assetManager

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "plugin"
    "sync"

    "github.com/a-h/templ"
    "github.com/labstack/echo/v4"

    "knovault/internal/types"
    // Plugin imports
    cssPlugin "knovault/internal/assetManager/plugins/CustomCSS/plugin"
    filePlugin "knovault/internal/assetManager/plugins/FileManager/plugin"
    helloPlugin "knovault/internal/assetManager/plugins/HelloWorld/plugin"
    mdPlugin "knovault/internal/assetManager/plugins/MarkdownParser/plugin"
    themePlugin "knovault/internal/assetManager/plugins/ThemeChanger/plugin"
    // Theme imports
    defaultTheme "knovault/internal/assetManager/themes/defaultTheme/plugin"
    dark "knovault/internal/assetManager/themes/dark/plugin"
)

type AssetManager struct {
    plugins      map[string]types.Plugin
    themes       map[string]types.Theme
    currentTheme types.Theme
    mutex        sync.RWMutex
}

// NewAssetManager creates a new instance of AssetManager
func NewAssetManager() *AssetManager {
    return &AssetManager{
        plugins: make(map[string]types.Plugin),
        themes:  make(map[string]types.Theme),
    }
}

// LoadAssets loads all plugins and themes
func (am *AssetManager) LoadAssets() error {
    // Load plugins
    if err := am.loadBuiltInPlugins(); err != nil {
        log.Printf("Warning: Error loading built-in plugins: %v", err)
    }

    if err := am.loadCompiledPlugins(); err != nil {
        log.Printf("Warning: Error loading compiled plugins: %v", err)
    }

    // Load themes
    if err := am.loadBuiltInThemes(); err != nil {
        log.Printf("Warning: Error loading built-in themes: %v", err)
    }

    if err := am.loadCompiledThemes(); err != nil {
        log.Printf("Warning: Error loading compiled themes: %v", err)
    }

    // Verify we have at least one theme
    if len(am.themes) == 0 {
        return fmt.Errorf("no themes were loaded")
    }

    return nil
}

// Plugin management
func (am *AssetManager) loadBuiltInPlugins() error {
    builtInPlugins := map[string]types.Plugin{
        "CustomCSS":      cssPlugin.Plugin,
        "FileManager":    filePlugin.Plugin,
        "HelloWorld":     helloPlugin.Plugin,
        "MarkdownParser": mdPlugin.Plugin,
        "ThemeChanger":   themePlugin.Plugin,
    }

    am.mutex.Lock()
    defer am.mutex.Unlock()

    for name, p := range builtInPlugins {
        am.plugins[name] = p
        log.Printf("Loaded built-in plugin: %s", name)
    }

    return nil
}

func (am *AssetManager) loadCompiledPlugins() error {
    pluginsDir := "./internal/assetManager/plugins"
    return am.loadCompiledAssets(pluginsDir, am.loadPluginFromSO)
}

// Theme management
func (am *AssetManager) loadBuiltInThemes() error {
    builtInThemes := map[string]types.Theme{
        "defaultTheme": &defaultTheme.Theme, // Changed to pointer
        "dark": &dark.Theme,
    }

    am.mutex.Lock()
    defer am.mutex.Unlock()

    for name, theme := range builtInThemes {
        am.themes[name] = theme
        log.Printf("Loaded built-in theme: %s", name)

        // Set first theme as current if none set
        if am.currentTheme == nil {
            am.currentTheme = theme
            log.Printf("Set default theme: %s", name)
        }
    }

    return nil
}

func (am *AssetManager) loadCompiledThemes() error {
    themesDir := "./internal/assetManager/themes"
    return am.loadCompiledAssets(themesDir, am.loadThemeFromSO)
}

// Shared loading functionality
func (am *AssetManager) loadCompiledAssets(dir string, loader func(string) error) error {
    entries, err := os.ReadDir(dir)
    if err != nil {
        return fmt.Errorf("failed to read directory %s: %v", dir, err)
    }

    for _, entry := range entries {
        if !entry.IsDir() && filepath.Ext(entry.Name()) == ".so" {
            path := filepath.Join(dir, entry.Name())
            if err := loader(path); err != nil {
                log.Printf("Failed to load from %s: %v", path, err)
            }
        }
    }

    return nil
}

func (am *AssetManager) loadPluginFromSO(path string) error {
    p, err := plugin.Open(path)
    if err != nil {
        return err
    }

    symPlugin, err := p.Lookup("Plugin")
    if err != nil {
        return err
    }

    pl, ok := symPlugin.(types.Plugin)
    if !ok {
        return fmt.Errorf("invalid plugin type")
    }

    am.mutex.Lock()
    defer am.mutex.Unlock()
    am.plugins[pl.Name()] = pl
    return nil
}

func (am *AssetManager) loadThemeFromSO(path string) error {
    p, err := plugin.Open(path)
    if err != nil {
        return err
    }

    symTheme, err := p.Lookup("Theme")
    if err != nil {
        return err
    }

    theme, ok := symTheme.(types.Theme)
    if !ok {
        return fmt.Errorf("invalid theme type")
    }

    am.mutex.Lock()
    defer am.mutex.Unlock()

    themeName := filepath.Base(path)
    themeName = themeName[:len(themeName)-3] // Remove .so extension
    am.themes[themeName] = theme

    if am.currentTheme == nil {
        am.currentTheme = theme
    }

    return nil
}

// Public accessors
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
        plugins = append(plugins, types.PluginInfo{
            Name:        name,
            Description: p.Description(),
        })
    }
    return plugins
}

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