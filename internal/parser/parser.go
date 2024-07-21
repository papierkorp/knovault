package parser

import (
    "io/ioutil"
    "os"
    "path/filepath"

    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/parser"
)

func ReadMarkdownFile(filename string) (string, error) {
    // Get the current working directory
    cwd, err := os.Getwd()
    if err != nil {
        return "", err
    }

    // Construct the path to the data directory
    dataDir := filepath.Join(cwd, "data")

    // Construct the full path to the file
    path := filepath.Join(dataDir, filename)

    content, err := ioutil.ReadFile(path)
    if err != nil {
        return "", err
    }

    extensions := parser.CommonExtensions | parser.AutoHeadingIDs
    parser := parser.NewWithExtensions(extensions)
    md := markdown.ToHTML(content, parser, nil)

    return string(md), nil
}
