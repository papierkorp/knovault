#!/bin/bash

# Create new directory structure
mkdir -p internal/pluginManager/builtin
mkdir -p internal/pluginManager/external
mkdir -p internal/themeManager/builtin
mkdir -p internal/themeManager/external

# Move and restructure builtin plugins
for plugin in FileManager MarkdownParser ThemeChanger; do
    mkdir -p "internal/pluginManager/builtin/${plugin,,}"
    mv "internal/assetManager/plugins/$plugin/plugin/plugin.go" "internal/pluginManager/builtin/${plugin,,}/${plugin,,}.go"
done

# Move external plugins
for plugin in CustomCSS HelloWorld; do
    mkdir -p "internal/pluginManager/external/$plugin"
    cp -r "internal/assetManager/plugins/$plugin"/* "internal/pluginManager/external/$plugin/"
done

# Move and restructure builtin theme
mkdir -p internal/themeManager/builtin/defaulttheme
cp -r internal/assetManager/themes/defaultTheme/plugin/plugin.go internal/themeManager/builtin/defaulttheme/defaulttheme.go
cp -r internal/assetManager/themes/defaultTheme/templates internal/themeManager/builtin/defaulttheme/

# Move external theme
mkdir -p internal/themeManager/external/dark
cp -r internal/assetManager/themes/dark/* internal/themeManager/external/dark/

# Create new README files
cp internal/assetManager/README.md internal/pluginManager/README.md
cp internal/assetManager/README.md internal/themeManager/README.md

# Remove old assetManager directory
rm -rf internal/assetManager

# Remove old JSONs
rm -f internal/assetManager/plugins_list.json internal/assetManager/themes_list.json

# Clean up empty directories
find . -type d -empty -delete
