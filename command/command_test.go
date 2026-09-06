package command

import (
	"log"
	"testing"
)

func TestOmarchyThemeList(t *testing.T) {
	themes, err := OmarchyThemeList()
	if err != nil {
		t.Error(err)
	}

	for _, theme := range themes {
		log.Println(theme)
	}
}

func TestOmarchyThemeDir(t *testing.T) {
	themes, err := OmarchyThemeList()
	if err != nil {
		t.Error(err)
	}

	for _, theme := range themes {
		path, err := OmarchyThemeDir(theme)
		if err != nil {
			t.Error(err)
		}

		log.Println(theme, path)
	}
}

func TestOmarchyThemeCurrent(t *testing.T) {
	theme, err := OmarchyThemeCurrent()
	if err != nil {
		t.Error(err)
	}

	log.Println("current theme:", theme)
}
