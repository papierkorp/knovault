// internal/pluginManager/pluginManagerBuiltinHelper.go
package pluginManager

import (
    "knovault/internal/types"
    "knovault/internal/pluginManager/builtin/filemanager"
    "knovault/internal/pluginManager/builtin/markdownparser"
    "knovault/internal/pluginManager/builtin/themechanger"
)

func loadBuiltinPlugins() (map[string]types.Plugin, error) {
    plugins := make(map[string]types.Plugin)

    // Initialize builtin plugins
    plugins["FileManager"] = &filemanager.FileManagerPlugin{}
    plugins["MarkdownParser"] = &markdownparser.MarkdownParserPlugin{}
    plugins["ThemeChanger"] = &themechanger.ThemeChangerPlugin{}

    return plugins, nil
}