# Plugins Directory

This directory contains all plugins for the application. Each plugin follows a standard Go package structure that allows both direct source code usage and compilation to .so files when needed.

## Directory Structure

```
plugins/
  CustomCSS/
    plugin/             # Package plugin contains implementation
      plugin.go
    main.go            # Package main for compilation
    templates/
      custom_css_editor.templ
  FileManager/
    plugin/
      plugin.go
    main.go
  HelloWorld/
    plugin/
      plugin.go
    main.go
```

## Using Plugins

Plugins can be used in two ways:

1. Direct Source Code (Default)

   - Import the plugin package directly: `import "knovault/internal/pluginManager/plugins/CustomCSS/plugin"`
   - No compilation needed
   - Better for development and debugging

2. Compiled Plugin (.so file)
   - Compile the plugin to a .so file
   - Useful for distribution or runtime loading

## Compiling Plugins

To compile a plugin to a .so file:

```bash
# Using make (recommended)
make compile-plugin PLUGIN=CustomCSS

# Or manually from the plugin directory
cd internal/pluginManager/plugins/PluginName
go build -buildmode=plugin -o PluginName.so main.go
```

Note: Compiled plugins (.so files) will only be loaded if a direct source implementation isn't already available.
