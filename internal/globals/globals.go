// internal/globals/globals.go
package globals

import (
    "knovault/internal/types"
    "sync"
)

var (
    pluginManager types.PluginManager
    themeManager types.ThemeManager

    // Add plugin registry
    pluginRegistry = struct {
        sync.RWMutex
        constructors map[string]func() types.Plugin
    }{
        constructors: make(map[string]func() types.Plugin),
    }
)

// Plugin registry functions
func RegisterPlugin(name string, constructor func() types.Plugin) {
    pluginRegistry.Lock()
    defer pluginRegistry.Unlock()
    pluginRegistry.constructors[name] = constructor
}

func GetPluginConstructors() map[string]func() types.Plugin {
    pluginRegistry.RLock()
    defer pluginRegistry.RUnlock()
    // Return a copy to prevent map modifications
    constructors := make(map[string]func() types.Plugin)
    for k, v := range pluginRegistry.constructors {
        constructors[k] = v
    }
    return constructors
}

// Existing functions
func SetPluginManager(pm types.PluginManager) {
    pluginManager = pm
}

func GetPluginManager() types.PluginManager {
    return pluginManager
}

func SetThemeManager(tm types.ThemeManager) {
    themeManager = tm
}

func GetThemeManager() types.ThemeManager {
    return themeManager
}