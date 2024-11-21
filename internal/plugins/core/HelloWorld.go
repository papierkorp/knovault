package core

import (
	"encoding/json"
	"knovault/internal/plugins"
	"time"

	"github.com/a-h/templ"
)

type HelloWorldPlugin struct{}

func (p *HelloWorldPlugin) Name() string {
	return "HelloWorld"
}

func (p *HelloWorldPlugin) Description() string {
	return "Returns the current time and date"
}

func (p *HelloWorldPlugin) Help() string {
	return "This plugin doesn't require any input. It returns the current time and date."
}

func (p *HelloWorldPlugin) TemplResponse() (templ.Component, error) {
	currentTime := time.Now().Format(time.RFC3339)
	return templ.Raw(currentTime), nil
}

func (p *HelloWorldPlugin) JsonResponse() ([]byte, error) {
	response := map[string]string{
		"currentTime": time.Now().Format(time.RFC3339),
	}
	return json.Marshal(response)
}

func (p *HelloWorldPlugin) Execute(params map[string]string) (interface{}, error) {
	return templ.Raw(""), nil
}

func init() {
	plugins.RegisterCorePlugin("HelloWorld", &HelloWorldPlugin{})
}