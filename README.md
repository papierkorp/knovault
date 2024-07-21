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
go get github.com/labstack/echo/v4
vi server.go
go mod tidy
go run server.go

# frontend vorbereiten
go install github.com/a-h/templ/cmd/templ@latest
go get github.com/a-h/templ
templ generate

npm install -D tailwindcss
npx tailwindcss init
vi tailwind.config.js
vi static/css/main.css
npx tailwindcss -i ./static/css/main.css -o ./static/css/output.css --watch

cd static && wget https://unpkg.com/htmx.org@1.9.12/dist/htmx.min.js # --minify


go get github.com/gomarkdown/markdown
go get github.com/labstack/echo/v4/middleware@v4.12.0
```

# Usage

```bash
make dev
```