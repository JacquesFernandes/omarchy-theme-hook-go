package omarchy

import (
	"log"
	"testing"
)

func TestThemeList(t *testing.T) {
	themes, err := ThemeList()
	if err != nil {
		t.Error(err)
	}

	for _, theme := range themes {
		log.Println(theme)
	}
}

func TestThemeDir(t *testing.T) {
	themes, err := ThemeList()
	if err != nil {
		t.Error(err)
	}

	for _, theme := range themes {
		path, err := ThemeDir(theme)
		if err != nil {
			t.Error(err)
		}

		log.Println(theme, path)
	}
}

func TestCurrentThemeName(t *testing.T) {
	theme, err := CurrentThemeName()
	if err != nil {
		t.Error(err)
	}

	log.Println("current theme:", theme)
}
