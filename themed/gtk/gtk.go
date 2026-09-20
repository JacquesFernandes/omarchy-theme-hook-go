package gtk

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"JacquesFernandes/omarchy_theme_hook_go/fileutils"
)

// createGtkTheme generates the CSS from a "template"
func CreateGtkTheme(themeMap map[string]string) string {
	gtkTemplate, err := template.New("gtk").Parse(templateCSSRules)
	if err != nil {
		log.Panic("error creating template:", err)
	}

	tempDirPath := filepath.Join(os.TempDir(), "omarchy-theme-hook-go", "themed")
	err = os.MkdirAll(tempDirPath, 0o755)
	if err != nil {
		log.Panic("error creating temp dir path:", err)
	}

	outputTempFile, err := os.CreateTemp(tempDirPath, "gtk.css")
	if err != nil {
		log.Panic("error creating outputTempFile:", err)
	}
	defer outputTempFile.Close()

	// Deliberately only execute and write the template on a limited set of lines
	err = gtkTemplate.Execute(outputTempFile, themeMap)
	if err != nil {
		log.Panic("error executing template:", err)
	}

	// write the remaining static CSS rules
	_, err = outputTempFile.WriteString(remainingCSSRules)
	if err != nil {
		log.Panic("error trying to write remaining GTK theme Rules:", err)
	}
	return outputTempFile.Name()
}

// UpdateGtkThemeFiles takes the path of the generated CSS file and updates the gtk-3.0 and gtk-4.0 css files
// The newThemeFilePath is usually created by the CreateGtkTheme function
func UpdateGtkThemeFiles(newThemeFilePath string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	gtk3DirPath := filepath.Join(homeDir, ".config", "gtk-3.0")
	gtk3CSSPath := filepath.Join(gtk3DirPath, "gtk.css")
	gtk3CSSBackupPath := filepath.Join(gtk3DirPath, fmt.Sprintf("gtk.css.%d", now))
	gtk4DirPath := filepath.Join(homeDir, ".config", "gtk-4.0")
	gtk4CSSPath := filepath.Join(gtk4DirPath, "gtk.css")
	gtk4CSSBackupPath := filepath.Join(gtk4DirPath, fmt.Sprintf("gtk.css.%d", now))

	log.Println("gtk 3", gtk3DirPath, gtk3CSSPath, gtk3CSSBackupPath)
	log.Println("gtk 4", gtk4DirPath, gtk4CSSPath, gtk4CSSBackupPath)

	// rename the existing gtk3 CSS file to a backup
	err = os.Rename(gtk3CSSPath, gtk3CSSBackupPath)
	if err != nil {
		return err
	}

	// Copy new CSS file to gtk3 location
	if err = fileutils.CopyFile(newThemeFilePath, gtk3CSSPath); err != nil {
		return err
	}

	// rename the existing gtk4 CSS file to a backup
	err = os.Rename(gtk4CSSPath, gtk4CSSBackupPath)
	if err != nil {
		return err
	}

	// Copy new CSS file to gtk4 location
	if err = fileutils.CopyFile(newThemeFilePath, gtk4CSSPath); err != nil {
		return err
	}

	return nil
}
