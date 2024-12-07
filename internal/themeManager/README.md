# Theme Manager

The Theme Manager handles theme management for Knovault. This directory contains all theme implementations and manages both builtin and external themes.

## Directory Structure

```
themeManager/
├── builtin/              # Builtin theme implementations
│   └── defaulttheme/
├── external/             # External theme implementations
│   └── dark/
├── themeManager.go          # Core theme manager
├── themeManagerBuiltinHelper.go
└── themeManagerExternalHelper.go
```

## Theme Types

### Builtin Themes

- Source code is directly in the builtin directory
- Always available and loaded by default
- DefaultTheme is the core theme

Available builtin themes:

- DefaultTheme: Default application appearance

### External Themes

- Located in the external directory
- Can be compiled to .so files
- Can be added/removed without changing core code
- Must follow theme interface requirements

Available external themes:

- Dark: Dark mode theme

## Creating a New Theme

### Builtin Theme

1. Create a new directory under `builtin/`:

   ```
   builtin/
   └── yourtheme/
       ├── yourtheme.go    # Theme implementation
       └── templates/      # Theme templates
           ├── layout/
           └── pages/
   ```

2. Implement the Theme interface:

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

### External Theme

1. Create a new directory under `external/`:

   ```
   external/
   └── YourTheme/
       ├── main.go           # Theme entry point
       ├── plugin/
       │   └── plugin.go     # Theme implementation
       └── templates/        # Theme templates
           ├── layout/
           └── pages/
   ```

2. Create main.go:

   ```go
   package main

   import (
       "knovault/internal/types"
       "knovault/internal/themeManager/external/YourTheme/plugin"
   )

   var Theme types.Theme = &plugin.YourTheme{}

   func main() {}
   ```

3. Create template files:
   - base.templ: Base layout template
   - Required page templates (home, help, settings, etc.)

## Template Structure

Each theme must provide these templates:

- Layout:
  - base.templ: Base layout with common elements
- Pages:
  - docs.templ: Documentation page
  - docsroot.templ: Documentation root
  - help.templ: Help page
  - home.templ: Home page
  - playground.templ: Plugin playground
  - plugins.templ: Plugin listing
  - search.templ: Search interface
  - settings.templ: Settings page

## Building External Themes

External themes are automatically compiled to .so files when detected:

1. The theme manager scans the external directory
2. Finds main.go files
3. Compiles them to .so files
4. Loads the resulting themes

You can also manually compile themes:

```bash
go build -buildmode=plugin -o yourtheme.so main.go
```

## Development Tips

- Start by copying an existing theme
- Test designs with the builtin theme first
- Keep styles consistent
- Use CSS variables for theming
- Test all templates thoroughly
- Support mobile views
- Consider accessibility
- Document your theme's requirements
- Follow existing theme patterns
