package main

import (
    "gowiki/internal/server"
)

func main() {
    server.Start()
}

func getFileList() []string {
    all_files, err := filemanager.GetAllFiles()
    fmt.Println("all_files: ", all_files)
    if err != nil {
        fmt.Println("Error fetching files:", err)
        all_files = []string{}
    }

    return all_files
}
