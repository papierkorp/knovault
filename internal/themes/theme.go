package themes

import (
	"fmt"
	"gowiki/internal/types"
)


var (
	currentTheme types.Theme
	themes       = make(map[string]types.Theme)
)

func RegisterTheme(name string, theme types.Theme) {
	fmt.Println("Register themes before: ", themes)
	themes[name] = theme
	fmt.Println("Register themes after: ", themes)
}

func SetCurrentTheme(name string) error {
	fmt.Println("THEME NAME: ", name)
	fmt.Println("SET CURRENT THEME")
	fmt.Println("SetCurrentTheme themes before: ", themes)
	theme, ok := themes[name]
	if !ok {
		return fmt.Errorf("theme %s not found", name)
	}
	currentTheme = theme
	fmt.Println("SetCurrentTheme themes after: ", themes)
	return nil
}

func GetCurrentTheme() types.Theme {
	fmt.Println("themes: ", themes)
	fmt.Println("currentTheme: ", currentTheme)
	return currentTheme
}