package core

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "path/filepath"
    "pewitima/internal/plugins"

    "github.com/a-h/templ"
    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/parser"
    "github.com/microcosm-cc/bluemonday"
)

type ParserPlugin struct{}

func (p *ParserPlugin) Name() string {
    return "Parser"
}

func (p *ParserPlugin) Description() string {
    return "Parses Markdown files to HTML"
}

func (p *ParserPlugin) Help() string {
    return "This plugin takes a relative path to a Markdown file and returns the parsed HTML content. Pass the file path as a parameter when calling the plugin."
}

func (p *ParserPlugin) TemplResponse() (templ.Component, error) {
    // This method doesn't take parameters, so we can't pass a file path here.
    // We'll return an explanation instead.
    explanation := "To use the Parser plugin, please provide a file path when executing the plugin."
    return templ.Raw(explanation), nil
}

func (p *ParserPlugin) JsonResponse() ([]byte, error) {
    // Similarly, we can't pass a file path here, so we'll return an explanation.
    response := map[string]string{
        "message": "To use the Parser plugin, please provide a file path when executing the plugin.",
    }
    return json.Marshal(response)
}

func (p *ParserPlugin) Execute(params map[string]string) (interface{}, error) {
    filePath, ok := params["filePath"]
    if !ok {
        return nil, fmt.Errorf("filePath parameter is required")
    }

    html, err := p.parseMarkdownToHtml(filePath)
    if err != nil {
        return nil, err
    }
    return templ.Raw(html), nil
}

func (p *ParserPlugin) parseMarkdownToHtml(relativePath string) (string, error) {
    path := filepath.Join("data", relativePath)

    content, err := ioutil.ReadFile(path)
    if err != nil {
        return "", fmt.Errorf("error reading file: %v", err)
    }

    extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
    parser := parser.NewWithExtensions(extensions)
    maybeUnsafeHTML := markdown.ToHTML(content, parser, nil)
    html := bluemonday.UGCPolicy().SanitizeBytes(maybeUnsafeHTML)

    return string(html), nil
}

func init() {
    plugins.RegisterCorePlugin("Parser", &ParserPlugin{})
}