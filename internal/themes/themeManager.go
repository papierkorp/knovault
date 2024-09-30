package themes

import (
	"fmt"
	"gowiki/internal/types"
	"gowiki/internal/plugins"
)


var (
	currentTheme types.Theme
	themes       = make(map[string]types.Theme)
	pluginManager *plugins.Manager
)

func RegisterTheme(name string, theme types.Theme) {
	fmt.Printf("Registering theme: %s\n", name)
	themes[name] = theme
	fmt.Printf("Registered themes: %v\n", themes)
}

func SetCurrentTheme(name string) error {
    fmt.Printf("Setting current theme to: %s\n", name)
    theme, ok := themes[name]
    if !ok {
        availableThemes := GetAvailableThemes()
        fmt.Printf("Theme %s not found. Available themes: %v\n", name, availableThemes)
        return fmt.Errorf("theme %s not found. Available themes: %v", name, availableThemes)
    }
    currentTheme = theme
    fmt.Printf("Current theme set to: %s\n", name)
    return nil
}

func GetCurrentTheme() types.Theme {
    fmt.Printf("Getting current theme. Available themes: %v\n", themes)
    fmt.Printf("Current theme: %T\n", currentTheme)
	return currentTheme
}

func GetAvailableThemes() []string {
    var availableThemes []string
    for name := range themes {
        availableThemes = append(availableThemes, name)
    }
    return availableThemes
}

func SetPluginManager(manager *plugins.Manager) {
    pluginManager = manager
}

func GetPluginManager() *plugins.Manager {
    return pluginManager
}