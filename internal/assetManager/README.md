# Asset Manager

The Asset Manager handles plugins and themes for Knovault. This directory contains all plugin and theme implementations, along with their configurations.

## Directory Structure

```
assetManager/
├── plugins/              # Plugin implementations
│   ├── CustomCSS/
│   ├── FileManager/
│   ├── MarkdownParser/
│   └── ThemeChanger/
├── themes/              # Theme implementations
│   ├── defaultTheme/
│   └── dark/
├── plugins_list.json    # Plugin configuration
└── themes_list.json     # Theme configuration
```

## Creating a New Plugin

1. Create a new directory under `plugins/`:

   ```
   plugins/
   └── YourPlugin/
       ├── main.go          # Plugin entry point
       ├── plugin/
       │   └── plugin.go    # Plugin implementation
       └── templates/       # Optional templates
   ```

2. Implement the Plugin interface in `plugin.go`:

   ```go
   type YourPlugin struct{}

   func (p *YourPlugin) Name() string
   func (p *YourPlugin) Description() string
   func (p *YourPlugin) Help() string
   func (p *YourPlugin) TemplResponse() (templ.Component, error)
   func (p *YourPlugin) JsonResponse() ([]byte, error)
   func (p *YourPlugin) Execute(params map[string]string) (interface{}, error)
   ```

3. Add plugin configuration to `plugins_list.json`:
   ```json
   {
     "name": "YourPlugin",
     "path": "internal/assetManager/plugins/YourPlugin/plugin",
     "enabled": true,
     "tags": ["your-tag"]
   }
   ```

## Creating a New Theme

1. Create a new directory under `themes/`:

   ```
   themes/
   └── YourTheme/
       ├── main.go           # Theme entry point
       ├── plugin/
       │   └── plugin.go     # Theme implementation
       └── templates/        # Theme templates
           ├── layout/
           └── pages/
   ```

2. Implement the Theme interface in `plugin.go`:

   ```go
   type YourTheme struct{}

   func (t *YourTheme) Home() (templ.Component, error)
   func (t *YourTheme) Help() (templ.Component, error)
   func (t *YourTheme) Settings() (templ.Component, error)
   func (t *YourTheme) Search() (templ.Component, error)
   func (t *YourTheme) DocsRoot() (templ.Component, error)
   func (t *YourTheme) Docs(content string) (templ.Component, error)
   func (t *YourTheme) Playground() (templ.Component, error)
   func (t *YourTheme) Plugins() (templ.Component, error)
   ```

3. Add theme configuration to `themes_list.json`:
   ```json
   {
     "name": "YourTheme",
     "path": "internal/assetManager/themes/YourTheme/plugin",
     "enabled": true,
     "tags": ["your-tag"]
   }
   ```

## Compiling Assets

### Plugins

```bash
make compile-plugin PLUGIN=YourPlugin
```

### Themes

```bash
make compile-theme THEME=YourTheme
```

## Plugin Interfaces

Additional plugin interfaces are available for extended functionality:

- `PluginWithRoute`: For plugins that need HTTP routes
- `PluginWithTemplateExtensions`: For plugins that extend templates

See `types/types.go` for detailed interface definitions.

## Development Tips

- Use the `dev` make target for live reloading during development
- Plugin and theme changes require recompilation
- Check build errors in `tmp/build-errors.log`
- Use `make clean-assets` to remove compiled assets
