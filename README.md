# knovault: Knowledgevault - Personal Wiki and To-dos

Personal wiki and to-do application built with Go, HTMX, and Templ.

## Design Choices

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

## Core Features

- Markdown file storage in `data` folder with Git version control
- Plugin system for extensibility
- Theme system for customizable UI
- HTMX integration for dynamic content

## Color palette

- primary: #81a15a
- accent: #9ece5a
- neutral: #6b6e72
- black: #4e5154
- white: #d1d1d3

# Installation process

```bash
# go projekt vorbereiten
mkdir knovault && cd knovault
go mod init knovault
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

vi static/css/main.css

cd static && wget https://unpkg.com/htmx.org@1.9.12/dist/htmx.min.js # --minify

```

# Setup

## Prerequisites

1. Go 1.16+
2. `go install github.com/a-h/templ/cmd/templ@latest`
3. `go install github.com/cosmtrek/air@latest`

## Development

```bash
make dev  # Starts development server with live reload
```

## Project Structure

```
.
├── internal/
│   ├── plugins/
│   │   ├── core/       # Core plugins with source code
│   │   │   └── PluginName/
│   │   │       ├── main.go
│   │   │       ├── PluginName.so
│   │   │       └── templates/
│   │   └── common/    # Runtime-loadable plugins (.so only)
│   └── themes/
│       ├── core/      # Core themes with source code
│       │   └── ThemeName/
│       │       ├── main.go
│       │       ├── ThemeName.so
│       │       └── templates/
│       └── common/    # Runtime-loadable themes (.so only)
```

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

## Creating Plugins

1. Create plugin directory structure:

```bash
mkdir -p internal/plugins/core/MyPlugin/templates
```

2. Implement the plugin interface in `main.go`:

```go
package main

type MyPlugin struct{}

func (p *MyPlugin) Name() string {
    return "MyPlugin"
}

// Implement other required methods...

var Plugin = MyPlugin{}
```

3. Add templates in the templates directory if needed
4. Plugin will be compiled automatically on server start

## Creating Themes

1. Create theme directory structure:

```bash
mkdir -p internal/themes/core/MyTheme/templates/layout
```

2. Implement the theme interface in `main.go`:

```go
package main

type MyTheme struct{}

func (t *MyTheme) Home() (templ.Component, error) {
    return templates.Home(), nil
}

// Implement other required methods...

var Theme = MyTheme{}
```

3. Add templates and layout files
4. Theme will be compiled automatically on server start

## Runtime Plugin/Theme Installation

- Copy compiled .so files to respective common directories
- System will detect and load automatically
- No server restart required

## Available Make Commands

- `make dev`: Development mode with live reload
- `make build`: Production build
- `make docker-build-dev`: Build development Docker image
- `make docker-build-prod`: Build production Docker image
-

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
