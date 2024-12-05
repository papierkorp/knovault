package assetManager

import (
    "fmt"
    "log"
    "path/filepath"
    "plugin"
    "os"

    "knovault/internal/types"
)

func (am *AssetManager) loadConfiguredThemes() error {
    config, err := loadConfig[types.ThemeConfig]("internal/assetManager/themes_list.json")
    if err != nil {
        return err
    }

    am.mutex.Lock()
    defer am.mutex.Unlock()

    for _, t := range config.Themes {
        if !t.Enabled {
            continue
        }

        // Store metadata regardless of loading success
        am.themeInfo[t.Name] = t

        if hasTag(t.Tags, "built-in") {
            themePath := filepath.Join(t.Path)
            soPath := filepath.Join(filepath.Dir(themePath), "plugin.so")
            if err := compileModule(themePath, soPath); err != nil {
                log.Printf("Warning: Could not compile theme %s: %v", t.Name, err)
                continue
            }

            if err := am.loadThemeFromPath(t.Name, soPath); err != nil {
                log.Printf("Warning: Could not load theme %s: %v", t.Name, err)
                continue
            }
        }
    }

    return nil
}

func (am *AssetManager) loadThemeFromPath(name, path string) error {
    themePlugin, err := plugin.Open(path)
    if err != nil {
        return fmt.Errorf("could not open theme: %v", err)
    }

    symTheme, err := themePlugin.Lookup("Theme")
    if err != nil {
        return fmt.Errorf("could not find Theme symbol: %v", err)
    }

    themeInstance, ok := symTheme.(types.Theme)
    if !ok {
        return fmt.Errorf("invalid theme type")
    }

    am.themes[name] = themeInstance
    log.Printf("Loaded theme: %s", name)

    // Set as current theme if it's tagged as default and no current theme
    if am.currentTheme == nil {
        metadata, exists := am.themeInfo[name]
        if exists && hasTag(metadata.Tags, "default") {
            am.currentTheme = themeInstance
            log.Printf("Set default theme: %s", name)
        }
    }

    return nil
}

func (am *AssetManager) loadCompiledThemes() error {
    themesDir := "./internal/assetManager/themes"
    entries, err := os.ReadDir(themesDir)
    if err != nil {
        return fmt.Errorf("failed to read themes directory: %v", err)
    }

    for _, entry := range entries {
        if !entry.IsDir() && filepath.Ext(entry.Name()) == ".so" {
            name := filepath.Base(entry.Name())
            name = name[:len(name)-3] // Remove .so extension
            path := filepath.Join(themesDir, entry.Name())

            if err := am.loadThemeFromPath(name, path); err != nil {
                log.Printf("Failed to load theme %s: %v", name, err)
            }
        }
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

func (am *AssetManager) GetThemesByTag(tag string) []string {
    am.mutex.RLock()
    defer am.mutex.RUnlock()

    var themes []string
    for name, metadata := range am.themeInfo {
        if hasTag(metadata.Tags, tag) {
            themes = append(themes, name)
        }
    }
    return themes
}
