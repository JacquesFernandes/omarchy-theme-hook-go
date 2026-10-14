package steam

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

func UpdateSteamTheme(newThemeFilePath string) error {
	afsIsInstalled := checkForAdwaitaForSteam()
	expectedPath := adwaitaForSteamPath()

	if !afsIsInstalled {
		return &missingDependencyError{
			commandName: "adwaita-for-steam",
			instructions: []string{
				fmt.Sprintf("Run `git clone https://github.com/tkashkin/Adwaita-for-Steam --revision 171aa65e8d186cf0fe723cafb4498d10f1ae8baf %s`", expectedPath),
			},
		}
	}

	omarchyThemeDirPath, err := createAdwaitaForSteamOmarchyThemeDir()
	if err != nil {
		return err
	}

	omarchyThemeFilePath := filepath.Join(omarchyThemeDirPath, "omarchy.css")
	err = fileutils.CopyFile(newThemeFilePath, omarchyThemeFilePath)
	if err != nil {
		fmt.Println("Error when copying file:", err)
		return err
	}

	err = os.Chdir(expectedPath)
	if err != nil {
		fmt.Println("Error when changing dir:", err)
		return err
	}

	cmd := exec.Command("python3", "./install.py", "--color-theme", "omarchy", "--extras", "library/hide-whats-new")
	cmd.Dir = expectedPath
	result, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error when running adwaita-for-steam:", string(result), cmd.Dir)
		return err
	}

	return nil
}
