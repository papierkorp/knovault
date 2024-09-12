package defaultTheme

import (
	"github.com/a-h/templ"
	"gowiki/internal/templates"
	"gowiki/internal/themes"
)

type defaultTheme struct{}

func (t *defaultTheme) Docs(content string) (templ.Component, error) {
	return templates.Docs(content), nil
}

func (t *defaultTheme) DocsRoot() (templ.Component, error) {
	return templates.DocsRoot(), nil
}

func (t *defaultTheme) Playground(content string) (templ.Component, error) {
	return templates.Playground(content), nil
}

func (t *defaultTheme) Search() (templ.Component, error) {
	return templates.Search(), nil
}

func (t *defaultTheme) Settings() (templ.Component, error) {
	return templates.Settings(), nil
}

func (t *defaultTheme) Home() (templ.Component, error) {
	return templates.Home(), nil
}

func (t *defaultTheme) Help() (templ.Component, error) {
	return templates.Help(), nil
}

func init() {
	themes.RegisterTheme("defaultTheme", &defaultTheme{})
}
