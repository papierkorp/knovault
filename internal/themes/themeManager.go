package themes

import (
    "fmt"
    "io/ioutil"
    "log"
    "path/filepath"
    "pewito/internal/types"
    "plugin"
    "sync"
)

var (
    themes       = make(map[string]types.Theme)
    currentTheme types.Theme
    themeMutex   sync.RWMutex
)

func LoadThemes() error {
    themesDir := "./internal/themes"
    entries, err := ioutil.ReadDir(themesDir)
    if err != nil {
        return fmt.Errorf("failed to read themes directory: %v", err)
    }

    for _, entry := range entries {
        if entry.IsDir() {
            themeName := entry.Name()
            themePluginPath := filepath.Join(themesDir, themeName, themeName+".so")

            p, err := plugin.Open(themePluginPath)
            if err != nil {
                log.Printf("Failed to load theme %s: %v", themeName, err)
                continue
            }

            symTheme, err := p.Lookup("Theme")
            if err != nil {
                log.Printf("Failed to find Theme symbol in %s: %v", themeName, err)
                continue
            }

            theme, ok := symTheme.(types.Theme)
            if !ok {
                log.Printf("Unexpected type for Theme in %s", themeName)
                continue
            }

            RegisterTheme(themeName, theme)
        }
    }

    if len(themes) == 0 {
        return fmt.Errorf("no themes were loaded")
    }

    return nil
}

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