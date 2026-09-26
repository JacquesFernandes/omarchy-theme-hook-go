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
