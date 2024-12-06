package assetManager

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "log"
)

func loadConfig[T any](path string) (*T, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("error reading config file %s: %v", path, err)
    }

    var config T
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("error parsing config file %s: %v", path, err)
    }

    return &config, nil
}

func hasTag(tags []string, target string) bool {
    for _, tag := range tags {
        if tag == target {
            return true
        }
    }
    return false
}

func findAssetFiles(baseDir string) ([]string, error) {
    var files []string

    // Walk through the directory
    err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Skip directories themselves
        if info.IsDir() {
            return nil
        }

        // Look for .so files
        if filepath.Ext(path) == ".so" {
            files = append(files, path)
            log.Printf("Found .so file: %s", path)
            return nil
        }

        // Look for main.go files in the plugin directory
        if filepath.Base(path) == "plugin.go" &&
           filepath.Base(filepath.Dir(path)) == "plugin" {
            files = append(files, path)
            log.Printf("Found plugin.go file: %s", path)
        }

        return nil
    })

    if err != nil {
        return nil, err
    }

    if len(files) == 0 {
        log.Printf("No asset files found in directory: %s", baseDir)
        // Log the directory contents for debugging
        entries, err := os.ReadDir(baseDir)
        if err != nil {
            log.Printf("Error reading directory: %v", err)
        } else {
            log.Printf("Directory contents of %s:", baseDir)
            for _, entry := range entries {
                log.Printf("  %s (isDir: %v)", entry.Name(), entry.IsDir())
            }
        }
    }

    return files, nil
}