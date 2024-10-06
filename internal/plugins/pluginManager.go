package plugins

import (
	"fmt"
	"pewitima/internal/types"
	"os"
	"path/filepath"
	"plugin"
	"sync"
)

var (
	corePlugins     = make(map[string]types.Plugin)
	commonPlugins   = make(map[string]types.Plugin)
	pluginMutex     sync.RWMutex
	pluginDirectory = "./internal/plugins"
)

func RegisterCorePlugin(name string, p types.Plugin) {
	pluginMutex.Lock()
	defer pluginMutex.Unlock()
	corePlugins[name] = p
}

func LoadCommonPlugins() error {
	commonPluginDir := filepath.Join(pluginDirectory, "common")
	entries, err := os.ReadDir(commonPluginDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".so" {
			pluginPath := filepath.Join(commonPluginDir, entry.Name())
			p, err := plugin.Open(pluginPath)
			if err != nil {
				return err
			}

			symPlugin, err := p.Lookup("Plugin")
			if err != nil {
				return err
			}

			var pl types.Plugin
			pl, ok := symPlugin.(types.Plugin)
			if !ok {
				return fmt.Errorf("unexpected type from module symbol")
			}

			pluginMutex.Lock()
			commonPlugins[pl.Name()] = pl
			pluginMutex.Unlock()
		}
	}

	return nil
}

func GetPlugin(name string) (types.Plugin, bool) {
	pluginMutex.RLock()
	defer pluginMutex.RUnlock()

	if p, ok := corePlugins[name]; ok {
		return p, true
	}
	if p, ok := commonPlugins[name]; ok {
		return p, true
	}
	return nil, false
}

func ListPlugins() []types.PluginInfo {
	pluginMutex.RLock()
	defer pluginMutex.RUnlock()

	var plugins []types.PluginInfo

	for name, p := range corePlugins {
		plugins = append(plugins, types.PluginInfo{
			Name:        name,
			Description: p.Description(),
			Type:        "core",
		})
	}

	for name, p := range commonPlugins {
		plugins = append(plugins, types.PluginInfo{
			Name:        name,
			Description: p.Description(),
			Type:        "common",
		})
	}

	return plugins
}

func InstallCommonPlugin(name string) error {
	// Implementation for installing a common plugin
	// This could involve downloading the plugin, verifying it, and then loading it
	return nil
}

func UninstallCommonPlugin(name string) error {
	pluginMutex.Lock()
	defer pluginMutex.Unlock()

	if _, exists := commonPlugins[name]; !exists {
		return fmt.Errorf("plugin %s not found", name)
	}

	delete(commonPlugins, name)
	// Additional cleanup if necessary

	return nil
}

func ExecutePlugin(name string, params map[string]string) (interface{}, error) {
	plugin, ok := GetPlugin(name)
	if !ok {
		return nil, fmt.Errorf("plugin %s not found", name)
	}

	return plugin.Execute(params)
}