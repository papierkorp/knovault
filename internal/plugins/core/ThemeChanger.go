package core

import (
	"encoding/json"
	"fmt"
	"knovault/internal/plugins"
	"knovault/internal/themes"
	"knovault/internal/types"
	pluginTemplates "knovault/internal/plugins/templates"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
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
	return pluginTemplates.ThemeChangerForm(availableThemes, currentTheme), nil
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

func (p *ThemeChangerPlugin) Route() types.PluginRoute {
	return types.PluginRoute{
		Method: "POST",
		Path:   "/plugins/ThemeChanger",
		Handler: func(c echo.Context) error {
			newTheme := c.FormValue("theme")
			if err := themes.SetCurrentTheme(newTheme); err != nil {
				return c.JSON(echo.ErrInternalServerError.Code, map[string]string{"error": err.Error()})
			}
			c.Response().Header().Set("HX-Refresh", "true")
			return c.NoContent(200)
		},
	}
}

func (p *ThemeChangerPlugin) ExtendTemplate(templateName string) (templ.Component, error) {
	switch templateName {
	case "settings":
		availableThemes := themes.GetAvailableThemes()
		currentTheme := themes.GetCurrentThemeName()
		return pluginTemplates.ThemeChangerForm(availableThemes, currentTheme), nil
	default:
		return nil, fmt.Errorf("template %s not supported by ThemeChanger plugin", templateName)
	}
}

func init() {
	plugins.RegisterCorePlugin("ThemeChanger", &ThemeChangerPlugin{})
}