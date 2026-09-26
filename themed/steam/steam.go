package steam

import (
	"fmt"
	"text/template"

	"JacquesFernandes/omarchy_theme_hook_go/fileutils"
	"JacquesFernandes/omarchy_theme_hook_go/themed"
)

// CreateSteamTheme generates the CSS from a template and writes it a temp-file using data in themeMap
// This then returns the path to the generated temp-file if successful
func CreateSteamTheme(themeMap themed.ThemeMap) (string, error) {
	steamTemplate, err := template.New("steam").Parse(templateCSSRules)
	if err != nil {
		return "", fmt.Errorf("error creating Steam tempalte: %w", err)
	}

	outputTempFile, err := fileutils.CreateTempProjectFile("themed", "steam.css")
	if err != nil {
		return "", fmt.Errorf("error creating outputTempFile: %w", err)
	}
	defer outputTempFile.Close()

	err = steamTemplate.Execute(outputTempFile, themeMap)
	if err != nil {
		return "", fmt.Errorf("error executing template: %w", err)
	}

	return outputTempFile.Name(), nil
}
