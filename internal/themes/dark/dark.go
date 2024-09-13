package dark

import (
	"github.com/a-h/templ"
    "gowiki/internal/themes"
    "gowiki/internal/themes/dark/templates"
)

type dark struct{}

func (t *dark) Docs(content string) (templ.Component, error) {
	return templates.Docs(content), nil
}

func (t *dark) DocsRoot() (templ.Component, error) {
	return templates.DocsRoot(), nil
}

func (t *dark) Playground(content string) (templ.Component, error) {
	return templates.Playground(content), nil
}

func (t *dark) Search() (templ.Component, error) {
	return templates.Search(), nil
}

func (t *dark) Settings() (templ.Component, error) {
	return templates.Settings(), nil
}

func (t *dark) Home() (templ.Component, error) {
	return templates.Home(), nil
}

func (t *dark) Help() (templ.Component, error) {
	return templates.Help(), nil
}

func init() {
	themes.RegisterTheme("dark", &dark{})
}
