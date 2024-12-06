package assetManager

import (
    "fmt"
    "log"
    "path/filepath"
    "plugin"
    "os"

    "knovault/internal/types"
)

func (am *AssetManager) loadConfiguredPlugins() error {
    config, err := loadConfig[types.PluginConfig]("internal/assetManager/plugins_list.json")
    if err != nil {
        return err
    }

    am.mutex.Lock()
    defer am.mutex.Unlock()

    for _, p := range config.Plugins {
        if !p.Enabled {
            continue
        }

        // Store metadata regardless of loading success
        am.pluginInfo[p.Name] = p

        if hasTag(p.Tags, "built-in") {
            pluginPath := filepath.Join(p.Path)
            soPath := filepath.Join(filepath.Dir(pluginPath), "plugin.so")
            if err := compileModule(pluginPath, soPath); err != nil {
                log.Printf("Warning: Could not compile plugin %s: %v", p.Name, err)
                continue
            }

            if err := am.loadPluginFromPath(p.Name, soPath); err != nil {
                log.Printf("Warning: Could not load plugin %s: %v", p.Name, err)
                continue
            }
        }
    }

    return nil
}

func (am *AssetManager) loadPluginFromPath(name, path string) error {
    plug, err := plugin.Open(path)
    if err != nil {
        return fmt.Errorf("could not open plugin: %v", err)
    }

    // Get the Plugin symbol
    symPlugin, err := plug.Lookup("Plugin")
    if err != nil {
        return fmt.Errorf("could not find Plugin symbol: %v", err)
    }

    // Check if it's a pointer to a Plugin interface
    pluginPtr, ok := symPlugin.(*types.Plugin)
    if !ok {
        // Try direct interface conversion if pointer conversion fails
        pluginInstance, ok := symPlugin.(types.Plugin)
        if !ok {
            return fmt.Errorf("invalid plugin type")
        }
        am.plugins[name] = pluginInstance
        log.Printf("Loaded plugin: %s", name)
        return nil
    }

    // If we got a pointer, dereference it
    am.plugins[name] = *pluginPtr
    log.Printf("Loaded plugin: %s", name)
    return nil
}

func (am *AssetManager) loadCompiledPlugins() error {
    pluginsDir := "./internal/assetManager/plugins"
    entries, err := os.ReadDir(pluginsDir)
    if err != nil {
        return fmt.Errorf("failed to read plugins directory: %v", err)
    }

    for _, entry := range entries {
        if !entry.IsDir() && filepath.Ext(entry.Name()) == ".so" {
            name := filepath.Base(entry.Name())
            name = name[:len(name)-3] // Remove .so extension
            path := filepath.Join(pluginsDir, entry.Name())

            if err := am.loadPluginFromPath(name, path); err != nil {
                log.Printf("Failed to load plugin %s: %v", name, err)
            }
        }
    }

    return nil
}

// Plugin management methods
func (am *AssetManager) GetPlugin(name string) (types.Plugin, bool) {
    am.mutex.RLock()
    defer am.mutex.RUnlock()
    p, ok := am.plugins[name]
    return p, ok
}

func (am *AssetManager) ListPlugins() []types.PluginInfo {
    am.mutex.RLock()
    defer am.mutex.RUnlock()

    var plugins []types.PluginInfo
    for name, p := range am.plugins {
        info := types.PluginInfo{
            Name:        name,
            Description: p.Description(),
        }
        if metadata, ok := am.pluginInfo[name]; ok {
            info.Tags = metadata.Tags
        }
        plugins = append(plugins, info)
    }
    return plugins
}

func (am *AssetManager) GetPluginsByTag(tag string) []types.PluginInfo {
    am.mutex.RLock()
    defer am.mutex.RUnlock()

    var plugins []types.PluginInfo
    for name, metadata := range am.pluginInfo {
        if hasTag(metadata.Tags, tag) && am.plugins[name] != nil {
            plugins = append(plugins, types.PluginInfo{
                Name:        name,
                Description: am.plugins[name].Description(),
                Tags:        metadata.Tags,
            })
        }
    }
    return plugins
}