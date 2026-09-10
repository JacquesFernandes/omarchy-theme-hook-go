package colors

import (
	"log"
	"strings"
	"testing"
)

func TestLoadColorsToml(t *testing.T) {
	catppuccinThemeToml := "/usr/share/omarchy/themes/catppuccin/colors.toml"

	expectedKeys := []string{
		"mode",

		"accent",
		"selection",
		"muted",

		"background",
		"dark_background",
		"darker_background",
		"lighter_background",

		"foreground",
		"dark_foreground",
		"light_foreground",
		"bright_foreground",

		"red",
		"yellow",
		"orange",
		"green",
		"cyan",
		"blue",
		"magenta",
		"brown",

		"bright_red",
		"bright_yellow",
		"bright_green",
		"bright_cyan",
		"bright_blue",
		"bright_magenta",
	}

	themeMap := LoadColorsToml(catppuccinThemeToml)

	for _, expectedKey := range expectedKeys {
		t.Run(expectedKey, func(t *testing.T) {
			value, ok := themeMap[expectedKey]

			if !ok {
				t.Fatal("Missing key:", expectedKey)
			}

			if len(value) == 0 {
				t.Fatal("Key", expectedKey, "has an empty value")
			}

			if len(strings.TrimSpace(value)) != len(value) {
				t.Fatal("Key", expectedKey, "value has a space around it")
			}
		})
	}

	log.Println("Loaded Catppuccin theme:", themeMap)
}
