package plugins

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"sync"
	"knovault/internal/types"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

var (
	corePlugins     = make(map[string]types.Plugin)
	commonPlugins   = make(map[string]types.Plugin)
	pluginMutex     sync.RWMutex
	pluginDirectory = "./internal/plugins"
)

// compileCorePlugins compiles all core plugins into .so files
func compileCorePlugins() error {
    absPluginDir, err := filepath.Abs(pluginDirectory)
    if err != nil {
        return fmt.Errorf("failed to get absolute path: %v", err)
    }
    corePluginDir := filepath.Join(absPluginDir, "core")
    log.Printf("Absolute core plugins directory: %s", corePluginDir)

    entries, err := os.ReadDir(corePluginDir)
    if err != nil {
        return fmt.Errorf("failed to read core plugin directory: %v", err)
    }

    for _, entry := range entries {
        if entry.IsDir() {
            pluginName := entry.Name()
            pluginDir := filepath.Join(corePluginDir, pluginName)
            outputFile := filepath.Join(pluginDir, pluginName+".so")

            log.Printf("Compiling plugin %s", pluginName)
            log.Printf("Plugin directory: %s", pluginDir)
            log.Printf("Output file: %s", outputFile)

            // Build the plugin
            cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", outputFile, "./main.go")
            cmd.Dir = pluginDir

            if output, err := cmd.CombinedOutput(); err != nil {
                log.Printf("Failed to compile plugin %s: %v\nOutput: %s", pluginName, err, string(output))
                continue
            }

            // Verify file exists after compilation
            if _, err := os.Stat(outputFile); err != nil {
                log.Printf("Plugin %s compiled but file not found at %s: %v", pluginName, outputFile, err)
                continue
            } else {
                log.Printf("Verified plugin file exists at: %s", outputFile)
            }

            // Print file size to confirm it's a valid file
            if info, err := os.Stat(outputFile); err == nil {
                log.Printf("Plugin file size: %d bytes", info.Size())
            }
        }
    }
    return nil
}



// LoadCorePlugins compiles and loads all core plugins
func LoadCorePlugins() error {
	if err := compileCorePlugins(); err != nil {
		return err
	}

	corePluginDir := filepath.Join(pluginDirectory, "core")
	return loadPluginsFromDir(corePluginDir, true)
}

// LoadCommonPlugins loads all plugins from the common plugins directory
func LoadCommonPlugins() error {
	commonPluginDir := filepath.Join(pluginDirectory, "common")
	return loadPluginsFromDir(commonPluginDir, false)
}

func loadPlugin(pluginPath string, isCore bool) error {
    absPath, err := filepath.Abs(pluginPath)
    if err != nil {
        return fmt.Errorf("failed to get absolute path: %v", err)
    }

    log.Printf("Attempting to load plugin from: %s", absPath)

    if _, err := os.Stat(absPath); err != nil {
        return fmt.Errorf("plugin file does not exist at %s: %v", absPath, err)
    }

    p, err := plugin.Open(absPath)
    if err != nil {
        return fmt.Errorf("failed to open plugin at %s: %v", absPath, err)
    }

    symPlugin, err := p.Lookup("Plugin")
    if err != nil {
        return fmt.Errorf("failed to lookup Plugin symbol: %v", err)
    }

    log.Printf("Plugin symbol type: %T", symPlugin)

    var pl types.Plugin

    // Try different type assertions
    switch v := symPlugin.(type) {
    case types.Plugin:
        pl = v
    case *types.Plugin:
        pl = *v
    default:
        // If it's a pointer to a struct that implements Plugin
        if plugin, ok := v.(types.Plugin); ok {
            pl = plugin
        } else {
            log.Printf("Found symbol of type %T", symPlugin)
            return fmt.Errorf("unexpected type from module symbol")
        }
    }

    pluginMutex.Lock()
    defer pluginMutex.Unlock()

    if isCore {
        corePlugins[pl.Name()] = pl
        log.Printf("Successfully loaded core plugin: %s", pl.Name())
    } else {
        commonPlugins[pl.Name()] = pl
        log.Printf("Successfully loaded common plugin: %s", pl.Name())
    }

    return nil
}

// loadPluginsFromDir loads all .so files from the specified directory
func loadPluginsFromDir(dir string, isCore bool) error {
    absDir, err := filepath.Abs(dir)
    if err != nil {
        return fmt.Errorf("failed to get absolute path: %v", err)
    }

    log.Printf("Loading plugins from directory: %s", absDir)

    entries, err := os.ReadDir(absDir)
    if err != nil {
        if os.IsNotExist(err) {
            log.Printf("Creating directory: %s", absDir)
            return os.MkdirAll(absDir, 0755)
        }
        return err
    }

    for _, entry := range entries {
        if entry.IsDir() && isCore {
            // For core plugins, look for .so files in subdirectories
            soPath := filepath.Join(absDir, entry.Name(), entry.Name()+".so")
            if err := loadPlugin(soPath, isCore); err != nil {
                log.Printf("Failed to load core plugin from %s: %v", soPath, err)
            }
        } else if !entry.IsDir() && filepath.Ext(entry.Name()) == ".so" {
            // For common plugins, load .so files directly
            pluginPath := filepath.Join(absDir, entry.Name())
            if err := loadPlugin(pluginPath, isCore); err != nil {
                log.Printf("Failed to load common plugin from %s: %v", pluginPath, err)
            }
        }
    }

    return nil
}

// GetPlugin retrieves a plugin by name from either core or common plugins
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

// ListPlugins returns a list of all available plugins
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

// InstallCommonPlugin installs a plugin from a .so file
func InstallCommonPlugin(soFile string) error {
	commonPluginDir := filepath.Join(pluginDirectory, "common")
	if err := os.MkdirAll(commonPluginDir, 0755); err != nil {
		return err
	}

	// Copy the .so file to the common plugins directory
	destPath := filepath.Join(commonPluginDir, filepath.Base(soFile))
	if err := copyFile(soFile, destPath); err != nil {
		return err
	}

	// Load the newly installed plugin
	return loadPlugin(destPath, false)
}

// UninstallCommonPlugin removes a plugin from the common plugins directory
func UninstallCommonPlugin(name string) error {
	pluginMutex.Lock()
	defer pluginMutex.Unlock()

	if _, exists := commonPlugins[name]; !exists {
		return fmt.Errorf("plugin %s not found", name)
	}

	delete(commonPlugins, name)

	// Remove the .so file
	commonPluginDir := filepath.Join(pluginDirectory, "common")
	soFile := filepath.Join(commonPluginDir, name+".so")
	return os.Remove(soFile)
}

// copyFile helper function to copy a file
func copyFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, input, 0644)
}

// ApplyPluginRoutes applies all plugin routes to the Echo instance
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

// GetPluginTemplateExtensions returns template extensions from all plugins for a given template
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