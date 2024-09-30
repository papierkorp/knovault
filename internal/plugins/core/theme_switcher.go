package core

import (
	"context"
	"gowiki/internal/themes"
	"io"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type ThemeSwitcherPlugin struct{}

func (p *ThemeSwitcherPlugin) Name() string {
    return "ThemeSwitcher"
}

func (p *ThemeSwitcherPlugin) Initialize() error {
    // Any initialization logic
    return nil
}

func (p *ThemeSwitcherPlugin) Shutdown() error {
    // Any cleanup logic
    return nil
}

func (p *ThemeSwitcherPlugin) Handlers() map[string]echo.HandlerFunc {
    return map[string]echo.HandlerFunc{
        "/change-theme": p.changeTheme,
    }
}

func (p *ThemeSwitcherPlugin) TemplateData() map[string]interface{} {
    return map[string]interface{}{
        "availableThemes": themes.GetAvailableThemes(),
    }
}

func (p *ThemeSwitcherPlugin) changeTheme(c echo.Context) error {
    var req struct {
        Theme string `json:"theme"`
    }
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false})
    }

    if err := themes.SetCurrentTheme(req.Theme); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{"success": false})
    }

    return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

func (p *ThemeSwitcherPlugin) Render() templ.Component {
    return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
        _, err = io.WriteString(w, `
            <select id="theme-select" class="px-4 py-2 bg-gray-200 dark:bg-gray-700 rounded">
                <option selected value="defaultTheme">Default Theme</option>
                <option value="dark">Dark</option>
                // Add more theme options here as needed
            </select>
            <script>
                document.addEventListener('DOMContentLoaded', () => {
                    const themeSelect = document.getElementById('theme-select');
                    
                    // Load the current theme from localStorage
                    const currentTheme = localStorage.getItem('appTheme') || 'defaultTheme';
                    themeSelect.value = currentTheme;

                    themeSelect.addEventListener('change', (event) => {
                        const selectedTheme = event.target.value;
                        localStorage.setItem('appTheme', selectedTheme);

                        // Make an AJAX request to change the theme
                        fetch('/change-theme', {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json',
                            },
                            body: JSON.stringify({ theme: selectedTheme }),
                        })
                        .then(response => {
                            console.log('Response status:', response.status);
                            return response.json().then(data => ({ status: response.status, body: data }));
                        })
                        .then(({ status, body }) => {
                            console.log('Response body:', body);
                            if (status === 200 && body.success) {
                                console.log('Theme changed successfully. Reloading...');
                                window.location.reload();
                            } else {
                                console.error('Failed to change theme. Status:', status, 'Body:', body);
                            }
                        })
                        .catch(error => {
                            console.error('Error:', error);
                        });
                    });
                });
            </script>
        `)
        return
    })
}