package secondTheme

import (
	"github.com/a-h/templ"
	"pewitima/internal/themes"
	"pewitima/internal/themes/secondTheme/templates"
)

type SecondTheme struct{}


func (t *SecondTheme) Home() (templ.Component, error) {
	return templates.Home(), nil
}

func (t *SecondTheme) Help() (templ.Component, error) {
	return templates.Help(), nil
}

func (t *SecondTheme) Settings() (templ.Component, error) {
	return templates.Settings(), nil
}

func (t *SecondTheme) Search() (templ.Component, error) {
	return templates.Search(), nil
}

func (t *SecondTheme) DocsRoot() (templ.Component, error) {
	return templates.DocsRoot(), nil
}

func (t *SecondTheme) Docs(content string) (templ.Component, error) {
	return templates.Docs(content), nil
}

func (t *SecondTheme) Playground() (templ.Component, error) {
	return templates.Playground(), nil
}

func (t *SecondTheme) Plugins() (templ.Component, error) {
	return templates.Plugins(), nil
}


func init() {
	themes.RegisterTheme("secondTheme", &SecondTheme{})
}
