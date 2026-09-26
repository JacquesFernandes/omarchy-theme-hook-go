package gtk

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"JacquesFernandes/omarchy_theme_hook_go/fileutils"
	"JacquesFernandes/omarchy_theme_hook_go/themed"
)

// createGtkTheme generates the CSS from a "template"
func CreateGtkTheme(themeMap themed.ThemeMap) (string, error) {
	gtkTemplate, err := template.New("gtk").Parse(templateCSSRules)
	if err != nil {
		return "", fmt.Errorf("error creating GTK template: %w", err)
	}

	outputTempFile, err := fileutils.CreateTempProjectFile("themed", "gtk.css")
	if err != nil {
		return "", fmt.Errorf("error creating outputTempFile: %w", err)
	}
	defer outputTempFile.Close()

	// Deliberately only execute and write the template on a limited set of lines
	err = gtkTemplate.Execute(outputTempFile, themeMap)
	if err != nil {
		return "", fmt.Errorf("error executing template:", err)
	}

	// write the remaining static CSS rules
	_, err = outputTempFile.WriteString(remainingCSSRules)
	if err != nil {
		return "", fmt.Errorf("error trying to write remaining GTK theme Rules:", err)
	}
	return outputTempFile.Name(), nil
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
