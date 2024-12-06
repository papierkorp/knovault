package assetManager

import (
    "fmt"
    "log"
    "path/filepath"
    "plugin"

    "knovault/internal/types"
)

func (am *AssetManager) loadPlugins() error {
    config, err := loadConfig[types.PluginConfig]("internal/assetManager/plugins_list.json")
    if err != nil {
        return err
    }

    am.mutex.Lock()
    defer am.mutex.Unlock()

    pluginsDir := "./internal/assetManager/plugins"
    files, err := findAssetFiles(pluginsDir)
    if err != nil {
        return fmt.Errorf("failed to scan plugins directory: %v", err)
    }

    // Create a map of configured plugins
    configuredPlugins := make(map[string]types.PluginMetadata)
    for _, p := range config.Plugins {
        if p.Enabled {
            configuredPlugins[p.Name] = p
            am.pluginInfo[p.Name] = p
        }
    }

    // Load plugins from found files
    for _, file := range files {
        // For .so files, get the plugin name from the parent directory
        // For main.go files, get the plugin name from the grandparent directory
        var pluginName string
        if filepath.Ext(file) == ".so" {
            pluginName = filepath.Base(filepath.Dir(file))
        } else {
            pluginName = filepath.Base(filepath.Dir(filepath.Dir(file)))
        }

        // Check if this plugin is configured and enabled
        if metadata, ok := configuredPlugins[pluginName]; ok {
            if err := am.loadPluginFromPath(pluginName, file); err != nil {
                log.Printf("Warning: Could not load plugin %s: %v", pluginName, err)
                continue
            }
            am.pluginInfo[pluginName] = metadata
        }
    }

    return nil
}

func (am *AssetManager) loadPluginFromPath(name, path string) error {
    // For .so files, load using plugin.Open
    if filepath.Ext(path) == ".so" {
        plug, err := plugin.Open(path)
        if err != nil {
            return fmt.Errorf("could not open plugin: %v", err)
        }

        symPlugin, err := plug.Lookup("Plugin")
        if err != nil {
            return fmt.Errorf("could not find Plugin symbol: %v", err)
        }

        // Try both direct interface and pointer conversion
        if pluginInstance, ok := symPlugin.(types.Plugin); ok {
            am.plugins[name] = pluginInstance
        } else if pluginPtr, ok := symPlugin.(*types.Plugin); ok {
            am.plugins[name] = *pluginPtr
        } else {
            return fmt.Errorf("invalid plugin type")
        }
    }

    log.Printf("Loaded plugin: %s", name)
    return nil
}