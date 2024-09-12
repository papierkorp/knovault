package themes

import (
	"fmt"
	"gowiki/internal/types"
)


var (
	currentTheme types.RoutesHandler
	themes       = make(map[string]types.RoutesHandler)
)

func RegisterTheme(name string, theme types.RoutesHandler) {
	fmt.Println("Register themes before: ", themes)
	themes[name] = theme
	fmt.Println("Register themes after: ", themes)
}

func SetCurrentTheme(name string) error {
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

func GetCurrentTheme() types.RoutesHandler {
	fmt.Println("themes: ", themes)
	fmt.Println("currentTheme: ", currentTheme)
	return currentTheme
}