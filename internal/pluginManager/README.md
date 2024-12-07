# Plugin Manager

The Plugin Manager handles plugin management for Knovault. This directory contains all plugin implementations and manages both builtin and external plugins.

## Directory Structure

```
pluginManager/
├── builtin/              # Builtin plugin implementations
│   ├── filemanager/
│   ├── markdownparser/
│   └── themechanger/
├── external/             # External plugin implementations
│   ├── CustomCSS/
│   └── HelloWorld/
├── pluginManager.go          # Core plugin manager
├── pluginManagerBuiltinHelper.go
└── pluginManagerExternalHelper.go
```

## Plugin Types

### Builtin Plugins

- Source code is directly in the builtin directory
- Always available and loaded by default
- Core functionality plugins

Available builtin plugins:

- FileManager: File system operations
- MarkdownParser: Markdown processing
- ThemeChanger: Theme switching functionality

### External Plugins

- Located in the external directory
- Can be compiled to .so files
- Can be added/removed without changing core code
- Must follow plugin interface requirements

Available external plugins:

- CustomCSS: Custom styling support
- HelloWorld: Example plugin

## Creating a New Plugin

### Builtin Plugin

1. Create a new directory under `builtin/`:

   ```
   builtin/
   └── yourplugin/
       └── yourplugin.go    # Plugin implementation
   ```

2. Implement the Plugin interface:

   ```go
   type YourPlugin struct{}

   func (p *YourPlugin) Name() string
   func (p *YourPlugin) Description() string
   func (p *YourPlugin) Help() string
   func (p *YourPlugin) TemplResponse() (templ.Component, error)
   func (p *YourPlugin) JsonResponse() ([]byte, error)
   func (p *YourPlugin) Execute(params map[string]string) (interface{}, error)
   ```

### External Plugin

1. Create a new directory under `external/`:

   ```
   external/
   └── YourPlugin/
       ├── main.go          # Plugin entry point
       ├── plugin/
       │   └── plugin.go    # Plugin implementation
       └── templates/       # Optional templates
   ```

2. Create main.go:

   ```go
   package main

   import (
       "knovault/internal/types"
       "knovault/internal/pluginManager/external/YourPlugin/plugin"
   )

   var Plugin types.Plugin = &plugin.YourPlugin{}

   func main() {}
   ```

3. Implement the plugin interface in plugin.go

## Optional Interfaces

Plugins can implement additional interfaces for extended functionality:

### PluginWithRoute

For plugins that need HTTP routes:

```go
type PluginWithRoute interface {
    Plugin
    Route() PluginRoute
}
```

### PluginWithTemplateExtensions

For plugins that extend templates:

```go
type PluginWithTemplateExtensions interface {
    Plugin
    ExtendTemplate(templateName string) (templ.Component, error)
}
```

## Building External Plugins

External plugins are automatically compiled to .so files when detected:

1. The plugin manager scans the external directory
2. Finds main.go files
3. Compiles them to .so files
4. Loads the resulting plugins

You can also manually compile plugins:

```bash
go build -buildmode=plugin -o yourplugin.so main.go
```

## Development Tips

- Use `go run` for testing builtin plugins
- Test external plugins by building the .so file
- Keep plugin dependencies minimal
- Document your plugin's functionality
- Provide clear error messages
- Use template extensions when modifying UI
- Follow the existing plugin patterns for consistency
