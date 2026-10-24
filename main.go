// Package omarchythemehookgo is a golang replacement for omarchy-theme-hook
//
// Honestly? A personal attempt at trying to fix broken themeing after Omarchy 3 -> 4
// Will stick to Go standard lib as much as possible
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"JacquesFernandes/omarchy_theme_hook_go/colors"
	"JacquesFernandes/omarchy_theme_hook_go/omarchy"
	"JacquesFernandes/omarchy_theme_hook_go/themed/gtk"
)

func prompt(question string) string {
	fmt.Println(question)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal("error trying to prompt:", err)
	}

	return scanner.Text()
}

func main() {
	currentTheme, err := omarchy.CurrentThemeName()
	if err != nil {
		log.Fatal("error trying to get current theme name:", err)
	}

	themeDirPath, err := omarchy.ThemeDirPath(currentTheme)
	if err != nil {
		log.Fatal("error trying to get current theme path:", err)
	}

	themeColorsTomlPath := filepath.Join(themeDirPath, "colors.toml")

	themeColorsToml := colors.LoadColorsToml(themeColorsTomlPath)

	gtkThemePath, err := gtk.CreateGtkTheme(themeColorsToml)
	if err != nil {
		log.Fatal("error trying to create (temp) GTK theme CSS:", err)
	}

	response := prompt(fmt.Sprintf("Temp GTK theme CSS file created at %q. Continue to apply it? (y/n):", gtkThemePath))

	if response != "y" {
		log.Fatal("Created theme file not confirmed. Quitting.")
	}
}
