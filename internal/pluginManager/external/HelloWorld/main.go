// internal/pluginManager/external/HelloWorld/main.go
package main

import (
    "knovault/internal/types"
    "knovault/internal/pluginManager/external/HelloWorld/plugin"
)

// Export Plugin symbol as an interface
var Plugin types.Plugin = &plugin.HelloWorldPlugin{}

func main() {}

