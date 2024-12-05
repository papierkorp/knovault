package assetManager

import (
    "fmt"
    "log"
    "sync"

    "knovault/internal/types"
)

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
    if err := am.loadConfiguredPlugins(); err != nil {
        log.Printf("Warning: Error loading configured plugins: %v", err)
    }

    if err := am.loadCompiledPlugins(); err != nil {
        log.Printf("Warning: Error loading compiled plugins: %v", err)
    }

    // Load themes
    if err := am.loadConfiguredThemes(); err != nil {
        log.Printf("Warning: Error loading configured themes: %v", err)
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

