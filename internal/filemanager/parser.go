package filemanager

import (
    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/parser"
    "github.com/microcosm-cc/bluemonday"
    "io/ioutil"
    "path/filepath"
)

func ParseMarkdownToHtml(relativePath string) (string) {
    // cwd, err := os.Getwd()
    // if err != nil {
    //     return "", err
    // }

    // dataDir := filepath.Join(cwd, "data")
    // path := filepath.Join(dataDir, filename)

    path := filepath.Join("data", relativePath)

    content, err := ioutil.ReadFile(path)
    if err != nil {
        return ""
    }

    extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
    parser := parser.NewWithExtensions(extensions)
    maybeUnsafeHTML := markdown.ToHTML(content, parser, nil)
    html := bluemonday.UGCPolicy().SanitizeBytes(maybeUnsafeHTML)

    return string(html)
}
