package types

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Theme interface {
	Home() (templ.Component, error)
	Help() (templ.Component, error)
	Settings() (templ.Component, error)
	Search() (templ.Component, error)
	DocsRoot() (templ.Component, error)
	Docs(content string) (templ.Component, error)
	Playground(content string) (templ.Component, error)
}

type Plugin interface {
    Name() string
    Initialize() error
    Shutdown() error
    Handlers() map[string]echo.HandlerFunc
    TemplateData() map[string]interface{}
    Render() templ.Component
}