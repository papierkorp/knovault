package plugins

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"pewito/internal/types"
	_ "pewito/internal/plugins/templates"
	"plugin"
	"sync"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
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

func ApplyPluginRoutes(e *echo.Echo) {
	for _, p := range corePlugins {
		if routePlugin, ok := p.(types.PluginWithRoute); ok {
			route := routePlugin.Route()
			e.Add(route.Method, route.Path, route.Handler)
		}
	}

	for _, p := range commonPlugins {
		if routePlugin, ok := p.(types.PluginWithRoute); ok {
			route := routePlugin.Route()
			e.Add(route.Method, route.Path, route.Handler)
		}
	}
}

func GetPluginTemplateExtensions(templateName string) []templ.Component {
	var extensions []templ.Component
	
	// Iterate over core plugins
	for _, plugin := range corePlugins {
		if extendablePlugin, ok := plugin.(types.PluginWithTemplateExtensions); ok {
			if extension, err := extendablePlugin.ExtendTemplate(templateName); err == nil {
				extensions = append(extensions, extension)
			}
		}
	}
	
	// Iterate over common plugins
	for _, plugin := range commonPlugins {
		if extendablePlugin, ok := plugin.(types.PluginWithTemplateExtensions); ok {
			if extension, err := extendablePlugin.ExtendTemplate(templateName); err == nil {
				extensions = append(extensions, extension)
			}
		}
	}
	
	return extensions
}

func handlePluginExecute(c echo.Context) error {
	pluginName := c.Param("pluginName")
	plugin, ok := GetPlugin(pluginName)
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Plugin not found"})
	}

	// Collect all form values as parameters
	params := make(map[string]string)
	formParams, err := c.FormParams()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse form parameters"})
	}
	for key, values := range formParams {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	// Execute the plugin
	response, err := plugin.Execute(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Handle different response types
	switch response := response.(type) {
	case []byte:
		return c.Blob(http.StatusOK, "application/json", response)
	case templ.Component:
		return response.Render(c.Request().Context(), c.Response().Writer)
	default:
		return c.JSON(http.StatusOK, response)
	}
}

func ExecutePlugin(name string, params map[string]string) (interface{}, error) {
	plugin, ok := GetPlugin(name)
	if !ok {
		return nil, fmt.Errorf("plugin %s not found", name)
	}

	return plugin.Execute(params)
}