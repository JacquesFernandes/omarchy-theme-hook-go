package gtk

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"
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
		log.Panic("error trying to get user home dir:", err)
	}

	gtk3DirPath := filepath.Join(homeDir, ".config", "gtk-3.0")
	gtk3CSSPath := filepath.Join(gtk3DirPath, "gtk.css")
	gtk3CSSBackupPath := filepath.Join(gtk3DirPath, fmt.Sprintf("gtk.css.%d", time.Now().Unix()))
	gtk4DirPath := filepath.Join(homeDir, ".config", "gtk-4.0")
	gtk4CSSPath := filepath.Join(gtk4DirPath, "gtk.css")
	gtk4CSSBackupPath := filepath.Join(gtk4DirPath, fmt.Sprintf("gtk.css.%d", time.Now().Unix()))

	log.Println("gtk 3", gtk3DirPath, gtk3CSSPath, gtk3CSSBackupPath)
	log.Println("gtk 4", gtk4DirPath, gtk4CSSPath, gtk4CSSBackupPath)

	// Get read access to generated css file. Panic if not available
	sourceCSSFile, err := os.Open(newThemeFilePath)
	if err != nil {
		log.Panic("error trying to read generated CSS temp file:", sourceCSSFile)
	}
	defer sourceCSSFile.Close()

	// rename the existing gtk3 CSS file to a backup
	err = os.Rename(gtk3CSSPath, gtk3CSSBackupPath)
	if err != nil {
		log.Panic("error when creating backup of gtk3 css:", err)
	}

	// Create the new GTK3 CSS file
	newGTK3CSSFile, err := os.Create(gtk3CSSPath)
	if err != nil {
		log.Panic("error trying to create target GTK3 CSS file:", err)
	}
	defer newGTK3CSSFile.Close()

	// Write/Copy the generated CSS file into the gtk3 dir as the new gtk.css
	_, err = io.Copy(newGTK3CSSFile, sourceCSSFile)
	if err != nil {
		log.Panic("error trying to copy generated CSS file into new GTK3 locationn")
	}

	// TODO: Write the logic for GTK4

	return nil
}
