// internal/globals/globals.go
package globals

import (
    "knovault/internal/types"
)

var (
    pluginManager types.PluginManager
    themeManager types.ThemeManager
)

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