// internal/pluginManager/pluginManagerExternalHelper.go
package pluginManager

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

func loadExternalPlugins() (map[string]types.Plugin, error) {
    plugins := make(map[string]types.Plugin)
    externalDir := "./internal/pluginManager/external"

    absExternalDir, err := filepath.Abs(externalDir)
    if err != nil {
        return nil, fmt.Errorf("failed to get absolute path: %v", err)
    }
    log.Printf("Loading plugins from: %s", absExternalDir)

    log.Println("Compiling external plugins...")
    err = compileExternalPlugins(absExternalDir)
    if err != nil {
        log.Printf("⚠️  Warning: Error compiling external plugins: %v", err)
    }

    files, err := os.ReadDir(absExternalDir)
    if err != nil {
        log.Printf("Error reading external directory: %v", err)
        return plugins, err
    }

    for _, file := range files {
        if !file.IsDir() {
            continue
        }

        pluginName := file.Name()
        pluginDir := filepath.Join(absExternalDir, pluginName)
        mainPath := filepath.Join(pluginDir, "main.go")
        soPath := filepath.Join(pluginDir, strings.ToLower(pluginName)+".so")

        log.Printf("Processing plugin directory: %s", pluginDir)
        log.Printf("Looking for main.go at: %s", mainPath)
        log.Printf("Looking for .so file at: %s", soPath)

        if mainFileInfo, err := os.Stat(mainPath); err == nil && !mainFileInfo.IsDir() {
            log.Printf("Found main.go for plugin: %s", pluginName)

            log.Printf("Loading plugin: %s", pluginName)
            plug, err := loadPluginFromFile(soPath)
            if err != nil {
                log.Printf("⚠️  Could not load plugin %s: %v", pluginName, err)
                continue
            }

            plugins[pluginName] = plug
            log.Printf("✓ Successfully loaded: %s", pluginName)
            log.Printf("✓ Loaded external plugin: %s - %s", pluginName, plug.Description())
        } else {
            log.Printf("Skipping directory %s - not a valid plugin (no main.go found)", pluginName)
        }
    }

    return plugins, nil
}

func compileExternalPlugins(absDir string) error {
    files, err := os.ReadDir(absDir)
    if err != nil {
        return err
    }

    for _, file := range files {
        if !file.IsDir() {
            continue
        }

        pluginName := file.Name()
        mainPath := filepath.Join(absDir, pluginName, "main.go")
        outPath := filepath.Join(absDir, pluginName, strings.ToLower(pluginName)+".so")

        log.Printf("Checking plugin directory: %s", pluginName)
        log.Printf("Looking for main.go at: %s", mainPath)

        if mainFileInfo, err := os.Stat(mainPath); err == nil && !mainFileInfo.IsDir() {
            log.Printf("Compiling plugin: %s", pluginName)

            // Remove existing .so file
            os.Remove(outPath)

            cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", outPath, mainPath)
            cmd.Dir = filepath.Dir(mainPath)

            output, err := cmd.CombinedOutput()
            if err != nil {
                log.Printf("⚠️  Could not compile plugin %s: %v\nOutput: %s", pluginName, err, string(output))
                continue
            }
            log.Printf("✓ Successfully compiled plugin: %s", pluginName)
        } else {
            log.Printf("No main.go found for plugin %s", pluginName)
        }
    }

    return nil
}

func loadPluginFromFile(path string) (types.Plugin, error) {
    // Check if file exists before attempting to open
    if _, err := os.Stat(path); err != nil {
        return nil, fmt.Errorf("plugin file does not exist: %v", err)
    }

    plug, err := plugin.Open(path)
    if err != nil {
        return nil, fmt.Errorf("could not open plugin: %v", err)
    }

    symPlugin, err := plug.Lookup("Plugin")
    if err != nil {
        return nil, fmt.Errorf("could not find Plugin symbol: %v", err)
    }

    var p types.Plugin
    switch v := symPlugin.(type) {
    case types.Plugin:
        p = v
    case *types.Plugin:
        p = *v
    default:
        return nil, fmt.Errorf("invalid plugin type: %T", symPlugin)
    }

    return p, nil
}