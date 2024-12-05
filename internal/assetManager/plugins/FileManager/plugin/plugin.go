package plugin

import (
    "context"
    "encoding/json"
    "io"
    "os"
    "path/filepath"
    "strings"

    "github.com/a-h/templ"
)

type FileManagerPlugin struct{}

func (p *FileManagerPlugin) Name() string {
    return "FileManager"
}

func (p *FileManagerPlugin) Description() string {
    return "Lists all files in the data directory"
}

func (p *FileManagerPlugin) Help() string {
    return "This plugin lists all files in the data directory. It doesn't require any input."
}

func (p *FileManagerPlugin) TemplResponse() (templ.Component, error) {
    files, err := getAllFiles()
    if err != nil {
        return nil, err
    }
    return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
        _, err = io.WriteString(w, "<ul>")
        if err != nil {
            return err
        }
        for _, file := range files {
            _, err = io.WriteString(w, "<li>"+file+"</li>")
            if err != nil {
                return err
            }
        }
        _, err = io.WriteString(w, "</ul>")
        return err
    }), nil
}

func (p *FileManagerPlugin) JsonResponse() ([]byte, error) {
    files, err := getAllFiles()
    if err != nil {
        return nil, err
    }
    response := map[string][]string{
        "files": files,
    }
    return json.Marshal(response)
}

func (p *FileManagerPlugin) Execute(params map[string]string) (interface{}, error) {
    files, err := getAllFiles()
    if err != nil {
        return nil, err
    }
    return files, nil
}

// Helper functions
func getAllFiles() ([]string, error) {
    var output []string

    cwd, err := os.Getwd()
    if err != nil {
        return nil, err
    }
    dataDir := filepath.Join(cwd, "data")

    err = filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            relativePath := strings.TrimPrefix(path, dataDir+string(os.PathSeparator))
            output = append(output, relativePath)
        }
        return nil
    })
    if err != nil {
        return nil, err
    }

    return output, nil
}

// Export the plugin instance
var Plugin = &FileManagerPlugin{}