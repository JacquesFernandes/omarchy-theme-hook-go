package gtk

import (
	"log"
	"testing"
)

// Taken from the Catppuccin theme
// This theme is slightly darker
var catppuccinThemeMap map[string]string = map[string]string{
	"mode": "dark",

	"accent":    "#89b4fa",
	"selection": "#45475a",
	"muted":     "#585b70",

	"background":         "#1e1e2e",
	"dark_background":    "#161622",
	"darker_background":  "#101019",
	"lighter_background": "#313244",

	"foreground":        "#cdd6f4",
	"dark_foreground":   "#6c7086",
	"light_foreground":  "#bac2de",
	"bright_foreground": "#cdd6f4",

	"red":     "#f38ba8",
	"yellow":  "#f9e2af",
	"orange":  "#f6b6ab",
	"green":   "#a6e3a1",
	"cyan":    "#94e2d5",
	"blue":    "#89b4fa",
	"magenta": "#f5c2e7",
	"brown":   "#7b5b55",

	"bright_red":     "#f38ba8",
	"bright_yellow":  "#f9e2af",
	"bright_green":   "#a6e3a1",
	"bright_cyan":    "#94e2d5",
	"bright_blue":    "#89b4fa",
	"bright_magenta": "#f5c2e7",
}

// Taken from the Lupine Theme
// This theme is light
var lupineThemeMap map[string]string = map[string]string{
	"mode": "light",

	"accent":    "#3264eb",
	"selection": "#d0d0d0",
	"muted":     "#9e9e9e",

	"background":         "#fafafa",
	"dark_background":    "#ececec",
	"darker_background":  "#dedede",
	"lighter_background": "#f5f5f5",

	"foreground":        "#212121",
	"dark_foreground":   "#757575",
	"light_foreground":  "#424242",
	"bright_foreground": "#000000",

	"red":     "#c900c4",
	"yellow":  "#026fde",
	"orange":  "#026fde",
	"green":   "#4a2fd0",
	"cyan":    "#0c67de",
	"blue":    "#3264eb",
	"magenta": "#8a4ad7",
	"brown":   "#013a6f",

	"bright_red":     "#f930fb",
	"bright_yellow":  "#358fff",
	"bright_green":   "#9f85e0",
	"bright_cyan":    "#3986ff",
	"bright_blue":    "#5482ff",
	"bright_magenta": "#b363ff",
}

func TestCreateGtkTheme(t *testing.T) {
	tempFilePath, err := CreateGtkTheme(lupineThemeMap)
	if err != nil {
		t.Error("Got unexpected error:", err)
	}
	log.Print(tempFilePath)
}

func TestUpdateGtkThemeFiles(t *testing.T) {
	tempThemePath, err := CreateGtkTheme(catppuccinThemeMap)
	if err != nil {
		t.Errorf("CreateGtkTheme failed for some reason:", err)
	}
	UpdateGtkThemeFiles(tempThemePath)
}
