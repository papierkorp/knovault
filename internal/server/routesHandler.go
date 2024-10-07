package server

import (
	"net/http"
	"pewitima/internal/plugins"
	"pewitima/internal/themes"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)


func handleHome(c echo.Context) error {
	component, err := themes.GetCurrentTheme().Home()
	if err != nil {
		return err
	}
	return _render(c, component)
}


func handlePlayground(c echo.Context) error {
    component, err := themes.GetCurrentTheme().Playground()
    if err != nil {
        return err
    }
    return _render(c, component)
}

func handleHelp(c echo.Context) error {
	component, err := themes.GetCurrentTheme().Help()
	if err != nil {
		return err
	}
	return _render(c, component)
}

func handleSettings(c echo.Context) error {
	component, err := themes.GetCurrentTheme().Settings()
	if err != nil {
		return err
	}
	return _render(c, component)
}

func handleSearch(c echo.Context) error {
	component, err := themes.GetCurrentTheme().Search()
	if err != nil {
		return err
	}
	return _render(c, component)
}

func handleDocsRoot(c echo.Context) error {
	component, err := themes.GetCurrentTheme().DocsRoot()
	if err != nil {
		return err
	}
	return _render(c, component)
}

func handleDocs(c echo.Context) error {
	title := c.Param("title")
	component, err := themes.GetCurrentTheme().Docs(title)
	if err != nil {
		return err
	}
	return _render(c, component)
}


func handlePlugins(c echo.Context) error {
	component, err := themes.GetCurrentTheme().Plugins()
	if err != nil {
		return err
	}
	return _render(c, component)
}

func handlePluginExecute(c echo.Context) error {
    pluginName := c.Param("pluginName")
    plugin, ok := plugins.GetPlugin(pluginName)
    if !ok {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Plugin not found"})
    }

    // Collect all form values as parameters
    params := make(map[string]string)
    formParams, err := c.FormParams()
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse form parameters"})
    }
    for key, values := range formParams {
        if len(values) > 0 {
            params[key] = values[0]
        }
    }

    // Special handling for ThemeChanger plugin
    if c.Request().Method == "POST" && pluginName == "ThemeChanger" {
        newTheme, ok := params["theme"]
        if ok {
            if err := themes.SetCurrentTheme(newTheme); err != nil {
                return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
            }
            c.Response().Header().Set("HX-Refresh", "true")
        }
    }

    // Execute the plugin
    response, err := plugin.Execute(params)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    // Handle different response types
    switch response := response.(type) {
    case []byte:
        return c.Blob(http.StatusOK, "application/json", response)
    case templ.Component:
        return _render(c, response)
    default:
        return c.JSON(http.StatusOK, response)
    }
}