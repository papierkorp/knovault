// internal/themeManager/themeManagerBuiltinHelper.go
package themeManager

import (
    "knovault/internal/types"
    "knovault/internal/themeManager/builtin/defaulttheme"
)

func loadBuiltinThemes() (map[string]types.Theme, error) {
    themes := make(map[string]types.Theme)

    // Initialize builtin themes
    themes["defaultTheme"] = &defaulttheme.DefaultTheme{}

    return themes, nil
}