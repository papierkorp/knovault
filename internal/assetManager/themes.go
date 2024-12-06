package assetManager

import (
    "fmt"
    "log"
    "path/filepath"
    "plugin"

    "knovault/internal/types"
)

func (am *AssetManager) loadThemes() error {
    config, err := loadConfig[types.ThemeConfig]("internal/assetManager/themes_list.json")
    if err != nil {
        return err
    }

    am.mutex.Lock()
    defer am.mutex.Unlock()

    themesDir := "./internal/assetManager/themes"
    files, err := findAssetFiles(themesDir)
    if err != nil {
        return fmt.Errorf("failed to scan themes directory: %v", err)
    }

    // Create a map of configured themes
    configuredThemes := make(map[string]types.ThemeMetadata)
    for _, t := range config.Themes {
        if t.Enabled {
            configuredThemes[t.Name] = t
            am.themeInfo[t.Name] = t
        }
    }

    // Load themes from found files
    for _, file := range files {
        // For .so files, get the theme name from the parent directory
        // For plugin.go files, get the theme name from the grandparent directory
        var themeName string
        if filepath.Ext(file) == ".so" {
            themeName = filepath.Base(filepath.Dir(file))
        } else {
            themeName = filepath.Base(filepath.Dir(filepath.Dir(filepath.Dir(file))))
        }

        // Check if this theme is configured and enabled
        if metadata, ok := configuredThemes[themeName]; ok {
            if err := am.loadThemeFromPath(themeName, file); err != nil {
                log.Printf("Warning: Could not load theme %s: %v", themeName, err)
                continue
            }
            am.themeInfo[themeName] = metadata
        }
    }

    // Verify that themes were actually loaded
    if len(am.themes) == 0 {
        return fmt.Errorf("no themes were loaded")
    }

    return nil
}

func (am *AssetManager) loadThemeFromPath(name, path string) error {
    if filepath.Ext(path) == ".so" {
        // For .so files, load using plugin.Open
        plug, err := plugin.Open(path)
        if err != nil {
            return fmt.Errorf("could not open theme: %v", err)
        }

        symTheme, err := plug.Lookup("Theme")
        if err != nil {
            return fmt.Errorf("could not find Theme symbol: %v", err)
        }

        // Try both direct interface and pointer conversion
        if themeInstance, ok := symTheme.(types.Theme); ok {
            am.themes[name] = themeInstance
        } else if themePtr, ok := symTheme.(*types.Theme); ok {
            am.themes[name] = *themePtr
        } else {
            return fmt.Errorf("invalid theme type")
        }
    } else if filepath.Base(path) == "plugin.go" {
        // For plugin.go files, assume Theme is a package-level variable
        // and will be loaded when the plugin is compiled to .so
        // We'll need to compile it first using go build -buildmode=plugin
        // pluginDir := filepath.Dir(filepath.Dir(path))
        // outPath := filepath.Join(pluginDir, name+".so")

        // Remove any existing .so file to avoid loading stale plugins
        // This is commented out for now as it might be better to handle this in the Makefile
        // os.Remove(outPath)

        log.Printf("Theme plugin source found at %s. Please compile using: make compile-theme THEME=%s", path, name)
        return nil
    }

    log.Printf("Loaded theme: %s", name)
    return nil
}