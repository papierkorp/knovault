package themes

import (
    "fmt"
    "pewitima/internal/types"
    "log"
    "sync"
)

var (
    themes       = make(map[string]types.Theme)
    currentTheme types.Theme
    themeMutex   sync.RWMutex
)

func RegisterTheme(name string, theme types.Theme) {
    themeMutex.Lock()
    defer themeMutex.Unlock()
    themes[name] = theme
    log.Printf("Theme registered: %s", name)
    if currentTheme == nil {
        currentTheme = theme
        log.Printf("Set default theme: %s", name)
    }
}

func SetCurrentTheme(name string) error {
    themeMutex.Lock()
    defer themeMutex.Unlock()
    theme, ok := themes[name]
    if !ok {
        return fmt.Errorf("theme %s not found", name)
    }
    currentTheme = theme
    log.Printf("Current theme set to: %s", name)
    return nil
}

func GetCurrentTheme() types.Theme {
    themeMutex.RLock()
    defer themeMutex.RUnlock()
    return currentTheme
}

func GetCurrentThemeName() string {
    themeMutex.RLock()
    defer themeMutex.RUnlock()
    for name, theme := range themes {
        if theme == currentTheme {
            return name
        }
    }
    return ""
}

func GetAvailableThemes() []string {
    themeMutex.RLock()
    defer themeMutex.RUnlock()
    var themeNames []string
    for name := range themes {
        themeNames = append(themeNames, name)
    }
    return themeNames
}