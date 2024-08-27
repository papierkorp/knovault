package filemanager

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetAllFiles() ([]string, error) {
	fmt.Println("RUNNING GetAllFiles()!!!!!!!!!!!!!!!!!!")

	output := []string{}

	cwd, err1 := os.Getwd()
	if err1 != nil {
		return output, err1
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
			output = append(output, path)
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	return output, err
}
