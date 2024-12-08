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

    log.Println("Loading builtin themes...")
    builtinThemes, err := loadBuiltinThemes()
    if err != nil {
        log.Printf("⚠️  Warning: Error loading builtin themes: %v", err)
    }
    for name, theme := range builtinThemes {
        log.Printf("✓ Loaded builtin theme: %s", name)
        tm.themes[name] = theme
        tm.themeInfo[name] = types.ThemeInfo{
            Name: name,
            Tags: []string{"builtin"},
        }
    }

    log.Println("Loading external themes...")
    externalThemes, err := loadExternalThemes()
    if err != nil {
        log.Printf("⚠️  Warning: Error loading external themes: %v", err)
    }
    for name, theme := range externalThemes {
        log.Printf("✓ Loaded external theme: %s", name)
        tm.themes[name] = theme
        tm.themeInfo[name] = types.ThemeInfo{
            Name: name,
            Tags: []string{"external"},
        }
    }

    if len(tm.themes) == 0 {
        return fmt.Errorf("❌ No themes were loaded")
    }

    // Set default theme
    if theme, ok := tm.themes["defaultTheme"]; ok {
        tm.currentTheme = theme
        log.Println("✓ Set default theme: defaultTheme")
    } else {
        // Set first available theme as default
        for name, theme := range tm.themes {
            tm.currentTheme = theme
            log.Printf("⚠️  Default theme not found, using: %s", name)
            break
        }
    }

    log.Printf("✓ Theme Manager initialized with %d themes", len(tm.themes))
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



