package plugin

import (
    "github.com/a-h/templ"
    "knovault/internal/assetManager/themes/dark/templates"
)

type Dark struct{}

func (t *Dark) Home() (templ.Component, error) {
    return templates.Home(), nil
}

func (t *Dark) Help() (templ.Component, error) {
    return templates.Help(), nil
}

func (t *Dark) Settings() (templ.Component, error) {
    return templates.Settings(), nil
}

func (t *Dark) Search() (templ.Component, error) {
    return templates.Search(), nil
}

func (t *Dark) DocsRoot() (templ.Component, error) {
    return templates.DocsRoot(), nil
}

func (t *Dark) Docs(content string) (templ.Component, error) {
    return templates.Docs(content), nil
}

func (t *Dark) Playground() (templ.Component, error) {
    return templates.Playground(), nil
}

func (t *Dark) Plugins() (templ.Component, error) {
    return templates.Plugins(), nil
}

// Initialize Theme as pointer
var Theme Dark