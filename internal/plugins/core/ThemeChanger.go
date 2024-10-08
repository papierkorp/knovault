package core

import (
	"encoding/json"
	"fmt"
	"io"
	"context"
	"pewito/internal/plugins"
	"pewito/internal/themes"

	"github.com/a-h/templ"
)
type ThemeChangerPlugin struct{}

func (p *ThemeChangerPlugin) Name() string {
	return "ThemeChanger"
}

func (p *ThemeChangerPlugin) Description() string {
	return "Lists available themes and allows changing the current theme"
}

func (p *ThemeChangerPlugin) Help() string {
	return "Use this plugin to view available themes and change the current theme."
}

func (p *ThemeChangerPlugin) TemplResponse() (templ.Component, error) {
	availableThemes := themes.GetAvailableThemes()
	currentTheme := themes.GetCurrentThemeName()

	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = fmt.Fprintf(w, `
			<div>
				<h3 class="text-lg font-semibold mb-2">Available Themes</h3>
				<ul class="list-disc pl-5 mb-4">
					%s
				</ul>
				<h4 class="text-md font-semibold mb-2">Current Theme: %s</h4>
				<form hx-post="/plugins/ThemeChanger" hx-swap="outerHTML">
					<select name="theme" class="border rounded p-1 mr-2">
						%s
					</select>
					<button type="submit" class="bg-blue-500 text-white px-4 py-1 rounded">Change Theme</button>
				</form>
			</div>
		`, p.generateThemeList(availableThemes), currentTheme, p.generateThemeOptions(availableThemes, currentTheme))
		return
	}), nil
}

func (p *ThemeChangerPlugin) JsonResponse() ([]byte, error) {
	availableThemes := themes.GetAvailableThemes()
	currentTheme := themes.GetCurrentThemeName()

	response := map[string]interface{}{
		"availableThemes": availableThemes,
		"currentTheme":    currentTheme,
	}
	return json.Marshal(response)
}

func (p *ThemeChangerPlugin) Execute(params map[string]string) (interface{}, error) {
	return templ.Raw(""), nil
}

func (p *ThemeChangerPlugin) generateThemeList(themes []string) string {
	var list string
	for _, theme := range themes {
		list += fmt.Sprintf("<li>%s</li>", theme)
	}
	return list
}

func (p *ThemeChangerPlugin) generateThemeOptions(themes []string, currentTheme string) string {
	var options string
	for _, theme := range themes {
		selected := ""
		if theme == currentTheme {
			selected = "selected"
		}
		options += fmt.Sprintf(`<option value="%s" %s>%s</option>`, theme, selected, theme)
	}
	return options
}

func init() {
	plugins.RegisterCorePlugin("ThemeChanger", &ThemeChangerPlugin{})
}