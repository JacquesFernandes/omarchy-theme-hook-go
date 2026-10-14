package steam

import (
	"os"
	"testing"

	"JacquesFernandes/omarchy_theme_hook_go/themed"
)

func TestCreateSteamTheme(t *testing.T) {
	themePath, err := CreateSteamTheme(themed.CatppuccinThemeMap)
	if err != nil {
		t.Error("Did not expect error when creating theme file, but got", err)
	}

	file, err := os.Open(themePath)
	defer file.Close()
	if err != nil {
		t.Error("Did not expect error when reading theme file, but got", err)
	}

	t.Log("File created at:", themePath)
}

func TestUpdateSteamTheme(t *testing.T) {
	themePath, err := CreateSteamTheme(themed.CatppuccinThemeMap)
	if err != nil {
		t.Error("Did not expect error when creating theme file, but got", err)
	}

	err = UpdateSteamTheme(themePath)
	if err != nil {
		t.Error("Did not expect error when updating theme file, but got", err)
	}

	t.Log("File created at:", themePath)
	t.Log("Steam theme upadted...")
}
