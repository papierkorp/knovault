package main

import (
    "knovault/internal/assetManager/plugins/HelloWorld/plugin"
    "knovault/internal/types"
)

// Export Plugin symbol as an interface
var Plugin types.Plugin = plugin.Plugin

func main() {}