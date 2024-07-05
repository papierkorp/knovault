package server

import (
	templruntime "github.com/a-h/templ/runtime"
	"github.com/labstack/echo/v4"
	"gowiki/views"
	"net/http"
	"strings"
)

func Start() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		var sb strings.Builder
		err := templruntime.RenderComponent(&sb, views.Hello("John"))
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error rendering template")
		}
		return c.String(http.StatusOK, sb.String())
	})
	e.Logger.Fatal(e.Start(":1323"))
}
