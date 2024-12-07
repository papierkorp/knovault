// internal/pluginManager/external/CustomCSS/main.go
package main

import (
    "knovault/internal/types"
    "knovault/internal/pluginManager/external/CustomCSS/plugin"
)

// Export Plugin symbol as an interface
var Plugin types.Plugin = plugin.NewCustomCSSPlugin()

func main() {}

