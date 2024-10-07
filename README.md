# pewitima

# Installation process

```bash
# go projekt vorbereiten
mkdir pewitima && cd pewitima
go mod init pewitima
go mod tidy

go install github.com/air-verse/air@latest
air init
vi .air.toml

# web framework installieren (echo vs gin..)
go get github.com/labstack/echo/v4@v4.12.0
go get github.com/labstack/echo/v4/middleware@v4.12.0
vi server.go
go mod tidy
go run server.go

# frontend vorbereiten
go install github.com/a-h/templ/cmd/templ@latest
go get github.com/a-h/templ@v0.2.747
templ generate

npm install -D tailwindcss
npx tailwindcss init
vi tailwind.config.js
vi static/css/main.css
npx tailwindcss -i ./static/css/main.css -o ./static/css/output.css --watch

cd static && wget https://unpkg.com/htmx.org@1.9.12/dist/htmx.min.js # --minify

```

# Usage

```bash
make dev
```

# markdown to html

```bash

go get github.com/gomarkdown/markdown
```

routes

```go
func handleHome(c echo.Context) error {
    content, err := parser.ReadMarkdownFile("example_markdown.md")
    if err != nil {
        log.Printf("Error reading markdown file: %v", err)
        return err
    }

    err = templates.Home(content).Render(c.Request().Context(), c.Response().Writer)
    if err != nil {
        log.Printf("Error rendering template: %v", err)
        return err
    }

    return nil
}
```

parser

```go
package parser

import (
    "io/ioutil"
    "os"
    "path/filepath"

    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/parser"
)

func ReadMarkdownFile(filename string) (string, error) {
    // Get the current working directory
    cwd, err := os.Getwd()
    fmt.println("cwd: ", cwd, "err: ", err)
    if err != nil {
        return "", err
    }

    // Construct the path to the data directory
    dataDir := filepath.Join(cwd, "data")

    // Construct the full path to the file
    path := filepath.Join(dataDir, filename)

    content, err := ioutil.ReadFile(path)
    if err != nil {
        return "", err
    }

    extensions := parser.CommonExtensions | parser.AutoHeadingIDs
    parser := parser.NewWithExtensions(extensions)
    md := markdown.ToHTML(content, parser, nil)

    return string(md), nil
}
```

template

```go
package templates

import (
    "pewitima/internal/templates/layout"
)

templ Home(content string) {
    @layout.Base("Home") {
        <h1 class="mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-white">HOME</h1>
        <div class="markdown-content">
            @templ.Raw(content)
        </div>
    }
}
```

# Creating Core Plugins and Themes

This guide explains how to create new core plugins and themes for the project.

## Creating a New Core Plugin

To create a new core plugin, follow these steps:

1. Create a new Go file in the `internal/plugins/core` directory with a descriptive name for your plugin (e.g., `MyNewPlugin.go`).

2. In this file, define a struct for your plugin and implement the `Plugin` interface. Here's a template:

```go
package core

import (
    "encoding/json"
    "pewitima/internal/plugins"
    "github.com/a-h/templ"
)

type MyNewPlugin struct{}

func (p *MyNewPlugin) Name() string {
    return "MyNewPlugin"
}

func (p *MyNewPlugin) Description() string {
    return "Description of what your plugin does"
}

func (p *MyNewPlugin) Help() string {
    return "Instructions on how to use your plugin"
}

func (p *MyNewPlugin) TemplResponse() (templ.Component, error) {
    // Implement this method if your plugin returns a templ Component
    return templ.Raw("Your templ component here"), nil
}

func (p *MyNewPlugin) JsonResponse() ([]byte, error) {
    // Implement this method if your plugin returns JSON
    response := map[string]string{
        "key": "value",
    }
    return json.Marshal(response)
}

func (p *MyNewPlugin) Execute(params map[string]string) (interface{}, error) {
    // Implement the main functionality of your plugin here
    // You can return a templ.Component, JSON, or any other type
    return "Plugin execution result", nil
}

func init() {
    plugins.RegisterCorePlugin("MyNewPlugin", &MyNewPlugin{})
}
```

3. Implement the methods of the `Plugin` interface according to your plugin's functionality.

4. In the `init()` function, register your plugin using `plugins.RegisterCorePlugin()`.

5. If your plugin requires any additional dependencies or setup, make sure to include them in the file.

6. Update any relevant UI components (e.g., `playground.templ`) to include your new plugin if necessary.

## Creating a New Theme

To create a new theme, follow these steps:

1. Create a new directory under `internal/themes` for your theme (e.g., `internal/themes/myNewTheme`).

2. Inside this directory, create a Go file named after your theme (e.g., `myNewTheme.go`). Use this template:

```go
package myNewTheme

import (
    "github.com/a-h/templ"
    "pewitima/internal/themes"
    "pewitima/internal/themes/myNewTheme/templates"
)

type MyNewTheme struct{}

func (t *MyNewTheme) Home() (templ.Component, error) {
    return templates.Home(), nil
}

func (t *MyNewTheme) Help() (templ.Component, error) {
    return templates.Help(), nil
}

func (t *MyNewTheme) Settings() (templ.Component, error) {
    return templates.Settings(), nil
}

func (t *MyNewTheme) Search() (templ.Component, error) {
    return templates.Search(), nil
}

func (t *MyNewTheme) DocsRoot() (templ.Component, error) {
    return templates.DocsRoot(), nil
}

func (t *MyNewTheme) Docs(content string) (templ.Component, error) {
    return templates.Docs(content), nil
}

func (t *MyNewTheme) Playground() (templ.Component, error) {
    return templates.Playground(), nil
}

func (t *MyNewTheme) Plugins() (templ.Component, error) {
    return templates.Plugins(), nil
}

func init() {
    themes.RegisterTheme("myNewTheme", &MyNewTheme{})
}
```

3. Create a `templates` subdirectory in your theme directory.

4. In the `templates` directory, create templ files for each component (e.g., `home.templ`, `help.templ`, etc.). Implement your theme's design in these files.

5. Create a `layout` subdirectory inside `templates` and add a `base.templ` file for the base layout of your theme.

6. Implement each method in your theme struct to return the appropriate templ component.

7. In the `init()` function, register your theme using `themes.RegisterTheme()`.

8. Update the `server.go` file to import your new theme:

```go
import (
    // ... other imports ...
    _ "pewitima/internal/themes/myNewTheme"
)
```

9. To use your new theme, you can set it as the current theme in `server.go` or use the ThemeChanger plugin:

```go
err := themes.SetCurrentTheme("myNewTheme")
if err != nil {
    log.Fatalf("Failed to set new theme: %v", err)
}
```

Remember to style your theme components using Tailwind CSS classes to maintain consistency with the project's styling approach.

By following these steps, you can extend the functionality of the project with new plugins and create custom themes to change the appearance of the application.
