package types

import (
	"github.com/a-h/templ"
)

type Theme interface {
    Home() (templ.Component, error)
    Help() (templ.Component, error)
    Settings() (templ.Component, error)
    Search() (templ.Component, error)
    DocsRoot() (templ.Component, error)
    Docs(content string) (templ.Component, error)
    Playground() (templ.Component, error)
    Plugins() (templ.Component, error)
}

type Plugin interface {
    Name() string
    Description() string
    Help() string
    TemplResponse() (templ.Component, error)
    JsonResponse() ([]byte, error)
    Execute(params map[string]string) (interface{}, error)
}

type PluginInfo struct {
    Name        string
    Description string
    Type        string // "core" or "common"
}