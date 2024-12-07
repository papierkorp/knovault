// internal/themeManager/themeManager.go
package themeManager

import (
    "fmt"
    "log"
    "sync"
    "knovault/internal/types"
)

type ThemeManager struct {
    themes      map[string]types.Theme
    themeInfo   map[string]types.ThemeInfo
    currentTheme types.Theme
    mutex       sync.RWMutex
}

func NewThemeManager() *ThemeManager {
    return &ThemeManager{
        themes:    make(map[string]types.Theme),
        themeInfo: make(map[string]types.ThemeInfo),
    }
}

func (tm *ThemeManager) Initialize() error {
    tm.mutex.Lock()
    defer tm.mutex.Unlock()

    // Load builtin themes
    builtinThemes, err := loadBuiltinThemes()
    if err != nil {
        log.Printf("Warning: Error loading builtin themes: %v", err)
    }
    for name, theme := range builtinThemes {
        tm.themes[name] = theme
        tm.themeInfo[name] = types.ThemeInfo{
            Name: name,
            Tags: []string{"builtin"},
        }
    }

    // Load external themes
    externalThemes, err := loadExternalThemes()
    if err != nil {
        log.Printf("Warning: Error loading external themes: %v", err)
    }
    for name, theme := range externalThemes {
        tm.themes[name] = theme
        tm.themeInfo[name] = types.ThemeInfo{
            Name: name,
            Tags: []string{"external"},
        }
    }

    if len(tm.themes) == 0 {
        return fmt.Errorf("no themes were loaded")
    }

    // Set default theme
    if theme, ok := tm.themes["defaultTheme"]; ok {
        tm.currentTheme = theme
    } else {
        // Set first available theme as default
        for _, theme := range tm.themes {
            tm.currentTheme = theme
            break
        }
    }

    return nil
}

func (tm *ThemeManager) GetCurrentTheme() types.Theme {
    tm.mutex.RLock()
    defer tm.mutex.RUnlock()
    return tm.currentTheme
}

func (tm *ThemeManager) SetCurrentTheme(name string) error {
    tm.mutex.Lock()
    defer tm.mutex.Unlock()

    theme, ok := tm.themes[name]
    if !ok {
        return fmt.Errorf("theme %s not found", name)
    }

    tm.currentTheme = theme
    return nil
}

func (tm *ThemeManager) GetCurrentThemeName() string {
    tm.mutex.RLock()
    defer tm.mutex.RUnlock()

    for name, theme := range tm.themes {
        if theme == tm.currentTheme {
            return name
        }
    }
    return ""
}

func (tm *ThemeManager) GetAvailableThemes() []string {
    tm.mutex.RLock()
    defer tm.mutex.RUnlock()

    var names []string
    for name := range tm.themes {
        names = append(names, name)
    }
    return names
}



