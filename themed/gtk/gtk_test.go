package gtk

import (
	"log"
	"testing"

	"JacquesFernandes/omarchy_theme_hook_go/themed"
)

func TestCreateGtkTheme(t *testing.T) {
	tempFilePath, err := CreateGtkTheme(themed.LupineThemeMap)
	if err != nil {
		t.Error("Got unexpected error:", err)
	}
	log.Print(tempFilePath)
}

func TestUpdateGtkThemeFiles(t *testing.T) {
	tempThemePath, err := CreateGtkTheme(themed.CatppuccinThemeMap)
	if err != nil {
		t.Error("CreateGtkTheme failed for some reason:", err)
	}
	UpdateGtkThemeFiles(tempThemePath)
}
