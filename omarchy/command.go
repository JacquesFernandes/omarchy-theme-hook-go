package omarchy

import (
	"os/exec"
	"strings"
)

func DoesCommandExist(cmd string) bool {
	_, err := exec.LookPath(cmd)

	return err == nil
}

func OmarchyThemeList() ([]string, error) {
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

func OmarchyThemeDir(theme string) (string, error) {
	cmd := exec.Command("omarchy-theme-dir", theme)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	stringOutput := string(output)
	return strings.Trim(stringOutput, "\n"), nil
}

func OmarchyThemeCurrent() (string, error) {
	cmd := exec.Command("omarchy-theme-current")
	output, err := cmd.Output()
	if err != nil {
		return "", nil
	}

	stringOutput := strings.Trim(string(output), "\n")
	lowered := strings.ToLower(stringOutput)

	return lowered, nil
}
