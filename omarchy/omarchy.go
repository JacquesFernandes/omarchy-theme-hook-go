// Package omarchy is used to make shell calls specific to the omarchy os
package omarchy

import (
	"os/exec"
	"strings"
)

// ThemeList returns a list of all themes on the system, inbuilt and extra.
// Internally calls "omarchy-theme-list".
func ThemeList() ([]string, error) {
	cmd := exec.Command("omarchy-theme-list")
	output, err := cmd.Output()
	if err != nil {
		return make([]string, 0), err
	}

	stringOutput := strings.Trim(string(output), "\n")

	themeList := strings.Split(stringOutput, "\n")

	for index, theme := range themeList {
		lowered := strings.ToLower(theme)
		kebab := strings.ReplaceAll(lowered, " ", "-")

		themeList[index] = kebab
	}

	return themeList, nil
}

// ThemeDirPath returns the path to the directory of a given theme.
// The theme is usually obtained from the output of either the
// ThemeList function or the CurrentThemeName function
func ThemeDirPath(theme string) (string, error) {
	cmd := exec.Command("omarchy-theme-dir", theme)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	stringOutput := string(output)
	return strings.Trim(stringOutput, "\n"), nil
}

// CurrentThemeName fetches the name of the current omarchy theme.
// This can be used for ThemeDirPath as well.
func CurrentThemeName() (string, error) {
	cmd := exec.Command("omarchy-theme-current")
	output, err := cmd.Output()
	if err != nil {
		return "", nil
	}

	stringOutput := strings.Trim(string(output), "\n")
	lowered := strings.ToLower(stringOutput)

	return lowered, nil
}
