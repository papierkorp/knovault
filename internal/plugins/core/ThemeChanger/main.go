// internal/plugins/core/ThemeChanger/main.go

package main

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "strings"

    "github.com/a-h/templ"
    "github.com/labstack/echo/v4"
    "knovault/internal/themes"
    "knovault/internal/types"
)

type ThemeChangerPlugin struct{}

type themeChangerTemplates struct{}

func (t themeChangerTemplates) ThemeChangerForm(availableThemes []string, currentTheme string) templ.Component {
    return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
        _, err = io.WriteString(w, `<div>
            <h2>Theme Settings</h2>
            <form hx-post="/plugins/ThemeChanger" hx-swap="outerHTML">
                <select name="theme">`)

        for _, theme := range availableThemes {
            selected := ""
            if theme == currentTheme {
                selected = ` selected="selected"`
            }
            _, err = fmt.Fprintf(w, `<option value="%s"%s>%s</option>`,
                theme, selected, strings.Title(theme))
        }

        _, err = io.WriteString(w, `</select>
                <button type="submit">Change Theme</button>
            </form>
        </div>`)
        return err
    })
}

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
    templates := themeChangerTemplates{}
    availableThemes := themes.GetAvailableThemes()
    currentTheme := themes.GetCurrentThemeName()
    return templates.ThemeChangerForm(availableThemes, currentTheme), nil
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
    templates := themeChangerTemplates{}
    switch templateName {
    case "settings":
        availableThemes := themes.GetAvailableThemes()
        currentTheme := themes.GetCurrentThemeName()
        return templates.ThemeChangerForm(availableThemes, currentTheme), nil
    default:
        return nil, fmt.Errorf("template %s not supported by ThemeChanger plugin", templateName)
    }
}

// Plugin is the exported symbol that will be loaded by the plugin system
var Plugin ThemeChangerPlugin