package main

import (
	"gowiki/internal/server"
	"gowiki/internal/themes"
	"log"
    _ "gowiki/internal/themes/defaultTheme" 
)

func main() {
    err := themes.SetCurrentTheme("defaultTheme")
    if err != nil {
        log.Fatalf("Failed to set default theme: %v", err)
    }

    server.Start()
}
