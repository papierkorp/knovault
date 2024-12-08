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
    "strings"
)

func loadExternalThemes() (map[string]types.Theme, error) {
    themes := make(map[string]types.Theme)
    externalDir := "./internal/themeManager/external"

    absExternalDir, err := filepath.Abs(externalDir)
    if err != nil {
        return nil, fmt.Errorf("failed to get absolute path: %v", err)
    }
    log.Printf("Loading themes from: %s", absExternalDir)

    log.Println("Compiling external themes...")
    err = compileExternalThemes(absExternalDir)
    if err != nil {
        log.Printf("⚠️  Warning: Error compiling external themes: %v", err)
    }

    files, err := os.ReadDir(absExternalDir)
    if err != nil {
        log.Printf("Error reading external directory: %v", err)
        return themes, err
    }

    for _, file := range files {
        if !file.IsDir() {
            continue
        }

        themeName := file.Name()
        themeDir := filepath.Join(absExternalDir, themeName)
        mainPath := filepath.Join(themeDir, "main.go")
        soPath := filepath.Join(themeDir, strings.ToLower(themeName)+".so")

        log.Printf("Processing theme directory: %s", themeDir)
        log.Printf("Looking for main.go at: %s", mainPath)
        log.Printf("Looking for .so file at: %s", soPath)

        if mainFileInfo, err := os.Stat(mainPath); err == nil && !mainFileInfo.IsDir() {
            log.Printf("Found main.go for theme: %s", themeName)

            log.Printf("Loading theme: %s", themeName)
            theme, err := loadThemeFromFile(soPath)
            if err != nil {
                log.Printf("⚠️  Could not load theme %s: %v", themeName, err)
                continue
            }

            themes[themeName] = theme
            log.Printf("✓ Successfully loaded theme: %s", themeName)
            log.Printf("✓ Loaded external theme: %s", themeName)
        } else {
            log.Printf("Skipping directory %s - not a valid theme (no main.go found)", themeName)
        }
    }

    log.Printf("Found %d external themes", len(themes))
    for name := range themes {
        log.Printf("✓ Available external theme: %s", name)
    }

    return themes, nil
}

func compileExternalThemes(absDir string) error {
    files, err := os.ReadDir(absDir)
    if err != nil {
        return err
    }

    for _, file := range files {
        if !file.IsDir() {
            continue
        }

        themeName := file.Name()
        mainPath := filepath.Join(absDir, themeName, "main.go")
        outPath := filepath.Join(absDir, themeName, strings.ToLower(themeName)+".so")

        log.Printf("Checking theme directory: %s", themeName)
        log.Printf("Looking for main.go at: %s", mainPath)

        if mainFileInfo, err := os.Stat(mainPath); err == nil && !mainFileInfo.IsDir() {
            log.Printf("Compiling theme: %s", themeName)

            // Remove existing .so file
            os.Remove(outPath)

            cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", outPath, mainPath)
            cmd.Dir = filepath.Dir(mainPath)

            output, err := cmd.CombinedOutput()
            if err != nil {
                log.Printf("⚠️  Could not compile theme %s: %v\nOutput: %s", themeName, err, string(output))
                continue
            }
            log.Printf("✓ Successfully compiled theme: %s", themeName)
        } else {
            log.Printf("No main.go found for theme %s", themeName)
        }
    }

    return nil
}

func loadThemeFromFile(path string) (types.Theme, error) {
    // Check if file exists before attempting to open
    if _, err := os.Stat(path); err != nil {
        return nil, fmt.Errorf("theme file does not exist: %v", err)
    }

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
        return nil, fmt.Errorf("invalid theme type: %T", symTheme)
    }

    return t, nil
}