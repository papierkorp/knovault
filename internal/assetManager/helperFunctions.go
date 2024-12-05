package assetManager

import (
    "encoding/json"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
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

func compileModule(pkgPath, outputPath string) error {
    // Ensure the output directory exists
    outputDir := filepath.Dir(outputPath)
    if err := os.MkdirAll(outputDir, 0755); err != nil {
        return fmt.Errorf("failed to create output directory: %v", err)
    }

    // The main.go file should be in the parent directory of the plugin path
    mainFile := filepath.Join(filepath.Dir(pkgPath), "main.go")
    if _, err := os.Stat(mainFile); err != nil {
        return fmt.Errorf("main.go not found at %s: %v", mainFile, err)
    }

    // Get absolute paths
    absMainFile, err := filepath.Abs(mainFile)
    if err != nil {
        return fmt.Errorf("failed to get absolute path for main.go: %v", err)
    }

    absOutputPath, err := filepath.Abs(outputPath)
    if err != nil {
        return fmt.Errorf("failed to get absolute path for output: %v", err)
    }

    // Get current working directory
    cwd, err := os.Getwd()
    if err != nil {
        return fmt.Errorf("failed to get current working directory: %v", err)
    }

    // Build the module
    cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", absOutputPath, absMainFile)
    cmd.Dir = cwd // Set working directory to project root
    cmd.Env = append(os.Environ(),
        "CGO_ENABLED=1",
        "GO111MODULE=on",
        fmt.Sprintf("GOPATH=%s", os.Getenv("GOPATH")),
    )

    if output, err := cmd.CombinedOutput(); err != nil {
        return fmt.Errorf("compilation failed: %v\nOutput: %s", err, output)
    }

    return nil
}