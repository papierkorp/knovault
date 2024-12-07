package pluginManager

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "plugin"
    "knovault/internal/types"
)

func loadExternalPlugins() (map[string]types.Plugin, error) {
    plugins := make(map[string]types.Plugin)
    externalDir := "./internal/pluginManager/external"

    // First, compile any main.go files to .so
    err := compileExternalPlugins(externalDir)
    if err != nil {
        log.Printf("Warning: Error compiling external plugins: %v", err)
    }

    // Then load all .so files
    err = filepath.Walk(externalDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() && filepath.Ext(path) == ".so" {
            pluginName := filepath.Base(filepath.Dir(path))
            plug, err := loadPluginFromFile(path)
            if err != nil {
                log.Printf("Warning: Could not load plugin %s: %v", pluginName, err)
                return nil
            }
            plugins[pluginName] = plug
        }
        return nil
    })

    return plugins, err
}

func compileExternalPlugins(dir string) error {
    return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() && filepath.Base(path) == "main.go" {
            pluginDir := filepath.Dir(path)
            pluginName := filepath.Base(filepath.Dir(pluginDir))
            outPath := filepath.Join(pluginDir, pluginName+".so")

            // Remove existing .so file
            os.Remove(outPath)

            cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", outPath, path)
            if err := cmd.Run(); err != nil {
                log.Printf("Warning: Could not compile plugin %s: %v", pluginName, err)
            }
        }
        return nil
    })
}

func loadPluginFromFile(path string) (types.Plugin, error) {
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
        return nil, fmt.Errorf("invalid plugin type")
    }

    return p, nil
}