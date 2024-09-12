package dark

import (
	"gowiki/internal/templates"
	"gowiki/internal/themes"
)

type DarkTheme struct{}

func (t *DarkTheme) Home() (templ.Component, error) {
	return templates.DarkHome(), nil
}

func (t *DarkTheme) Help() (templ.Component, error) {
	return templates.DarkHelp(), nil
}

func init() {
	themes.RegisterTheme("dark", &DarkTheme{})
}