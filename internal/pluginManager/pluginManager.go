// internal/pluginManager/pluginManager.go
package pluginManager

import (
    "fmt"
    "log"
    "sync"
    "knovault/internal/types"
)

type PluginManager struct {
    plugins    map[string]types.Plugin
    pluginInfo map[string]types.PluginInfo
    mutex      sync.RWMutex
}

func NewPluginManager() *PluginManager {
    return &PluginManager{
        plugins:    make(map[string]types.Plugin),
        pluginInfo: make(map[string]types.PluginInfo),
    }
}

func (pm *PluginManager) Initialize() error {
    pm.mutex.Lock()
    defer pm.mutex.Unlock()

    log.Println("Loading builtin plugins...")
    builtinPlugins, err := loadBuiltinPlugins()
    if err != nil {
        log.Printf("⚠️  Warning: Error loading builtin plugins: %v", err)
    }
    for name, plugin := range builtinPlugins {
        log.Printf("✓ Loaded builtin plugin: %s - %s", name, plugin.Description())
        pm.plugins[name] = plugin
        pm.pluginInfo[name] = types.PluginInfo{
            Name:        name,
            Description: plugin.Description(),
            Tags:        []string{"builtin"},
        }
    }

    log.Println("Loading external plugins...")
    externalPlugins, err := loadExternalPlugins()
    if err != nil {
        log.Printf("⚠️  Warning: Error loading external plugins: %v", err)
    }
    for name, plugin := range externalPlugins {
        log.Printf("✓ Loaded external plugin: %s - %s", name, plugin.Description())
        pm.plugins[name] = plugin
        pm.pluginInfo[name] = types.PluginInfo{
            Name:        name,
            Description: plugin.Description(),
            Tags:        []string{"external"},
        }
    }

    if len(pm.plugins) == 0 {
        return fmt.Errorf("❌ No plugins were loaded")
    }

    log.Printf("✓ Plugin Manager initialized with %d plugins", len(pm.plugins))
    return nil
}

func (pm *PluginManager) GetPlugin(name string) (types.Plugin, bool) {
    pm.mutex.RLock()
    defer pm.mutex.RUnlock()
    p, ok := pm.plugins[name]
    return p, ok
}

func (pm *PluginManager) ListPlugins() []types.PluginInfo {
    pm.mutex.RLock()
    defer pm.mutex.RUnlock()

    var plugins []types.PluginInfo
    for name, p := range pm.plugins {
        info := types.PluginInfo{
            Name:        name,
            Description: p.Description(),
        }
        if metadata, ok := pm.pluginInfo[name]; ok {
            info.Tags = metadata.Tags
        }
        plugins = append(plugins, info)
    }
    return plugins
}



