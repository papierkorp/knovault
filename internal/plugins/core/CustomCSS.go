package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"knovault/internal/plugins"
	pluginTemplates "knovault/internal/plugins/templates"
	"knovault/internal/types"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CustomCSSPlugin struct {
    CSSFilePath string
}

func (p *CustomCSSPlugin) Name() string {
    return "CustomCSS"
}

func (p *CustomCSSPlugin) Description() string {
    return "Manages a custom CSS file for user-defined styles"
}

func (p *CustomCSSPlugin) Help() string {
    return "Use this plugin to view, edit, and apply custom CSS styles. Your styles will override Tailwind styles."
}

func (p *CustomCSSPlugin) TemplResponse() (templ.Component, error) {
    css, err := p.readCustomCSS()
    if err != nil {
        css = getDefaultCSS()
    }
    return pluginTemplates.CustomCSSEditor(css, ""), nil
}

func (p *CustomCSSPlugin) JsonResponse() ([]byte, error) {
    css, err := p.readCustomCSS()
    if err != nil {
        css = getDefaultCSS()
    }
    response := map[string]string{
        "css": css,
    }
    return json.Marshal(response)
}

func (p *CustomCSSPlugin) Execute(params map[string]string) (interface{}, error) {
    action, ok := params["action"]
    var css string
    var err error
    message := ""

    if ok && action == "reset" {
        css = getDefaultCSS()
        err = p.writeCustomCSS(css)
        if err == nil {
            message = "Custom CSS reset successfully!"
        }
    } else if css, ok = params["css"]; ok {
        err = p.writeCustomCSS(css)
        if err == nil {
            message = "Custom CSS updated successfully!"
        }
    } else {
        css, err = p.readCustomCSS()
        if err != nil {
            css = getDefaultCSS()
        }
    }

    if err != nil {
        return nil, err
    }

    return pluginTemplates.CustomCSSEditor(css, message), nil
}

func (p *CustomCSSPlugin) Route() types.PluginRoute {
    return types.PluginRoute{
        Method: "POST",
        Path:   "/plugins/CustomCSS",
        Handler: func(c echo.Context) error {
            params := make(map[string]string)
            if err := c.Bind(&params); err != nil {
                return err
            }

            result, err := p.Execute(params)
            if err != nil {
                return err
            }

            component, ok := result.(templ.Component)
            if !ok {
                return fmt.Errorf("unexpected result type")
            }

            // Set headers to trigger a full page reload
            c.Response().Header().Set("HX-Refresh", "true")

            return component.Render(c.Request().Context(), c.Response())
        },
    }
}


func (p *CustomCSSPlugin) readCustomCSS() (string, error) {
    data, err := os.ReadFile(p.CSSFilePath)
    if err != nil {
        if os.IsNotExist(err) {
            return "", nil
        }
        return "", err
    }
    return string(data), nil
}

func (p *CustomCSSPlugin) writeCustomCSS(css string) error {
    dir := filepath.Dir(p.CSSFilePath)
    if err := os.MkdirAll(dir, 0755); err != nil {
        return err
    }
    return os.WriteFile(p.CSSFilePath, []byte(css), 0644)
}

func getDefaultCSS() string {
    return `/* Add your custom CSS styles here. These styles will override the default styles. */
/* In some cases you need to hard reload your page to see your changes (SHIFT + F5) */

/*
.bg-slate-100 {
	background-color: yellow;
}

#customcsstextarea {
	background-color: yellow;
}

#search-input {
	background-color: blue;
}
*/
`
}

func (p *CustomCSSPlugin) ExtendTemplate(templateName string) (templ.Component, error) {
	switch templateName {
	case "settings":
		return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			css, _ := p.readCustomCSS()
			return pluginTemplates.CustomCSSEditor(css, "").Render(ctx, w)
		}), nil
	default:
		return nil, fmt.Errorf("template %s not supported by CustomCSS plugin", templateName)
	}
}

func init() {
    cssFilePath := filepath.Join("static", "css", "custom.css")
    plugins.RegisterCorePlugin("CustomCSS", &CustomCSSPlugin{CSSFilePath: cssFilePath})
}