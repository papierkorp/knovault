// internal/pluginManager/builtin/themechanger/themechanger.go
package themechanger

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "strings"

    "github.com/a-h/templ"
    "github.com/labstack/echo/v4"
    "knovault/internal/types"
    "knovault/internal/globals"
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
    tm := globals.GetThemeManager()
    availableThemes := tm.GetAvailableThemes()
    currentTheme := tm.GetCurrentThemeName()
    return p.renderThemeChangerForm(availableThemes, currentTheme), nil
}

func (p *ThemeChangerPlugin) JsonResponse() ([]byte, error) {
    tm := globals.GetThemeManager()
    availableThemes := tm.GetAvailableThemes()
    currentTheme := tm.GetCurrentThemeName()

    response := map[string]interface{}{
        "availableThemes": availableThemes,
        "currentTheme":    currentTheme,
    }
    return json.Marshal(response)
}

func (p *ThemeChangerPlugin) Execute(params map[string]string) (interface{}, error) {
    return nil, nil
}

func (p *ThemeChangerPlugin) Route() types.PluginRoute {
    return types.PluginRoute{
        Method: "POST",
        Path:   "/plugins/ThemeChanger",
        Handler: func(c echo.Context) error {
            tm := globals.GetThemeManager()
            newTheme := c.FormValue("theme")
            if err := tm.SetCurrentTheme(newTheme); err != nil {
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
        tm := globals.GetThemeManager()
        availableThemes := tm.GetAvailableThemes()
        currentTheme := tm.GetCurrentThemeName()
        return p.renderThemeChangerForm(availableThemes, currentTheme), nil
    default:
        return nil, fmt.Errorf("template %s not supported by ThemeChanger plugin", templateName)
    }
}

func (p *ThemeChangerPlugin) renderThemeChangerForm(availableThemes []string, currentTheme string) templ.Component {
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
            if err != nil {
                return err
            }
        }

        _, err = io.WriteString(w, `</select>
                <button type="submit">Change Theme</button>
            </form>
        </div>`)
        return err
    })
}