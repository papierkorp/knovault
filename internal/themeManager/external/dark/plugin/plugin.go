// internal/themeManager/external/dark/plugin/plugin.go
package plugin

import (
    "github.com/a-h/templ"
    "knovault/internal/themeManager/external/dark/templates"
)

type DarkTheme struct{}

func (t *DarkTheme) Home() (templ.Component, error) {
    return templates.Home(), nil
}

func (t *DarkTheme) Help() (templ.Component, error) {
    return templates.Help(), nil
}

func (t *DarkTheme) Settings() (templ.Component, error) {
    return templates.Settings(), nil
}

func (t *DarkTheme) Search() (templ.Component, error) {
    return templates.Search(), nil
}

func (t *DarkTheme) DocsRoot() (templ.Component, error) {
    return templates.DocsRoot(), nil
}

func (t *DarkTheme) Docs(content string) (templ.Component, error) {
    return templates.Docs(content), nil
}

func (t *DarkTheme) Playground() (templ.Component, error) {
    return templates.Playground(), nil
}

func (t *DarkTheme) Plugins() (templ.Component, error) {
    return templates.Plugins(), nil
}

