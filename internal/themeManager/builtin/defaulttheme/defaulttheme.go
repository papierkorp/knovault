// internal/themeManager/builtin/defaulttheme/defaulttheme.go
package defaulttheme

import (
    "github.com/a-h/templ"
    "knovault/internal/themeManager/builtin/defaulttheme/templates"
)

type DefaultTheme struct{}

func (t *DefaultTheme) Home() (templ.Component, error) {
    return templates.Home(), nil
}

func (t *DefaultTheme) Help() (templ.Component, error) {
    return templates.Help(), nil
}

func (t *DefaultTheme) Settings() (templ.Component, error) {
    return templates.Settings(), nil
}

func (t *DefaultTheme) Search() (templ.Component, error) {
    return templates.Search(), nil
}

func (t *DefaultTheme) DocsRoot() (templ.Component, error) {
    return templates.DocsRoot(), nil
}

func (t *DefaultTheme) Docs(content string) (templ.Component, error) {
    return templates.Docs(content), nil
}

func (t *DefaultTheme) Playground() (templ.Component, error) {
    return templates.Playground(), nil
}

func (t *DefaultTheme) Plugins() (templ.Component, error) {
    return templates.Plugins(), nil
}