// internal/pluginManager/builtin/cssswitcher/cssswitcher.go
package cssswitcher

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "github.com/a-h/templ"
    "github.com/labstack/echo/v4"
    "knovault/internal/pluginManager/builtin/cssswitcher/templates"
    "knovault/internal/globals"
    "knovault/internal/types"
)

type CSSSwitcherPlugin struct{}

func init() {
    globals.RegisterPlugin("CSSSwitcher", NewCSSSwitcherPlugin)
}

func NewCSSSwitcherPlugin() types.Plugin {
    return &CSSSwitcherPlugin{}
}

func (p *CSSSwitcherPlugin) Name() string {
    return "CSSSwitcher"
}

func (p *CSSSwitcherPlugin) Description() string {
    return "Switch between different visual themes using CSS"
}

func (p *CSSSwitcherPlugin) Help() string {
    return "Select from available visual themes to change the application's appearance. Changes take effect immediately."
}

func (p *CSSSwitcherPlugin) TemplResponse() (templ.Component, error) {
    styles, err := p.getAvailableStyles()
    if err != nil {
        return nil, err
    }
    currentStyle := p.getCurrentStyle()
    return templates.StyleSelector(styles, currentStyle), nil
}

func (p *CSSSwitcherPlugin) JsonResponse() ([]byte, error) {
    styles, err := p.getAvailableStyles()
    if err != nil {
        return nil, err
    }
    currentStyle := p.getCurrentStyle()

    response := map[string]interface{}{
        "availableStyles": styles,
        "currentStyle": currentStyle,
    }
    return json.Marshal(response)
}

func (p *CSSSwitcherPlugin) Execute(params map[string]string) (interface{}, error) {
    if style, ok := params["style"]; ok {
        if err := p.setStyle(style); err != nil {
            return nil, fmt.Errorf("failed to set style: %v", err)
        }
    }

    styles, err := p.getAvailableStyles()
    if err != nil {
        return nil, err
    }
    currentStyle := p.getCurrentStyle()
    return templates.StyleSelector(styles, currentStyle), nil
}

func (p *CSSSwitcherPlugin) Route() types.PluginRoute {
    return types.PluginRoute{
        Method: "POST",
        Path:   "/plugins/CSSSwitcher",
        Handler: func(c echo.Context) error {
            style := c.FormValue("style")
            if err := p.setStyle(style); err != nil {
                return c.JSON(400, map[string]string{"error": err.Error()})
            }
            c.Response().Header().Set("HX-Refresh", "true")
            return c.NoContent(200)
        },
    }
}

func (p *CSSSwitcherPlugin) ExtendTemplate(templateName string) (templ.Component, error) {
    if templateName != "settings" {
        return nil, fmt.Errorf("template %s not supported by CSSSwitcher plugin", templateName)
    }

    styles, err := p.getAvailableStyles()
    if err != nil {
        return nil, err
    }
    currentStyle := p.getCurrentStyle()
    return templates.StyleSelector(styles, currentStyle), nil
}

// Helper functions
func (p *CSSSwitcherPlugin) getAvailableStyles() ([]string, error) {
    cwd, err := os.Getwd()
    if err != nil {
        return nil, fmt.Errorf("failed to get working directory: %v", err)
    }

    cssDir := filepath.Join(cwd, "static", "css")
    if _, err := os.Stat(cssDir); os.IsNotExist(err) {
        return nil, fmt.Errorf("CSS directory does not exist: %s", cssDir)
    }

    entries, err := os.ReadDir(cssDir)
    if err != nil {
        return nil, fmt.Errorf("failed to read styles directory: %v", err)
    }

    var styles []string
    for _, entry := range entries {
        if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
            styleName := entry.Name()
            // Check if the directory has a CSS file matching its name
            if _, err := os.Stat(filepath.Join(cssDir, styleName, styleName+".css")); err == nil {
                styles = append(styles, styleName)
            }
        }
    }

    if len(styles) == 0 {
        return nil, fmt.Errorf("no valid styles found in %s", cssDir)
    }
    return styles, nil
}

func (p *CSSSwitcherPlugin) getCurrentStyle() string {
    cwd, err := os.Getwd()
    if err != nil {
        return "default"
    }

    mainCSSPath := filepath.Join(cwd, "static", "css", "style.css")
    content, err := os.ReadFile(mainCSSPath)
    if err != nil {
        return "default"
    }

    // Look for import statements like '@import url('themename/themename.css');'
    importContent := string(content)
    entries, err := p.getAvailableStyles()
    if err != nil {
        return "default"
    }

    for _, style := range entries {
        if strings.Contains(importContent, fmt.Sprintf("'%s/%s.css'", style, style)) {
            return style
        }
    }
    return "default"
}

func (p *CSSSwitcherPlugin) setStyle(style string) error {
    cwd, err := os.Getwd()
    if err != nil {
        return fmt.Errorf("failed to get working directory: %v", err)
    }

    cssDir := filepath.Join(cwd, "static", "css")
    styleDir := filepath.Join(cssDir, style)
    mainCSSPath := filepath.Join(cssDir, "style.css")

    // Validate style exists with matching CSS file
    styleCSSPath := filepath.Join(styleDir, style+".css")
    if _, err := os.Stat(styleCSSPath); os.IsNotExist(err) {
        return fmt.Errorf("style '%s' does not exist or is invalid", style)
    }

    // Create new style.css content with the new naming pattern
    content := fmt.Sprintf("@import url('%s/%s.css');", style, style)

    // Write to main CSS file
    if err := os.WriteFile(mainCSSPath, []byte(content), 0644); err != nil {
        return fmt.Errorf("failed to write style file: %v", err)
    }

    return nil
}