package types

// Config types for asset manager
type ThemeConfig struct {
    Themes []ThemeMetadata `json:"themes"`
}

type PluginConfig struct {
    Plugins []PluginMetadata `json:"plugins"`
}

type ThemeMetadata struct {
    Name    string   `json:"name"`
    Path    string   `json:"path"`
    Enabled bool     `json:"enabled"`
    Tags    []string `json:"tags"`
}

type PluginMetadata struct {
    Name    string   `json:"name"`
    Path    string   `json:"path"`
    Enabled bool     `json:"enabled"`
    Tags    []string `json:"tags"`
}