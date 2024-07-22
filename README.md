# gowiki

# Installation process

```bash
# go projekt vorbereiten
mkdir gowiki && cd gowiki
go mod init gowiki
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
    "gowiki/internal/templates/layout"
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