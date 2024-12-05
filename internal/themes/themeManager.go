package themes

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "plugin"
    "sync"
    "knovault/internal/types"
    "strings"
    "reflect"
)

var (
    themes       = make(map[string]types.Theme)
    currentTheme types.Theme
    themeMutex   sync.RWMutex
    themeDirectory = "./internal/themes"
)

// compileCoreThemes compiles all core themes into .so files
func compileCoreThemes() error {
    coreThemeDir := filepath.Join(themeDirectory, "core")
    log.Printf("Compiling core themes from directory: %s", coreThemeDir)

    entries, err := os.ReadDir(coreThemeDir)
    if err != nil {
        return fmt.Errorf("failed to read core theme directory: %v", err)
    }

    for _, entry := range entries {
        if entry.IsDir() {
            themeName := entry.Name()
            themeDir := filepath.Join(coreThemeDir, themeName)
            outputFile := filepath.Join(themeDir, themeName+".so")

            log.Printf("Attempting to compile theme %s", themeName)
            log.Printf("Theme directory: %s", themeDir)
            log.Printf("Output file: %s", outputFile)

            // Change to theme directory and build
            currentDir, err := os.Getwd()
            if err != nil {
                log.Printf("Failed to get current directory for theme %s: %v", themeName, err)
                continue
            }

            if err := os.Chdir(themeDir); err != nil {
                log.Printf("Failed to change to theme directory for %s: %v", themeName, err)
                continue
            }

            cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", outputFile, "main.go")
            if output, err := cmd.CombinedOutput(); err != nil {
                log.Printf("Failed to compile core theme %s: %v\nOutput: %s", themeName, err, string(output))
                os.Chdir(currentDir)
                continue
            }

            log.Printf("Successfully compiled theme %s", themeName)

            if err := os.Chdir(currentDir); err != nil {
                log.Printf("Failed to change back to original directory after building %s: %v", themeName, err)
            }
        }
    }
    return nil
}

// LoadThemes loads both core and common themes
func LoadThemes() error {
    // First compile core themes
    if err := compileCoreThemes(); err != nil {
        return err
    }

    // Load core themes
    if err := loadThemesFromDir(filepath.Join(themeDirectory, "core"), true); err != nil {
        return err
    }

    // Load common themes
    if err := loadThemesFromDir(filepath.Join(themeDirectory, "common"), false); err != nil {
        return err
    }

    if len(themes) == 0 {
        return fmt.Errorf("no themes were loaded")
    }

    return nil
}

// loadThemesFromDir loads all theme .so files from the specified directory
func loadThemesFromDir(dir string, isCore bool) error {
    entries, err := os.ReadDir(dir)
    if err != nil {
        if os.IsNotExist(err) {
            return os.MkdirAll(dir, 0755)
        }
        return err
    }

    for _, entry := range entries {
        if entry.IsDir() && isCore {
            // For core themes, look for .so files in subdirectories
            soPath := filepath.Join(dir, entry.Name(), entry.Name()+".so")
            if err := loadTheme(soPath, entry.Name()); err != nil {
                log.Printf("Failed to load theme from %s: %v", soPath, err)
            }
        } else if !entry.IsDir() && filepath.Ext(entry.Name()) == ".so" {
            // For common themes, load .so files directly
            themeName := strings.TrimSuffix(entry.Name(), ".so")
            themePath := filepath.Join(dir, entry.Name())
            if err := loadTheme(themePath, themeName); err != nil {
                log.Printf("Failed to load theme from %s: %v", themePath, err)
            }
        }
    }

    return nil
}

// loadTheme loads a single theme from a .so file
func loadTheme(themePath, name string) error {
    log.Printf("Loading theme from: %s", themePath)
    if _, err := os.Stat(themePath); err != nil {
        return fmt.Errorf("theme file does not exist: %s", themePath)
    }

    p, err := plugin.Open(themePath)
    if err != nil {
        return fmt.Errorf("failed to open theme plugin: %v", err)
    }

    symTheme, err := p.Lookup("Theme")
    if err != nil {
        return fmt.Errorf("failed to lookup Theme symbol: %v", err)
    }

    log.Printf("Theme symbol type: %T", symTheme)

    // Convert any valid theme type to types.Theme interface
    var theme types.Theme

    // Try to get the theme interface
    if t, ok := symTheme.(types.Theme); ok {
        theme = t
    } else if ptr, ok := symTheme.(*types.Theme); ok {
        theme = *ptr
    } else {
        // Try to use reflection to get a pointer to the struct if needed
        valueOf := reflect.ValueOf(symTheme)
        if valueOf.Kind() == reflect.Ptr {
            if iface := valueOf.Interface(); iface != nil {
                if t, ok := iface.(types.Theme); ok {
                    theme = t
                }
            }
        } else if valueOf.CanAddr() {
            if iface := valueOf.Addr().Interface(); iface != nil {
                if t, ok := iface.(types.Theme); ok {
                    theme = t
                }
            }
        }
    }

    if theme == nil {
        log.Printf("Found symbol of type %T", symTheme)
        log.Printf("Value: %+v", symTheme)
        return fmt.Errorf("unexpected type from theme symbol")
    }

    RegisterTheme(name, theme)
    return nil
}

// InstallTheme installs a theme from a .so file into the common themes directory
func InstallTheme(soFile string) error {
    commonThemeDir := filepath.Join(themeDirectory, "common")
    if err := os.MkdirAll(commonThemeDir, 0755); err != nil {
        return err
    }

    // Copy the .so file to the common themes directory
    destPath := filepath.Join(commonThemeDir, filepath.Base(soFile))
    if err := copyFile(soFile, destPath); err != nil {
        return err
    }

    // Load the newly installed theme
    themeName := strings.TrimSuffix(filepath.Base(soFile), ".so")
    return loadTheme(destPath, themeName)
}

// UninstallTheme removes a theme from the common themes directory
func UninstallTheme(name string) error {
    themeMutex.Lock()
    defer themeMutex.Unlock()

    if _, exists := themes[name]; !exists {
        return fmt.Errorf("theme %s not found", name)
    }

    delete(themes, name)

    // Remove the .so file
    commonThemeDir := filepath.Join(themeDirectory, "common")
    soFile := filepath.Join(commonThemeDir, name+".so")
    return os.Remove(soFile)
}

// Helper function to copy files
func copyFile(src, dst string) error {
    input, err := os.ReadFile(src)
    if err != nil {
        return err
    }
    return os.WriteFile(dst, input, 0644)
}

// Existing theme management functions
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