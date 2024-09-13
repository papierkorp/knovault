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
	Playground(content string) (templ.Component, error)
}