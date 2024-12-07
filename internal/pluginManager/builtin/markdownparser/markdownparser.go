// internal/pluginManager/builtin/markdownparser/markdownparser.go
package markdownparser

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "path/filepath"

    "github.com/a-h/templ"
    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/parser"
    "github.com/microcosm-cc/bluemonday"
)

type MarkdownParserPlugin struct{}

func (p *MarkdownParserPlugin) Name() string {
    return "Parser"
}

func (p *MarkdownParserPlugin) Description() string {
    return "Parses Markdown files to HTML"
}

func (p *MarkdownParserPlugin) Help() string {
    return "This plugin takes a relative path to a Markdown file and returns the parsed HTML content. Pass the file path as a parameter when calling the plugin."
}

func (p *MarkdownParserPlugin) TemplResponse() (templ.Component, error) {
    explanation := "To use the Parser plugin, please provide a file path when executing the plugin."
    return templ.Raw(explanation), nil
}

func (p *MarkdownParserPlugin) JsonResponse() ([]byte, error) {
    response := map[string]string{
        "message": "To use the Parser plugin, please provide a file path when executing the plugin.",
    }
    return json.Marshal(response)
}

func (p *MarkdownParserPlugin) Execute(params map[string]string) (interface{}, error) {
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

func (p *MarkdownParserPlugin) parseMarkdownToHtml(relativePath string) (string, error) {
    path := filepath.Join("data", relativePath)

    content, err := ioutil.ReadFile(path)
    if err != nil {
        return "", fmt.Errorf("error reading file: %v", err)
    }

    extensions := parser.CommonExtensions |
                 parser.AutoHeadingIDs |
                 parser.NoEmptyLineBeforeBlock |
                 parser.Tables |
                 parser.FencedCode |
                 parser.Footnotes |
                 parser.HeadingIDs |
                 parser.Strikethrough

    mdParser := parser.NewWithExtensions(extensions)
    maybeUnsafeHTML := markdown.ToHTML(content, mdParser, nil)

    policy := bluemonday.UGCPolicy()
    policy.AllowElements("table", "thead", "tbody", "tr", "th", "td")
    policy.AllowElements("h1", "h2", "h3", "h4", "h5", "h6")
    policy.AllowElements("p", "br", "hr")
    policy.AllowElements("ul", "ol", "li")
    policy.AllowElements("strong", "em", "code", "pre")
    policy.AllowElements("a")
    policy.AllowAttrs("href").OnElements("a")
    policy.AllowAttrs("class").OnElements("code", "pre")
    policy.AllowAttrs("id").OnElements("h1", "h2", "h3", "h4", "h5", "h6")

    html := policy.SanitizeBytes(maybeUnsafeHTML)
    return string(html), nil
}