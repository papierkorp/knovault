package core

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type DarkModePlugin struct{}

func (p *DarkModePlugin) Name() string {
    return "DarkMode"
}

func (p *DarkModePlugin) Initialize() error {
    // Any initialization logic
    return nil
}

func (p *DarkModePlugin) Shutdown() error {
    // Any cleanup logic
    return nil
}

func (p *DarkModePlugin) Handlers() map[string]echo.HandlerFunc {
    return map[string]echo.HandlerFunc{}
}

func (p *DarkModePlugin) TemplateData() map[string]interface{} {
    return map[string]interface{}{
        "darkModeEnabled": true,
    }
}

func (p *DarkModePlugin) Render() templ.Component {
    return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
        _, err = io.WriteString(w, `
            <button id="theme-toggle">Toggle Dark Mode</button>
            <script>
                document.addEventListener('DOMContentLoaded', () => {
                  const toggleButton = document.getElementById('theme-toggle');
                  const htmlElement = document.documentElement;

                  // Check localStorage for theme preference
                  const currentTheme = localStorage.getItem('theme');
                  if (currentTheme === 'dark') {
                    htmlElement.classList.add('dark');
                  } else {
                    htmlElement.classList.remove('dark');
                  }

                  // Toggle dark mode on button click
                  toggleButton.addEventListener('click', () => {
                    if (htmlElement.classList.contains('dark')) {
                      htmlElement.classList.remove('dark');
                      localStorage.setItem('theme', 'light'); // Save preference
                    } else {
                      htmlElement.classList.add('dark');
                      localStorage.setItem('theme', 'dark'); // Save preference
                    }
                  });
                });
            </script>
        `)
        return
    })
}