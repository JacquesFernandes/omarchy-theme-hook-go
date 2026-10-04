package steam

import (
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// adwaitaForSteamPath generates and returns the path that we expect
// "Adwaita-for-Steam" to be at, if successful
func adwaitaForSteamPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Panic("error trying to get user HOME dir:", err)
	}

	return filepath.Join(homeDir, ".local", "share", "steam-adwaita")
}

// checkForAdwaitaForSteam checks if "Adwaita-for-Steam" is cloned to "~/.local/share/steam-adwaita/"
// Then, it checks if the "install.py" is present
//
// In the chance that "Adwaita-for-Steam" is installed somewhere else, this
// allows the user to just symlink the install.py
func checkForAdwaitaForSteam() bool {
	afsPath := adwaitaForSteamPath()

	files, err := os.ReadDir(afsPath)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		log.Panic("unhandled error when trying to read adwaita-for-steam dir:", err)
	}

	return slices.ContainsFunc(files, func(file os.DirEntry) bool {
		return strings.Contains(file.Name(), "install.py")
	})
}

// createAdwaitaForSteamOmarchyThemeDir uses os.MkdirAll to create the directory where
// the omarchy.css file will live (that will then be applied to Steam).
//
// This is safe to run multiple times as it'll just finish quietly if the directory
// already exists
func createAdwaitaForSteamOmarchyThemeDir() error {
	afsPath := adwaitaForSteamPath()
	omarchyThemeDirPath := filepath.Join(afsPath, "adwaita", "colorthemes", "omarchy")
	return os.MkdirAll(omarchyThemeDirPath, 0o755)
}
