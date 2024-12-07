// internal/themeManager/themeManagerExternalHelper.go
package themeManager

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "plugin"
    "os/exec"
    "knovault/internal/types"
)

func loadExternalThemes() (map[string]types.Theme, error) {
    themes := make(map[string]types.Theme)
    externalDir := "./internal/themeManager/external"

    // First, compile any main.go files to .so
    err := compileExternalThemes(externalDir)
    if err != nil {
        log.Printf("Warning: Error compiling external themes: %v", err)
    }

    // Then load all .so files
    err = filepath.Walk(externalDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() && filepath.Ext(path) == ".so" {
            themeName := filepath.Base(filepath.Dir(path))
            theme, err := loadThemeFromFile(path)
            if err != nil {
                log.Printf("Warning: Could not load theme %s: %v", themeName, err)
                return nil
            }
            themes[themeName] = theme
        }
        return nil
    })

    return themes, err
}

func compileExternalThemes(dir string) error {
    return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() && filepath.Base(path) == "main.go" {
            themeDir := filepath.Dir(path)
            themeName := filepath.Base(filepath.Dir(themeDir))
            outPath := filepath.Join(themeDir, themeName+".so")

            // Remove existing .so file
            os.Remove(outPath)

            cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", outPath, path)
            if err := cmd.Run(); err != nil {
                log.Printf("Warning: Could not compile theme %s: %v", themeName, err)
            }
        }
        return nil
    })
}

func loadThemeFromFile(path string) (types.Theme, error) {
    plug, err := plugin.Open(path)
    if err != nil {
        return nil, fmt.Errorf("could not open theme: %v", err)
    }

    symTheme, err := plug.Lookup("Theme")
    if err != nil {
        return nil, fmt.Errorf("could not find Theme symbol: %v", err)
    }

    var t types.Theme
    switch v := symTheme.(type) {
    case types.Theme:
        t = v
    case *types.Theme:
        t = *v
    default:
        return nil, fmt.Errorf("invalid theme type")
    }

    return t, nil
}