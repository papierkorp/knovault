package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"pewitima/internal/plugins"
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
	files := getAllFiles()
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
	files := getAllFiles()
	response := map[string][]string{
		"files": files,
	}
	return json.Marshal(response)
}

func (p *FileManagerPlugin) Execute(params map[string]string) (interface{}, error) {
	return templ.Raw(""), nil
}

func init() {
	plugins.RegisterCorePlugin("FileManager", &FileManagerPlugin{})
}


func getAllFiles() ([]string) {
	output := []string{}

	cwd, err1 := os.Getwd()
	if err1 != nil {
		return output
	}
	dataDir := filepath.Join(cwd, "data")

	err := filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("os.FileInfo: ", info)
		fmt.Println("FileName: ", info.Name())
		if !info.IsDir() {
			relativePath := strings.TrimPrefix(path, dataDir+"/")
			output = append(output, relativePath)
		}

		fmt.Println("output: ", output)

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	return output
}
