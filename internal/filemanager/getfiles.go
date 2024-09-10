package filemanager

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)


func GetAllFiles() ([]string) {
	fmt.Println("RUNNING GetAllFiles()!!!!!!!!!!!!!!!!!!")

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

