// internal/pluginManager/pluginManagerBuiltinHelper.go
package pluginManager

import (
    "fmt"
    "log"
    "path/filepath"
    "knovault/internal/types"
    "knovault/internal/pluginManager/builtin/filemanager"
    "knovault/internal/pluginManager/builtin/markdownparser"
    "knovault/internal/pluginManager/builtin/themechanger"
    "knovault/internal/pluginManager/builtin/cssswitcher"
)

func loadBuiltinPlugins() (map[string]types.Plugin, error) {
    plugins := make(map[string]types.Plugin)
    builtinDir := "./internal/pluginManager/builtin"

    absBuiltinDir, err := filepath.Abs(builtinDir)
    if err != nil {
        return nil, fmt.Errorf("failed to get absolute path: %v", err)
    }
    log.Printf("Loading builtin plugins from: %s", absBuiltinDir)

    // Initialize the plugins with consistent naming
    plugins["FileManager"] = &filemanager.FileManagerPlugin{}
    plugins["MarkdownParser"] = &markdownparser.MarkdownParserPlugin{}
    plugins["ThemeChanger"] = &themechanger.ThemeChangerPlugin{}
    plugins["CSSSwitcher"] = &cssswitcher.CSSSwitcherPlugin{}

    log.Printf("✓ Loaded %d builtin plugins", len(plugins))
    for name, plugin := range plugins {
        log.Printf("✓ Loaded builtin plugin: %s - %s", name, plugin.Description())
    }

    return plugins, nil
}