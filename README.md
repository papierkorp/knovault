# Pewito: Personal Wiki and To-dos

Pewito is a personal wiki and to-do application built with Go, HTMX, Tailwind CSS, and Templ. It uses a `data` folder for storing markdown files and settings, which can be version-controlled with Git.

it should just have one topbar which includes the logo, a search bar, a settings icon, a select project button and a latest changes button

as the main content (everything below the topbar) i want to display the markdown documents as cards where each file has this metadata:

- Title: title
- Description: first 100 words from the content
- Type: doc/task/filter
- Categorie: devops/cooking/develop...
- Tags: kubernetes/pizza/...
- content: md content..
- Project: can be read from the path
- path: /data/project1/tasks/task1.md

while the front of the card should just include title, type and the first 3 tags + a button for more details

the color palette is:

- primary: #81a15a
- accent: #9ece5a
- neutral: #6b6e72
- black: #4e5154
- white: #d1d1d3


# Installation process

```bash
# go projekt vorbereiten
mkdir pewito && cd pewito
go mod init pewito
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

# Setup

## Setting Up the Development Environment

Follow these steps to set up your development environment:

1. Install Go (1.16 or later): https://golang.org/doc/install

2. Install Node.js and npm: https://nodejs.org/

3. Clone the repository:
4. `git clone https://github.com/your-username/pewito.git`

5. Install Go dependencies: `go mod tidy`

6. Install Templ: `go install github.com/a-h/templ/cmd/templ@latest`

7. Install Tailwind CSS: `npm install`

8. Install Air for live reloading (optional): `go install github.com/cosmtrek/air@latest`

## Running the Application

The project includes a Makefile to simplify common development tasks. Here are the main commands you'll use:

1. Generate Templ files and build Tailwind CSS: `make build`

2. Run the application in development mode (with live reloading): `make dev`

3. Build Tailwind CSS only: `make tailwind-build`

4. Watch for Tailwind CSS changes: `make tailwind-watch`

5. Generate Templ files only: `make templ-generate`

To start developing, run:

```bash
make dev
```

This command will generate Templ files, build Tailwind CSS, and start the application with live reloading.

Open your browser and navigate to [http://localhost:132](http://localhost:1323) to see the application.

## Creating a New Theme

To create a new theme:

1. Create a new directory under `internal/themes` for your theme (e.g., `internal/themes/myNewTheme`).

2. Create a Go file for your theme (e.g., `myNewTheme.go`) with the following structure:

```go
package myNewTheme

import (
    "github.com/a-h/templ"
    "pewito/internal/themes"
    "pewito/internal/themes/myNewTheme/templates"
)

type MyNewTheme struct{}

func (t *MyNewTheme) Home() (templ.Component, error) {
    return templates.Home(), nil
}

// Implement other methods (Help, Settings, Search, DocsRoot, Docs, Playground, Plugins)

func init() {
    themes.RegisterTheme("myNewTheme", &MyNewTheme{})
}
```

3. Create a `templates` directory inside your theme directory and add Templ files for each component (e.g., `home.templ`, `help.templ`, etc.).

4. Implement the theme's design in these Templ files using Tailwind CSS classes.

5. Update `server.go` to import your new theme:

```go
import (
    // ... other imports ...
    _ "pewito/internal/themes/myNewTheme"
)
```

6. To use the new theme, set it as the current theme in `server.go`:

```go
err := themes.SetCurrentTheme("myNewTheme")
if err != nil {
 log.Fatalf("Failed to set new theme: %v", err)
}
```

7. After creating or modifying Templ files, run: `make templ-generate`

8. If you've made changes to Tailwind classes, rebuild the CSS: `make tailwind-build`

## Creating a New Plugin

To create a new plugin:

1. Create a new Go file in the `internal/plugins/core` directory (e.g., `MyNewPlugin.go`).

2. Implement the `Plugin` interface:

```go
package core

import (
    "encoding/json"
    "pewito/internal/plugins"
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
    return "Plugin execution result", nil
}

func init() {
    plugins.RegisterCorePlugin("MyNewPlugin", &MyNewPlugin{})
}
```

3. If your plugin requires a custom route, implement the `PluginWithRoute` interface:

```go
func (p *MyNewPlugin) Route() types.PluginRoute {
    return types.PluginRoute{
        Method: "POST",
        Path:   "/plugins/MyNewPlugin",
        Handler: func(c echo.Context) error {
            // Implement your route handler here
        },
    }
}
```

4. If your plugin extends templates, implement the `PluginWithTemplateExtensions` interface:

```go
func (p *MyNewPlugin) ExtendTemplate(templateName string) (templ.Component, error) {
    switch templateName {
    case "settings":
        return templ.Raw("<div>My plugin extension</div>"), nil
    default:
        return nil, fmt.Errorf("template %s not supported by MyNewPlugin", templateName)
    }
}
```

5. Update any relevant UI components (e.g., `playground.templ` or `settings.templ`) to include your new plugin if necessary.

6. After creating or modifying a plugin, rebuild the application: `make build`

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
