package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"knovault/internal/plugins"
	"knovault/internal/plugins/templates/CardCreator"
	"knovault/internal/types"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CardCreatorPlugin struct{}

func (p *CardCreatorPlugin) Name() string {
    return "CardCreator"
}

func (p *CardCreatorPlugin) Description() string {
    return "Create and manage custom card layouts for documents"
}

func (p *CardCreatorPlugin) Help() string {
    return "Use this plugin to design custom card layouts for your documents using a drag-and-drop interface."
}

func (p *CardCreatorPlugin) TemplResponse() (templ.Component, error) {
    return CardCreator.CardCreatorInterface(), nil
}

func (p *CardCreatorPlugin) JsonResponse() ([]byte, error) {
    return json.Marshal(map[string]string{"status": "ok"})
}

func (p *CardCreatorPlugin) Execute(params map[string]string) (interface{}, error) {
    // Implement execution logic if needed
    return nil, nil
}

func (p *CardCreatorPlugin) Route() types.PluginRoute {
    return types.PluginRoute{
        Method: "POST",
        Path:   "/plugins/CardCreator/update-preview",
        Handler: func(c echo.Context) error {
            // Parse form data
            cardWidth := c.FormValue("cardWidth")
            cardHeight := c.FormValue("cardHeight")
            // ... parse other form values as needed

            // Generate preview HTML
            previewHTML := fmt.Sprintf(`
                <div style="width:%spx;height:%spx;border:1px solid black;position:relative;">
                    <!-- Add more elements based on the form data -->
                    <div style="position:absolute;top:10px;left:10px;">Preview Content</div>
                </div>
            `, cardWidth, cardHeight)

            return c.HTML(http.StatusOK, previewHTML)
        },
    }
}

func init() {
    plugins.RegisterCorePlugin("CardCreator", &CardCreatorPlugin{})
}